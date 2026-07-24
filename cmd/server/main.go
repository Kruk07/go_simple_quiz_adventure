package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"go_simple_quiz_adventure/internal/repository"
	ws "go_simple_quiz_adventure/internal/websocket"

	_ "github.com/glebarez/sqlite"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	dbPath := os.Getenv("QUIZ_DB_PATH")
	if dbPath == "" {
		dbPath = "./quiz.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		logger.Error("open database", "path", dbPath, "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error("ping database", "path", dbPath, "error", err)
		os.Exit(1)
	}

	repo := repository.NewQuestionRepository(db)
	if err := repo.InitSchema(); err != nil {
		logger.Error("init schema", "error", err)
		os.Exit(1)
	}
	if err := repo.Seed(); err != nil {
		logger.Error("seed data", "error", err)
		os.Exit(1)
	}

	server := ws.NewWebSocketServer(repo)
	http.HandleFunc("/create-room", server.CreateRoomHandler)
	http.HandleFunc("/join-room", server.JoinRoomHandler)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	port := os.Getenv("QUIZ_SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("server started", "database", dbPath, "port", port)
	logger.Error("server stopped", "error", http.ListenAndServe(":"+port, nil))
}

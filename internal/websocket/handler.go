package websocket

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"go_simple_quiz_adventure/internal/game"
	"go_simple_quiz_adventure/internal/repository"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketServer struct {
	Rooms       map[string]*game.Room
	Connections map[string]map[string]*websocket.Conn
	Mutex       sync.RWMutex
	Repo        *repository.QuestionRepository
}

func NewWebSocketServer(repo *repository.QuestionRepository) *WebSocketServer {
	return &WebSocketServer{
		Rooms:       make(map[string]*game.Room),
		Connections: make(map[string]map[string]*websocket.Conn),
		Repo:        repo,
	}
}

func (s *WebSocketServer) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomCode := game.GenerateRoomCode()
	room := game.NewRoom(roomCode)

	categories, err := s.Repo.ListCategories()
	if err != nil {
		http.Error(w, fmt.Sprintf("load categories: %v", err), http.StatusInternalServerError)
		return
	}
	room.AvailableCategories = categories
	room.LoadQuestions = s.Repo.ListQuestionsByCategory
	room.Broadcast = func(envelope game.EventEnvelope) error {
		return s.broadcastRoom(room, envelope)
	}
	slog.Info("created room", "roomCode", roomCode, "categories", len(categories))

	s.Rooms[roomCode] = room

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"roomCode": roomCode})
}

func (s *WebSocketServer) JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomCode := r.URL.Query().Get("roomCode")
	nickname := r.URL.Query().Get("nickname")
	if roomCode == "" || nickname == "" {
		http.Error(w, "roomCode and nickname are required", http.StatusBadRequest)
		return
	}

	slog.Info("join room request", "roomCode", roomCode, "nickname", nickname)

	room, ok := s.Rooms[roomCode]
	if !ok {
		http.Error(w, "room not found", http.StatusNotFound)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("upgrade: %v", err), http.StatusInternalServerError)
		return
	}

	playerID := fmt.Sprintf("p-%d", len(room.Players)+1)
	player := &game.Player{ID: playerID, Nickname: nickname}
	room.AddPlayer(player)

	s.Mutex.Lock()
	if _, ok := s.Connections[roomCode]; !ok {
		s.Connections[roomCode] = make(map[string]*websocket.Conn)
	}
	s.Connections[roomCode][playerID] = conn
	s.Mutex.Unlock()

	if err := s.broadcastRoom(room, room.LobbyUpdate()); err != nil {
		slog.Error("broadcast lobby update failed", "roomCode", roomCode, "playerID", playerID, "error", err)
		conn.Close()
		room.RemovePlayer(playerID)
		return
	}
	slog.Info("player joined room", "roomCode", roomCode, "playerID", playerID, "nickname", nickname)

	go s.readMessages(conn, room, player)
}

func (s *WebSocketServer) broadcastRoom(room *game.Room, envelope game.EventEnvelope) error {
	s.Mutex.RLock()
	conns := s.Connections[room.Code]
	s.Mutex.RUnlock()

	for _, conn := range conns {
		if err := conn.WriteJSON(envelope); err != nil {
			return err
		}
	}
	return nil
}

func (s *WebSocketServer) readMessages(conn *websocket.Conn, room *game.Room, player *game.Player) {
	defer func() {
		conn.Close()
		room.RemovePlayer(player.ID)
		s.Mutex.Lock()
		if conns, ok := s.Connections[room.Code]; ok {
			delete(conns, player.ID)
		}
		s.Mutex.Unlock()
		_ = s.broadcastRoom(room, room.LobbyUpdate())
		slog.Info("player disconnected", "roomCode", room.Code, "playerID", player.ID)
	}()

	slog.Info("websocket listener started", "roomCode", room.Code, "playerID", player.ID)
	for {
		var envelope game.EventEnvelope
		if err := conn.ReadJSON(&envelope); err != nil {
			slog.Info("websocket disconnect", "roomCode", room.Code, "playerID", player.ID, "error", err)
			return
		}

		slog.Info("received event", "roomCode", room.Code, "playerID", player.ID, "type", envelope.Type)
		switch envelope.Type {
		case "START_GAME":
			if player.ID != room.HostID {
				continue
			}
			if err := room.StartGame(); err != nil {
				slog.Error("start game failed", "roomCode", room.Code, "playerID", player.ID, "error", err)
				continue
			}
		case "VOTE_CATEGORY":
			payloadData, err := json.Marshal(envelope.Payload)
			if err != nil {
				continue
			}
			var votePayload struct {
				CategoryID string `json:"categoryId"`
			}
			if err := json.Unmarshal(payloadData, &votePayload); err != nil {
				continue
			}
			if err := room.CastCategoryVote(player.ID, votePayload.CategoryID); err != nil {
				slog.Error("category vote failed", "roomCode", room.Code, "playerID", player.ID, "categoryID", votePayload.CategoryID, "error", err)
			}
		case "SUBMIT_ANSWER":
			payloadData, err := json.Marshal(envelope.Payload)
			if err != nil {
				continue
			}
			var answerPayload struct {
				QuestionID string `json:"questionId"`
				Answer     string `json:"answer"`
			}
			if err := json.Unmarshal(payloadData, &answerPayload); err != nil {
				continue
			}
			if err := room.SubmitAnswer(player.ID, answerPayload.QuestionID, answerPayload.Answer); err != nil {
				slog.Error("submit answer failed", "roomCode", room.Code, "playerID", player.ID, "questionID", answerPayload.QuestionID, "error", err)
			}
		}
	}
}

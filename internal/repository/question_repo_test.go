package repository

import (
	"database/sql"
	"testing"

	_ "github.com/glebarez/sqlite"
)

func TestQuestionRepositoryInitSchemaAndSeed(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open database: %v", err)
	}
	defer db.Close()

	repo := NewQuestionRepository(db)
	if err := repo.InitSchema(); err != nil {
		t.Fatalf("init schema: %v", err)
	}
	if err := repo.Seed(); err != nil {
		t.Fatalf("seed: %v", err)
	}

	categories, err := repo.ListCategories()
	if err != nil {
		t.Fatalf("list categories: %v", err)
	}
	if len(categories) == 0 {
		t.Fatal("expected categories to be seeded")
	}

	questions, err := repo.ListQuestionsByCategory("cat-programming")
	if err != nil {
		t.Fatalf("list questions: %v", err)
	}
	if len(questions) == 0 {
		t.Fatal("expected questions to be seeded")
	}
}

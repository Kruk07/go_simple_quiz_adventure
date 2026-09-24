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
	if len(categories) < 20 {
		t.Fatalf("expected at least 20 categories, got %d", len(categories))
	}

	questions, err := repo.ListQuestionsByCategory("cat-programming")
	if err != nil {
		t.Fatalf("list questions: %v", err)
	}
	if len(questions) < 40 {
		t.Fatalf("expected at least 40 questions in programming, got %d", len(questions))
	}

	totalQuestions := 0
	for _, category := range categories {
		categoryQuestions, err := repo.ListQuestionsByCategory(category.ID)
		if err != nil {
			t.Fatalf("list questions for %s: %v", category.ID, err)
		}
		totalQuestions += len(categoryQuestions)
	}
	if totalQuestions < 1000 {
		t.Fatalf("expected at least 1000 questions in total, got %d", totalQuestions)
	}

	seenCorrectLetters := map[string]struct{}{}
	for _, category := range categories {
		categoryQuestions, err := repo.ListQuestionsByCategory(category.ID)
		if err != nil {
			t.Fatalf("list questions for %s: %v", category.ID, err)
		}
		for _, question := range categoryQuestions[:min(10, len(categoryQuestions))] {
			seenCorrectLetters[question.CorrectOption] = struct{}{}
		}
	}
	if len(seenCorrectLetters) < 2 {
		t.Fatalf("expected varied correct answer letters across generated questions, got %d unique letters", len(seenCorrectLetters))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

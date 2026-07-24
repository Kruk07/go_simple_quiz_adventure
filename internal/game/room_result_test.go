package game

import (
	"go_simple_quiz_adventure/internal/models"
	"testing"
)

func TestQuestionResultBroadcast(t *testing.T) {
	room := NewRoom("123456")
	room.AvailableCategories = []models.Category{
		{ID: "cat-1", Name: "Programming"},
		{ID: "cat-2", Name: "History"},
		{ID: "cat-3", Name: "Science"},
		{ID: "cat-4", Name: "Geography"},
	}
	room.LoadQuestions = func(categoryID string) ([]models.Question, error) {
		return []models.Question{
			{ID: "q-1", CategoryID: categoryID, Text: "Test?", OptionA: "A", OptionB: "B", OptionC: "C", OptionD: "D", CorrectOption: "A"},
		}, nil
	}
	captured := make([]EventEnvelope, 0)
	room.Broadcast = func(envelope EventEnvelope) error {
		captured = append(captured, envelope)
		return nil
	}

	room.AddPlayer(&Player{ID: "p-1", Nickname: "Alice"})
	room.AddPlayer(&Player{ID: "p-2", Nickname: "Bob"})

	if err := room.StartGame(); err != nil {
		t.Fatalf("StartGame failed: %v", err)
	}

	if err := room.CastCategoryVote("p-1", room.CategoryOptions[0].ID); err != nil {
		t.Fatalf("CastCategoryVote failed: %v", err)
	}
	if err := room.CastCategoryVote("p-2", room.CategoryOptions[0].ID); err != nil {
		t.Fatalf("CastCategoryVote failed: %v", err)
	}

	if err := room.SubmitAnswer("p-1", room.Questions[0].ID, "A"); err != nil {
		t.Fatalf("SubmitAnswer failed: %v", err)
	}
	if err := room.SubmitAnswer("p-2", room.Questions[0].ID, "B"); err != nil {
		t.Fatalf("SubmitAnswer failed: %v", err)
	}

	found := false
	for _, envelope := range captured {
		if envelope.Type == "QUESTION_RESULT" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected QUESTION_RESULT broadcast")
	}
}

package game

import (
	"go_simple_quiz_adventure/internal/models"
	"testing"
)

func TestRoomHostAssignmentAndLobbyUpdate(t *testing.T) {
	room := NewRoom("123456")
	first := &Player{ID: "p-1", Nickname: "Alice"}
	second := &Player{ID: "p-2", Nickname: "Bob"}

	room.AddPlayer(first)
	if !first.IsHost {
		t.Fatal("expected first player to become host")
	}
	if room.HostID != first.ID {
		t.Fatalf("expected host id %q, got %q", first.ID, room.HostID)
	}

	room.AddPlayer(second)
	if second.IsHost {
		t.Fatal("expected second player not to be host")
	}

	update := room.LobbyUpdate()
	if update.Type != "LOBBY_UPDATE" {
		t.Fatalf("expected event type LOBBY_UPDATE, got %q", update.Type)
	}

	payload, ok := update.Payload.(LobbyUpdatePayload)
	if !ok {
		t.Fatal("expected LobbyUpdatePayload type")
	}
	if payload.RoomCode != "123456" {
		t.Fatalf("expected room code 123456, got %q", payload.RoomCode)
	}
	if len(payload.Players) != 2 {
		t.Fatalf("expected 2 players, got %d", len(payload.Players))
	}

	room.RemovePlayer(first.ID)
	if room.HostID != second.ID {
		t.Fatalf("expected host transfer to %q, got %q", second.ID, room.HostID)
	}
}

func TestRoomStartGameAndCategoryVoting(t *testing.T) {
	room := NewRoom("123456")
	room.AvailableCategories = []models.Category{
		{ID: "cat-1", Name: "Programming"},
		{ID: "cat-2", Name: "History"},
		{ID: "cat-3", Name: "Science"},
		{ID: "cat-4", Name: "Geography"},
	}
	room.LoadQuestions = func(categoryID string) ([]models.Question, error) {
		return []models.Question{{ID: "q-1", CategoryID: categoryID, Text: "Test?", OptionA: "A", OptionB: "B", OptionC: "C", OptionD: "D", CorrectOption: "A"}}, nil
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

	if room.State != RoomStateCategoryVoting {
		t.Fatalf("expected room to be in CATEGORY_VOTING, got %s", room.State)
	}
	if len(captured) == 0 || captured[0].Type != "CATEGORY_VOTE_START" {
		t.Fatalf("expected first broadcast to be CATEGORY_VOTE_START, got %v", captured)
	}

	if err := room.CastCategoryVote("p-1", room.CategoryOptions[0].ID); err != nil {
		t.Fatalf("CastCategoryVote failed: %v", err)
	}
	if err := room.CastCategoryVote("p-2", room.CategoryOptions[0].ID); err != nil {
		t.Fatalf("CastCategoryVote failed: %v", err)
	}

	// after all votes cast, state should eventually move to QUESTION_ACTIVE
	if room.State != RoomStateQuestionActive {
		t.Fatalf("expected room to be in QUESTION_ACTIVE after voting, got %s", room.State)
	}
}

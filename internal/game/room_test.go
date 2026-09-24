package game

import (
	"go_simple_quiz_adventure/internal/models"
	"testing"
	"time"
)

func TestGameConfigFromEnvOverridesDefaults(t *testing.T) {
	t.Setenv("QUIZ_MAX_ROUNDS", "3")
	t.Setenv("QUIZ_QUESTIONS_PER_ROUND", "4")
	t.Setenv("QUIZ_CATEGORY_VOTE_SECONDS", "30")
	t.Setenv("QUIZ_QUESTION_SECONDS", "45")
	t.Setenv("QUIZ_ROUND_SUMMARY_SECONDS", "8")

	cfg := GameConfigFromEnv()
	if cfg.MaxRounds != 3 {
		t.Fatalf("expected MaxRounds 3, got %d", cfg.MaxRounds)
	}
	if cfg.QuestionsPerRound != 4 {
		t.Fatalf("expected QuestionsPerRound 4, got %d", cfg.QuestionsPerRound)
	}
	if cfg.CategoryVoteDuration != 30*time.Second {
		t.Fatalf("expected CategoryVoteDuration 30s, got %s", cfg.CategoryVoteDuration)
	}
	if cfg.QuestionDuration != 45*time.Second {
		t.Fatalf("expected QuestionDuration 45s, got %s", cfg.QuestionDuration)
	}
	if cfg.RoundSummaryDuration != 8*time.Second {
		t.Fatalf("expected RoundSummaryDuration 8s, got %s", cfg.RoundSummaryDuration)
	}
}

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

func TestRoomTransferHostClearsPreviousHostFlag(t *testing.T) {
	room := NewRoom("123456")
	first := &Player{ID: "p-1", Nickname: "Alice"}
	second := &Player{ID: "p-2", Nickname: "Bob"}

	room.AddPlayer(first)
	room.AddPlayer(second)
	room.RemovePlayer(first.ID)

	if room.HostID != second.ID {
		t.Fatalf("expected host transfer to %q, got %q", second.ID, room.HostID)
	}
	if !second.IsHost {
		t.Fatal("expected remaining player to be host")
	}
	if first.IsHost {
		t.Fatal("expected removed host to lose host flag")
	}
}

func TestRoomRejectsLateAnswersUsingConfiguredDeadline(t *testing.T) {
	room := NewRoom("123456")
	room.Config = GameConfig{QuestionDuration: 2, CategoryVoteDuration: 15, RoundSummaryDuration: 5}
	room.State = RoomStateQuestionActive
	room.Questions = []models.Question{{
		ID:            "q-1",
		CategoryID:    "cat-1",
		Text:          "Test?",
		OptionA:       "A",
		OptionB:       "B",
		OptionC:       "C",
		OptionD:       "D",
		CorrectOption: "A",
	}}
	room.CurrentQuestionIndex = 0
	room.QuestionStartedAt = time.Now().Add(-3 * time.Second)
	room.QuestionDeadline = time.Now().Add(-1 * time.Second)
	room.Players = map[string]*Player{
		"p-1": {ID: "p-1", Nickname: "Alice"},
	}
	room.AnswerSubmissions = map[string]AnswerSubmission{}
	room.Scoreboard = map[string]int{"p-1": 0}

	if err := room.SubmitAnswer("p-1", "q-1", "A"); err == nil {
		t.Fatal("expected late answer to be rejected")
	}
}

func TestRoomKeepsDisconnectedPlayerAvailableForReconnect(t *testing.T) {
	room := NewRoom("123456")
	player := &Player{ID: "p-1", Nickname: "Alice"}
	room.AddPlayer(player)

	if _, err := room.MarkDisconnected(player.ID, func() {}); err != nil {
		t.Fatalf("expected mark disconnected without error: %v", err)
	}
	if player.Connected {
		t.Fatal("expected player to be marked disconnected")
	}
	if !room.RestorePlayerConnection(player.ID) {
		t.Fatal("expected player to reconnect successfully")
	}
	if !player.Connected {
		t.Fatal("expected player connection state to be restored")
	}
}

func TestRoomFindsDisconnectedPlayerByNickname(t *testing.T) {
	room := NewRoom("123456")
	player := &Player{ID: "p-1", Nickname: "Alice"}
	room.AddPlayer(player)

	room.MarkDisconnected(player.ID, func() {})
	resolved := room.FindDisconnectedPlayerByNickname("Alice")
	if resolved == nil {
		t.Fatal("expected disconnected player to be found by nickname")
	}
	if resolved.ID != player.ID {
		t.Fatalf("expected resolved player id %q, got %q", player.ID, resolved.ID)
	}
}

func TestRoomReconnectPlayerByNickname(t *testing.T) {
	room := NewRoom("123456")
	player := &Player{ID: "p-1", Nickname: "Alice"}
	room.AddPlayer(player)

	if _, err := room.MarkDisconnected(player.ID, func() {}); err != nil {
		t.Fatalf("mark disconnected failed: %v", err)
	}

	reconnected, ok := room.ReconnectPlayer("Alice")
	if !ok {
		t.Fatal("expected disconnected player to reconnect by nickname")
	}
	if reconnected.ID != player.ID {
		t.Fatalf("expected reconnected id %q, got %q", player.ID, reconnected.ID)
	}
	if !reconnected.Connected {
		t.Fatal("expected reconnected player to be active")
	}
}

func TestRoomShutdownStopsTimersAndClearsState(t *testing.T) {
	room := NewRoom("123456")
	player := &Player{ID: "p-1", Nickname: "Alice"}
	room.AddPlayer(player)
	room.categoryTimer = time.AfterFunc(time.Minute, func() {})
	room.questionTimer = time.AfterFunc(time.Minute, func() {})
	room.DisconnectTimers[player.ID] = time.AfterFunc(time.Minute, func() {})

	room.Shutdown()

	if room.State != RoomStateGameOver {
		t.Fatalf("expected room to be in GAME_OVER after shutdown, got %s", room.State)
	}
	if len(room.Players) != 0 {
		t.Fatalf("expected shutdown to clear players, got %d remaining", len(room.Players))
	}
	if len(room.DisconnectTimers) != 0 {
		t.Fatalf("expected shutdown to clear disconnect timers, got %d", len(room.DisconnectTimers))
	}
}

func TestRoomProgressesThroughFullQuestionCycle(t *testing.T) {
	room := NewRoom("123456")
	room.Config = GameConfig{
		MinPlayers:           2,
		MaxRounds:            2,
		QuestionsPerRound:    2,
		CategoryVoteDuration: 20 * time.Millisecond,
		QuestionDuration:     30 * time.Millisecond,
		RoundSummaryDuration: 15 * time.Millisecond,
	}
	room.AvailableCategories = []models.Category{
		{ID: "cat-1", Name: "Science"},
		{ID: "cat-2", Name: "Programming"},
		{ID: "cat-3", Name: "History"},
		{ID: "cat-4", Name: "Geography"},
	}
	room.LoadQuestions = func(categoryID string) ([]models.Question, error) {
		return []models.Question{
			{ID: "q-1", CategoryID: categoryID, Text: "First question?", OptionA: "A", OptionB: "B", OptionC: "C", OptionD: "D", CorrectOption: "A"},
			{ID: "q-2", CategoryID: categoryID, Text: "Second question?", OptionA: "A", OptionB: "B", OptionC: "C", OptionD: "D", CorrectOption: "A"},
		}, nil
	}
	room.Broadcast = func(EventEnvelope) error { return nil }

	room.AddPlayer(&Player{ID: "p-1", Nickname: "Alice"})
	room.AddPlayer(&Player{ID: "p-2", Nickname: "Bob"})

	if err := room.StartGame(); err != nil {
		t.Fatalf("StartGame failed: %v", err)
	}
	if err := room.CastCategoryVote("p-1", "cat-1"); err != nil {
		t.Fatalf("first vote failed: %v", err)
	}
	if err := room.CastCategoryVote("p-2", "cat-1"); err != nil {
		t.Fatalf("second vote failed: %v", err)
	}
	if room.State != RoomStateQuestionActive {
		t.Fatalf("expected QUESTION_ACTIVE state after voting, got %s", room.State)
	}

	if err := room.SubmitAnswer("p-1", "q-1", "A"); err != nil {
		t.Fatalf("first answer failed: %v", err)
	}
	if err := room.SubmitAnswer("p-2", "q-1", "A"); err != nil {
		t.Fatalf("second answer failed: %v", err)
	}

	deadline := time.Now().Add(250 * time.Millisecond)
	for time.Now().Before(deadline) && room.State != RoomStateQuestionActive {
		time.Sleep(10 * time.Millisecond)
	}
	if room.State != RoomStateQuestionActive {
		t.Fatalf("expected question cycle to continue, got state %s", room.State)
	}
	if room.CurrentQuestionIndex != 1 {
		t.Fatalf("expected second question to start, got index %d", room.CurrentQuestionIndex)
	}

	if err := room.SubmitAnswer("p-1", "q-2", "A"); err != nil {
		t.Fatalf("first answer for second question failed: %v", err)
	}
	if err := room.SubmitAnswer("p-2", "q-2", "A"); err != nil {
		t.Fatalf("second answer for second question failed: %v", err)
	}

	deadline = time.Now().Add(250 * time.Millisecond)
	for time.Now().Before(deadline) && room.State != RoomStateGameOver {
		time.Sleep(10 * time.Millisecond)
	}
	if room.State != RoomStateGameOver {
		t.Fatalf("expected game over after final question, got state %s", room.State)
	}
}

func TestRoomAllowsRestartAfterGameOver(t *testing.T) {
	room := NewRoom("123456")
	room.Config = GameConfig{
		MinPlayers:           2,
		MaxRounds:            1,
		QuestionsPerRound:    1,
		CategoryVoteDuration: 20 * time.Millisecond,
		QuestionDuration:     30 * time.Millisecond,
		RoundSummaryDuration: 15 * time.Millisecond,
	}
	room.AvailableCategories = []models.Category{
		{ID: "cat-1", Name: "Science"},
		{ID: "cat-2", Name: "Programming"},
		{ID: "cat-3", Name: "History"},
		{ID: "cat-4", Name: "Geography"},
	}
	room.LoadQuestions = func(categoryID string) ([]models.Question, error) {
		return []models.Question{{ID: "q-1", CategoryID: categoryID, Text: "Test?", OptionA: "A", OptionB: "B", OptionC: "C", OptionD: "D", CorrectOption: "A"}}, nil
	}
	room.Broadcast = func(EventEnvelope) error { return nil }

	room.AddPlayer(&Player{ID: "p-1", Nickname: "Alice"})
	room.AddPlayer(&Player{ID: "p-2", Nickname: "Bob"})

	if err := room.StartGame(); err != nil {
		t.Fatalf("start game failed: %v", err)
	}
	if err := room.CastCategoryVote("p-1", "cat-1"); err != nil {
		t.Fatalf("vote 1 failed: %v", err)
	}
	if err := room.CastCategoryVote("p-2", "cat-1"); err != nil {
		t.Fatalf("vote 2 failed: %v", err)
	}
	if err := room.SubmitAnswer("p-1", "q-1", "A"); err != nil {
		t.Fatalf("submit answer 1 failed: %v", err)
	}
	if err := room.SubmitAnswer("p-2", "q-1", "A"); err != nil {
		t.Fatalf("submit answer 2 failed: %v", err)
	}

	deadline := time.Now().Add(250 * time.Millisecond)
	for time.Now().Before(deadline) && room.State != RoomStateGameOver {
		time.Sleep(10 * time.Millisecond)
	}
	if room.State != RoomStateGameOver {
		t.Fatalf("expected final state GAME_OVER, got %s", room.State)
	}

	room.ResetForNewGame()
	if room.State != RoomStateLobby {
		t.Fatalf("expected reset state LOBBY, got %s", room.State)
	}
	if room.Round != 1 {
		t.Fatalf("expected round reset to 1, got %d", room.Round)
	}
	if len(room.Scoreboard) != 2 {
		t.Fatalf("expected scoreboard to preserve players after reset, got %d entries", len(room.Scoreboard))
	}
	if err := room.StartGame(); err != nil {
		t.Fatalf("restart failed: %v", err)
	}
	if room.State != RoomStateCategoryVoting {
		t.Fatalf("expected restarted room to begin in CATEGORY_VOTING, got %s", room.State)
	}
}

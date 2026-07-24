package game

import (
	"errors"
	"log/slog"
	"math/rand"
	"sort"
	"sync"
	"time"

	"go_simple_quiz_adventure/internal/models"
)

type Player struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	IsHost    bool   `json:"isHost"`
	Connected bool   `json:"connected"`
}

type RoomState string

const (
	RoomStateLobby          RoomState = "LOBBY"
	RoomStateCategoryVoting RoomState = "CATEGORY_VOTING"
	RoomStateQuestionActive RoomState = "QUESTION_ACTIVE"
	RoomStateQuestionResult RoomState = "QUESTION_RESULT"
	RoomStateGameOver       RoomState = "GAME_OVER"
)

type LobbyUpdatePayload struct {
	RoomCode string   `json:"roomCode"`
	Players  []Player `json:"players"`
}

type CategoryVoteStartPayload struct {
	Round           int               `json:"round"`
	DurationSeconds int               `json:"durationSeconds"`
	Categories      []models.Category `json:"categories"`
}

type QuestionStartPayload struct {
	QuestionIndex   int               `json:"questionIndex"`
	TotalQuestions  int               `json:"totalQuestions"`
	Category        string            `json:"category"`
	QuestionID      string            `json:"questionId"`
	Text            string            `json:"text"`
	Options         map[string]string `json:"options"`
	DurationSeconds int               `json:"durationSeconds"`
}

type LeaderboardEntry struct {
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
	Gained   int    `json:"gained"`
}

type QuestionResultPayload struct {
	QuestionID    string             `json:"questionId"`
	CorrectAnswer string             `json:"correctAnswer"`
	Leaderboard   []LeaderboardEntry `json:"leaderboard"`
}

type RoundSummaryPayload struct {
	Round             int                `json:"round"`
	Category          string             `json:"category"`
	QuestionsAnswered int                `json:"questionsAnswered"`
	TotalQuestions    int                `json:"totalQuestions"`
	Leaderboard       []LeaderboardEntry `json:"leaderboard"`
}

type GameOverPayload struct {
	Winner      string             `json:"winner"`
	Leaderboard []LeaderboardEntry `json:"leaderboard"`
}

type EventEnvelope struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type AnswerSubmission struct {
	Answer      string
	SubmittedAt time.Time
	Score       int
}

type Room struct {
	Code                 string
	Players              map[string]*Player
	DisconnectTimers     map[string]*time.Timer
	HostID               string
	State                RoomState
	Round                int
	AvailableCategories  []models.Category
	CategoryOptions      []models.Category
	SelectedCategories   []models.Category
	CategoryVotes        map[string]string
	SelectedCategoryID   string
	SelectedCategoryName string
	Questions            []models.Question
	CurrentQuestionIndex int
	QuestionStartedAt    time.Time
	AnswerSubmissions    map[string]AnswerSubmission
	Scoreboard           map[string]int
	Broadcast            func(EventEnvelope) error
	LoadQuestions        func(categoryID string) ([]models.Question, error)
	questionTimer        *time.Timer
	categoryTimer        *time.Timer
	Mutex                sync.RWMutex
}

func NewRoom(code string) *Room {
	return &Room{
		Code:             code,
		Players:          make(map[string]*Player),
		DisconnectTimers: make(map[string]*time.Timer),
		State:            RoomStateLobby,
		Scoreboard:       make(map[string]int),
	}
}

func (r *Room) AddPlayer(player *Player) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	player.Connected = true
	if len(r.Players) == 0 {
		player.IsHost = true
		r.HostID = player.ID
	}
	r.Players[player.ID] = player
	if _, ok := r.Scoreboard[player.ID]; !ok {
		r.Scoreboard[player.ID] = 0
	}
}

func (r *Room) RemovePlayer(playerID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if _, ok := r.Players[playerID]; !ok {
		return
	}
	r.stopDisconnectTimer(playerID)
	delete(r.Players, playerID)
	delete(r.Scoreboard, playerID)
	if r.HostID == playerID {
		r.transferHost()
	}
}

func (r *Room) stopDisconnectTimer(playerID string) {
	if timer, ok := r.DisconnectTimers[playerID]; ok {
		timer.Stop()
		delete(r.DisconnectTimers, playerID)
	}
}

func (r *Room) MarkDisconnected(playerID string, removeFn func()) (bool, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	player, ok := r.Players[playerID]
	if !ok {
		return false, errors.New("player not found")
	}
	player.Connected = false
	r.stopDisconnectTimer(playerID)
	r.DisconnectTimers[playerID] = time.AfterFunc(30*time.Second, func() {
		removeFn()
	})

	if player.ID == r.HostID && r.State == RoomStateLobby {
		return true, nil
	}
	if player.ID == r.HostID {
		r.transferHost()
	}

	return false, nil
}

func (r *Room) RestorePlayerConnection(playerID string) bool {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	player, ok := r.Players[playerID]
	if !ok {
		return false
	}
	player.Connected = true
	r.stopDisconnectTimer(playerID)
	return true
}

func (r *Room) transferHost() {
	for _, player := range r.Players {
		player.IsHost = true
		r.HostID = player.ID
		return
	}
	r.HostID = ""
}

func (r *Room) PlayerList() []Player {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	players := make([]Player, 0, len(r.Players))
	for _, p := range r.Players {
		players = append(players, *p)
	}
	return players
}

func (r *Room) LobbyUpdate() EventEnvelope {
	return EventEnvelope{
		Type: "LOBBY_UPDATE",
		Payload: LobbyUpdatePayload{
			RoomCode: r.Code,
			Players:  r.PlayerList(),
		},
	}
}

func (r *Room) StartGame() error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if r.State != RoomStateLobby {
		return errors.New("game already started")
	}
	if len(r.Players) < 2 {
		return errors.New("at least 2 players required to start")
	}
	if len(r.AvailableCategories) < 4 {
		return errors.New("not enough categories")
	}

	r.State = RoomStateCategoryVoting
	r.Round = 1
	r.CategoryVotes = make(map[string]string)
	r.SelectedCategories = append([]models.Category(nil), r.AvailableCategories...)

	slog.Info("game started", "roomCode", r.Code, "round", r.Round, "players", len(r.Players), "availableCategories", len(r.SelectedCategories))
	return r.startCategoryVoting()
}

func (r *Room) startCategoryVoting() error {
	r.State = RoomStateCategoryVoting
	if len(r.SelectedCategories) == 0 {
		r.SelectedCategories = append([]models.Category(nil), r.AvailableCategories...)
	}
	r.CategoryOptions = sampleCategories(r.SelectedCategories, 4)
	r.CategoryVotes = make(map[string]string)
	slog.Info("category voting started", "roomCode", r.Code, "round", r.Round, "options", len(r.CategoryOptions), "remainingCategories", len(r.SelectedCategories))
	if r.Broadcast == nil {
		return errors.New("broadcast callback not configured")
	}

	envelope := EventEnvelope{
		Type: "CATEGORY_VOTE_START",
		Payload: CategoryVoteStartPayload{
			Round:           r.Round,
			DurationSeconds: 15,
			Categories:      r.CategoryOptions,
		},
	}

	if err := r.Broadcast(envelope); err != nil {
		return err
	}

	r.categoryTimer = time.AfterFunc(15*time.Second, func() {
		_ = r.FinalizeCategoryVoting()
	})

	return nil
}

func (r *Room) CastCategoryVote(playerID, categoryID string) error {
	r.Mutex.Lock()
	if r.State != RoomStateCategoryVoting {
		r.Mutex.Unlock()
		return errors.New("not accepting category votes")
	}
	if _, ok := r.Players[playerID]; !ok {
		r.Mutex.Unlock()
		return errors.New("player not in room")
	}
	r.CategoryVotes[playerID] = categoryID
	votes := len(r.CategoryVotes)
	playerCount := len(r.Players)
	if r.categoryTimer != nil && votes == playerCount {
		r.categoryTimer.Stop()
	}
	r.Mutex.Unlock()

	slog.Info("category vote cast", "roomCode", r.Code, "playerID", playerID, "categoryID", categoryID, "votes", votes, "players", playerCount)
	if votes == playerCount {
		return r.FinalizeCategoryVoting()
	}

	return nil
}

func (r *Room) FinalizeCategoryVoting() error {
	r.Mutex.Lock()
	if r.State != RoomStateCategoryVoting {
		r.Mutex.Unlock()
		return nil
	}
	selected := r.pickWinningCategoryLocked()
	r.SelectedCategoryID = selected.ID
	r.SelectedCategoryName = selected.Name
	r.State = RoomStateQuestionActive
	r.removeSelectedCategory(selected.ID)
	r.Mutex.Unlock()

	slog.Info("category voting finalized", "roomCode", r.Code, "round", r.Round, "categoryID", r.SelectedCategoryID, "categoryName", r.SelectedCategoryName, "remainingCategories", len(r.SelectedCategories))
	return r.startQuestionRound()
}

func (r *Room) pickWinningCategoryLocked() models.Category {
	counts := make(map[string]int)
	for _, vote := range r.CategoryVotes {
		counts[vote]++
	}

	if len(counts) == 0 {
		return r.CategoryOptions[rand.Intn(len(r.CategoryOptions))]
	}

	bestCount := 0
	var winners []models.Category
	for _, category := range r.CategoryOptions {
		count := counts[category.ID]
		if count > bestCount {
			bestCount = count
			winners = []models.Category{category}
		} else if count == bestCount {
			winners = append(winners, category)
		}
	}

	return winners[rand.Intn(len(winners))]
}

func (r *Room) removeSelectedCategory(categoryID string) {
	remaining := make([]models.Category, 0, len(r.SelectedCategories))
	for _, category := range r.SelectedCategories {
		if category.ID != categoryID {
			remaining = append(remaining, category)
		}
	}
	r.SelectedCategories = remaining
}

func (r *Room) startQuestionRound() error {
	if r.LoadQuestions == nil {
		return errors.New("question loader not configured")
	}

	questions, err := r.LoadQuestions(r.SelectedCategoryID)
	if err != nil {
		return err
	}

	r.Questions = sampleQuestions(questions, 5)
	r.CurrentQuestionIndex = 0
	r.AnswerSubmissions = make(map[string]AnswerSubmission)
	slog.Info("question round started", "roomCode", r.Code, "category", r.SelectedCategoryName, "questions", len(r.Questions))

	return r.advanceQuestion()
}

func (r *Room) advanceQuestion() error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if r.CurrentQuestionIndex >= len(r.Questions) {
		r.State = RoomStateGameOver
		slog.Info("all questions complete", "roomCode", r.Code)
		return r.broadcastGameOver()
	}

	r.State = RoomStateQuestionActive
	question := r.Questions[r.CurrentQuestionIndex]
	r.QuestionStartedAt = time.Now()
	r.AnswerSubmissions = make(map[string]AnswerSubmission)
	slog.Info("question started", "roomCode", r.Code, "questionIndex", r.CurrentQuestionIndex+1, "questionID", question.ID)
	slog.Info("state transition", "roomCode", r.Code, "from", "QUESTION_RESULT", "to", "QUESTION_ACTIVE", "questionIndex", r.CurrentQuestionIndex+1)

	if r.Broadcast == nil {
		return errors.New("broadcast callback not configured")
	}

	envelope := EventEnvelope{
		Type: "QUESTION_START",
		Payload: QuestionStartPayload{
			QuestionIndex:  r.CurrentQuestionIndex + 1,
			TotalQuestions: len(r.Questions),
			Category:       r.SelectedCategoryName,
			QuestionID:     question.ID,
			Text:           question.Text,
			Options: map[string]string{
				"A": question.OptionA,
				"B": question.OptionB,
				"C": question.OptionC,
				"D": question.OptionD,
			},
			DurationSeconds: 20,
		},
	}

	if err := r.Broadcast(envelope); err != nil {
		return err
	}

	r.questionTimer = time.AfterFunc(20*time.Second, func() {
		r.completeQuestion()
	})

	return nil
}

func (r *Room) SubmitAnswer(playerID, questionID, answer string) error {
	r.Mutex.Lock()
	if r.State != RoomStateQuestionActive {
		r.Mutex.Unlock()
		return errors.New("not accepting answers")
	}
	question := r.CurrentQuestion()
	if question == nil || question.ID != questionID {
		r.Mutex.Unlock()
		return errors.New("invalid question")
	}
	if _, ok := r.Players[playerID]; !ok {
		r.Mutex.Unlock()
		return errors.New("player not in room")
	}
	if _, answered := r.AnswerSubmissions[playerID]; answered {
		r.Mutex.Unlock()
		return errors.New("already answered")
	}

	timeElapsed := time.Now().Sub(r.QuestionStartedAt)
	remaining := 20*time.Second - timeElapsed
	if remaining < 0 {
		remaining = 0
	}
	score := 0
	if answer == question.CorrectOption {
		bonus := int(50 * remaining.Seconds() / 20)
		score = 100 + bonus
	}

	r.AnswerSubmissions[playerID] = AnswerSubmission{
		Answer:      answer,
		SubmittedAt: time.Now(),
		Score:       score,
	}
	r.Scoreboard[playerID] += score
	allAnswered := len(r.AnswerSubmissions) == len(r.Players)
	r.Mutex.Unlock()

	slog.Info("answer submitted", "roomCode", r.Code, "playerID", playerID, "questionID", questionID, "answer", answer, "score", score, "allAnswered", allAnswered)
	if allAnswered {
		if r.questionTimer != nil {
			r.questionTimer.Stop()
		}
		r.completeQuestion()
	}

	return nil
}

func (r *Room) CurrentQuestion() *models.Question {
	if r.CurrentQuestionIndex >= len(r.Questions) {
		return nil
	}
	return &r.Questions[r.CurrentQuestionIndex]
}

func (r *Room) completeQuestion() {
	r.Mutex.Lock()
	if r.State != RoomStateQuestionActive {
		r.Mutex.Unlock()
		return
	}
	r.State = RoomStateQuestionResult
	question := r.CurrentQuestion()
	leaderboard := r.buildLeaderboard(question)
	r.Mutex.Unlock()

	slog.Info("question complete", "roomCode", r.Code, "questionID", question.ID, "leaderboardCount", len(leaderboard))
	if r.Broadcast != nil {
		_ = r.Broadcast(EventEnvelope{
			Type: "QUESTION_RESULT",
			Payload: QuestionResultPayload{
				QuestionID:    question.ID,
				CorrectAnswer: question.CorrectOption,
				Leaderboard:   leaderboard,
			},
		})
	}

	time.AfterFunc(5*time.Second, func() {
		r.Mutex.Lock()
		r.CurrentQuestionIndex++
		finishedRound := r.CurrentQuestionIndex >= len(r.Questions)
		r.Mutex.Unlock()

		if finishedRound {
			_ = r.broadcastRoundSummary()
			if r.Round >= 4 || len(r.SelectedCategories) == 0 {
				_ = r.broadcastGameOver()
				return
			}
			r.Round++
			_ = r.startCategoryVoting()
			return
		}

		_ = r.advanceQuestion()
	})
}

func (r *Room) buildLeaderboard(question *models.Question) []LeaderboardEntry {
	entries := make([]LeaderboardEntry, 0, len(r.Players))
	for _, player := range r.Players {
		score := r.Scoreboard[player.ID]
		gained := 0
		if answer, ok := r.AnswerSubmissions[player.ID]; ok {
			gained = answer.Score
		}
		entries = append(entries, LeaderboardEntry{
			Nickname: player.Nickname,
			Score:    score,
			Gained:   gained,
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Score > entries[j].Score
	})
	return entries
}

func (r *Room) broadcastGameOver() error {
	if r.Broadcast == nil {
		return nil
	}
	leaderboard := r.buildLeaderboard(nil)
	winner := ""
	if len(leaderboard) > 0 {
		winner = leaderboard[0].Nickname
	}
	slog.Info("game over", "roomCode", r.Code, "winner", winner)
	if err := r.Broadcast(EventEnvelope{
		Type: "GAME_OVER",
		Payload: GameOverPayload{
			Winner:      winner,
			Leaderboard: leaderboard,
		},
	}); err != nil {
		return err
	}
	return nil
}

func (r *Room) broadcastRoundSummary() error {
	if r.Broadcast == nil {
		return nil
	}
	leaderboard := r.buildLeaderboard(nil)
	summary := RoundSummaryPayload{
		Round:             r.Round,
		Category:          r.SelectedCategoryName,
		QuestionsAnswered: len(r.Questions),
		TotalQuestions:    len(r.Questions),
		Leaderboard:       leaderboard,
	}
	slog.Info("round summary", "roomCode", r.Code, "round", r.Round, "category", r.SelectedCategoryName, "questions", len(r.Questions))
	return r.Broadcast(EventEnvelope{
		Type:    "ROUND_SUMMARY",
		Payload: summary,
	})
}

func (r *Room) broadcast(envelope EventEnvelope) error {
	if r.Broadcast == nil {
		return nil
	}
	return r.Broadcast(envelope)
}

func sampleCategories(all []models.Category, count int) []models.Category {
	if len(all) <= count {
		return append([]models.Category(nil), all...)
	}
	indices := rand.Perm(len(all))[:count]
	out := make([]models.Category, 0, count)
	for _, idx := range indices {
		out = append(out, all[idx])
	}
	return out
}

func sampleQuestions(all []models.Question, count int) []models.Question {
	if len(all) <= count {
		return append([]models.Question(nil), all...)
	}
	indices := rand.Perm(len(all))[:count]
	out := make([]models.Question, 0, count)
	for _, idx := range indices {
		out = append(out, all[idx])
	}
	return out
}

func GenerateRoomCode() string {
	const digits = "0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = digits[rand.Intn(len(digits))]
	}
	return string(code)
}

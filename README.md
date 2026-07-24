# go_simple_quiz_adventure

A multiplayer quiz game backend built in Go with WebSocket support, SQLite storage, and a browser-based test client.

## Requirements

- Go 1.22 or newer
- SQLite support via `github.com/glebarez/sqlite`

## Setup

1. Install Go.
2. From the project root, fetch dependencies:
   ```bash
   go mod tidy
   ```

## Run the server

1. Start the server from the project root:
   ```bash
   go run ./cmd/server
   ```

2. By default, the server listens on port `8080` and uses `./quiz.db` as the SQLite database file.

3. Optional environment variables:
   - `QUIZ_DB_PATH` — path to the SQLite file (`./quiz.db` by default)
   - `QUIZ_SERVER_PORT` — server port (`8080` by default)

## Browser test client

The project includes a simple browser-based test client in `web/index.html`.

1. Start the server.
2. Open `http://localhost:8080/` in a browser.
3. Use the page to:
   - create a room
   - join a room with a nickname
   - start the game
   - vote for categories from a set of 4 options
   - submit answers
   - see event output from the WebSocket server

## Game Flow

- The game consists of 4 rounds.
- Each round begins with category voting: 4 random categories are selected from the remaining available pool.
- Players vote on one category; the winning category is used for the round.
- Each round contains 5 questions from the selected category.
- After each question, the server broadcasts the result and leaderboard.
- After each round, the server broadcasts a round summary.
- The chosen category is removed from the pool so future rounds do not repeat it.
- After 4 rounds, the game ends and the winner is declared.

## API Endpoints

- `GET /create-room`
  - Creates a new room and returns a JSON payload with `roomCode`.

- `GET /join-room?roomCode=<code>&nickname=<name>`
  - Upgrades the request to a WebSocket connection and joins the player to the room.

### WebSocket events

Client -> Server:
- `START_GAME` — start the game as the host
- `VOTE_CATEGORY` — vote for a category
- `SUBMIT_ANSWER` — submit an answer to the current question

Server -> Client:
- `LOBBY_UPDATE`
- `CATEGORY_VOTE_START`
- `QUESTION_START`
- `QUESTION_RESULT`
- `ROUND_SUMMARY`
- `GAME_OVER`

## Testing

Run all Go tests with:

```bash
go test ./...
```

## Notes

- The server bootstraps the database schema and seeds a set of sample categories and questions on startup.
- The game selects 4 categories at random from the available pool on each round and removes the chosen category from later rounds.
- Questions are sampled randomly from the selected category, and chosen questions are not repeated within the same round.
- The browser test client is designed for quick manual verification and does not represent the final production UI.
- Room state, question flow, and scoring are implemented in the backend, with room lifecycle driven by WebSocket messages.
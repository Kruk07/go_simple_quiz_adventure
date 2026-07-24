# 5. Implementation Plan

This plan translates the game rules, WebSocket protocol, state machine, and database schema into a concrete build roadmap for the Go backend.

## 1. Overall Goal

Build a multiplayer quiz game backend that supports:
- room creation and joining via a 6-digit room code
- lobby management and host-controlled game start
- category voting and question rounds
- scoring with speed bonuses
- game-over flow
- real-time updates over WebSockets
- persistent question and category storage

## 2. Assumptions

- The first implementation focuses on the backend and WebSocket gameplay loop.
- A simple frontend can connect later using the documented event format.
- SQLite is the recommended default for local development and early testing.
- PostgreSQL can be introduced later if persistent multi-user deployment is needed.

## 3. Implementation Phases

### Phase 1 — Project Foundation

Tasks:
1. Initialize the Go module and create the proposed package structure.
2. Add the core dependencies:
   - `gorilla/websocket` for real-time connections
   - a database driver (`sqlite3` or `pgx`)
   - a logger and config loader
3. Create a basic server entrypoint in `cmd/server/main.go`.
4. Define configuration for:
   - server host/port
   - database path or DSN
   - room timeout values

Deliverables:
- runnable server skeleton
- project layout aligned with the docs

---

### Phase 2 — Data Model and Repository Layer

Tasks:
1. Define the domain models for:
   - `Category`
   - `Question`
   - `Player`
   - `Room`
   - `Client`
   - `GameMessage`
2. Create repository interfaces and implementations for:
   - loading categories
   - loading questions by category
   - storing and retrieving quiz data
3. Add seed data for a small initial question set.
4. Implement database schema creation on startup.

Deliverables:
- working repository layer
- seeded question/category data for development

---

### Phase 3 — Room Engine and State Machine

Tasks:
1. Implement a `Room` type that owns the game lifecycle for one room.
2. Introduce a `GameState` enum covering:
   - `LOBBY`
   - `CATEGORY_VOTING`
   - `QUESTION_ACTIVE`
   - `QUESTION_RESULT`
   - `ROUND_SUMMARY`
   - `GAME_OVER`
3. Build the room event loop using the hub pattern described in the docs.
4. Implement state transitions for:
   - lobby start
   - category voting timeout / fallback
   - question progression
   - round completion
   - game end
5. Add timers and phase durations based on the spec.

Deliverables:
- room engine capable of advancing through game phases automatically
- deterministic transition logic for each state

---

### Phase 4 — WebSocket Communication Layer

Tasks:
1. Implement a WebSocket connection handler.
2. Support the client-to-server event types:
   - `VOTE_CATEGORY`
   - `SUBMIT_ANSWER`
3. Support the server-to-client broadcast events:
   - `LOBBY_UPDATE`
   - `CATEGORY_VOTE_START`
   - `QUESTION_START`
   - `QUESTION_RESULT`
4. Add JSON envelope validation and basic error handling.
5. Route client messages to the appropriate room and player context.

Deliverables:
- working real-time gameplay transport
- event contract matching the documented protocol

---

### Phase 5 — Core Gameplay Features

Tasks:
1. Implement room creation and join flow.
2. Enforce the minimum player requirement before game start.
3. Implement category voting:
   - show 4 random categories
   - accept votes from players
   - select winner by vote count
   - fall back to random selection if no votes are cast
4. Implement quiz question flow:
   - show one question at a time
   - enforce answer submission deadline
   - track response timing for speed scoring
5. Implement scoring:
   - base 100 points for correct answers
   - speed bonus as defined in the spec
6. Implement round summaries and leaderboard updates.
7. Implement end-of-game flow and winner announcement.
8. Implement game-over screen logic and room shutdown/reset.

Deliverables:
- fully playable game loop from lobby to completion
- scoring behavior matching the spec

---

### Phase 6 — Edge Cases and Reliability

Tasks:
1. Handle player disconnects:
   - keep the session alive in memory for 30 seconds
   - remove the player if reconnect does not occur
2. Handle host leaving:
   - dissolve the room during lobby
   - transfer host role mid-game to the next connected player
3. Handle invalid client actions:
   - unknown room
   - duplicate vote
   - late answer submissions
4. Add logging and basic metrics for room events.

Deliverables:
- graceful handling of common connection and room-management problems

---

### Phase 7 — Testing and Polish

Tasks:
1. Add unit tests for:
   - scoring logic
   - category selection logic
   - state transitions
2. Add integration tests for:
   - room creation/join
   - WebSocket message handling
   - full round flow
3. Test the server with a small set of simulated clients.
4. Improve failure messages and edge-case behavior.

Deliverables:
- test coverage for the main gameplay paths
- confidence that the room loop behaves correctly

---

## 4. Suggested Implementation Order

1. Setup module and basic server
2. Database + repositories
3. Room model and state machine
4. WebSocket transport
5. Lobby and room lifecycle
6. Category voting and question flow
7. Scoring and game-over flow
8. Disconnect/host-transfer handling
9. Testing and stabilization

## 5. Recommended Deliverables by Milestone

### Milestone 1
- server boots
- DB schema initializes
- categories/questions can be loaded

### Milestone 2
- rooms can be created and joined
- lobby updates are broadcast

### Milestone 3
- host can start a game
- category voting and question flow run end-to-end

### Milestone 4
- scoring and game-over flow work
- disconnect/host-transfer edge cases are covered

### Milestone 5
- full client/server integration testing
- host disconnect, reconnection, and lobby recovery behavior
- simple frontend or WebSocket test harness for manual verification
- polish error handling, logging, and deployment configuration

## 6. Open Questions

The following decisions should be confirmed before implementation begins:
1. Should the initial version use SQLite or PostgreSQL?
- I'm more familiar with PostgreSQL.
2. Is the first milestone backend-only, or should a simple frontend also be scaffolded?
- We should have simple frontend implemented to be able to test app
3. Should room codes be generated as numeric 6-digit strings only, or should they support a custom prefix format?
- 6 number digit or just create a link which host can share for players. In that way players after clicking this link should be able to join the game
4. Should the system support persistent room history and saved game results beyond the current session?
- No. System should save current state of the game. When the game is ending, show final results screen, after closing this screen system shouldn't have any info about last game

## 7. Suggested Next Step

Start by finishing Milestone 4 and then move into Milestone 5 for full integration testing, reconnect/recovery behavior, and polish.

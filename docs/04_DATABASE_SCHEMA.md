Markdown

# 4. Database Schema & Directory Layout

## 4.1. Database Schema (PostgreSQL / SQLite)

```sql
CREATE TABLE categories (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE questions (
    id VARCHAR(36) PRIMARY KEY,
    category_id VARCHAR(36) REFERENCES categories(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    option_a VARCHAR(255) NOT NULL,
    option_b VARCHAR(255) NOT NULL,
    option_c VARCHAR(255) NOT NULL,
    option_d VARCHAR(255) NOT NULL,
    correct_option CHAR(1) NOT NULL -- 'A', 'B', 'C', or 'D'
);

4.2. Proposed Directory Layout (Standard Go Project Layout)
Plaintext

quiz-game-backend/
├── cmd/
│   └── server/
│       └── main.go           # Application entrypoint
├── internal/
│   ├── game/                 # Core engine: rooms, states, logic
│   │   ├── room.go
│   │   ├── state.go
│   │   └── player.go
│   ├── websocket/            # WS connection handler (e.g., gorilla/websocket)
│   │   ├── client.go
│   │   └── handler.go
│   ├── repository/           # Database queries (questions & categories)
│   │   └── question_repo.go
│   └── models/               # DTOs and data structs
│       ├── event.go
│       └── question.go
├── docs/                     # Project documentation
├── go.mod
└── go.sum
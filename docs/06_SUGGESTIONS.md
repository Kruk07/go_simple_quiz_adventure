# 6. Project Improvement Suggestions

## 6.1. Gameplay and features
- Add support for more than 4 rounds or configurable round count.
- Introduce a timed bonus multiplier and explicit scoring rules in docs.
- Add category preview or “most voted” feedback during voting.
- Add player answer reveal UI state so everyone sees the correct answer after each question.
- Support “skip question” or “hint” mechanics for future game modes.

## 6.2. Architecture and reliability
- Extract room state transitions into a dedicated state machine package/file.
- Decouple timers from game logic by using an injectable scheduler/test timer.
- Persist room history and leaderboard state in SQLite for reconnect/recovery.
- Add an explicit `RoomStateRoundSummary` state if the round summary phase is part of the flow.

## 6.3. Testing and validation
- Add unit tests for:
  - category sampling and non-repeating selection
  - vote winner tie-break logic
  - question advance / state transition flow
  - answer scoring and leaderboard ordering
- Add integration tests for WebSocket event sequences:
  - lobby → category voting → question flow → round summary → game over
  - partial answer submission and timeout handling

## 6.4. Documentation
- Add a dedicated `docs/06_SUGGESTIONS.md` for future roadmap items.
- Expand `docs/02_WEBSOCKET_PROTOCOL.md` with concrete event examples for `ROUND_SUMMARY` and `GAME_OVER`.
- Add a `docs/07_CHANGELOG.md` or project notes file for implemented changes and known issues.

## 6.5. Developer experience
- Add environment configuration examples in `README.md`.
- Add `make test` / `go test ./...` instructions and any common troubleshooting notes.
- Add a `web/README.md` or client usage guide for the browser test client.
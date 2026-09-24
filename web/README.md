# Browser Test Client

This folder contains the browser-based playtest client for the quiz game.

## Quick start

1. Start the server from the project root:
   ```bash
   go run ./cmd/server
   ```
2. Open the app in a browser:
   ```text
   http://localhost:8080/
   ```
3. Create a room and share the code with other players.
4. Join with a nickname from each device.
5. Start the game from the host device.

## Mobile use

- Open the page on each phone using the same Wi-Fi network.
- Keep the browser at 100% zoom.
- Use short nicknames for leaderboard readability.
- Use the host device for room creation and game start.

## Client actions

- Create Room: generates a new room code.
- Join Room: connects the current browser to the selected room.
- Start Game: starts the round flow as the host.
- Vote Category: chooses one category from the available options.
- Submit Answer: sends the selected answer for the active question.

## Notes

This client is meant for quick validation and playtesting, not for the final production UI. It is intentionally lightweight so you can test the actual multiplayer room logic on real devices.

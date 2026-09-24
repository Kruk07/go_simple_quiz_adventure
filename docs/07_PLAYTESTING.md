# 7. Playtesting Guide

## Mobile-first setup
This project is designed for quick local playtesting with phones on the same network. The browser client is intentionally simple, but it is usable on mobile devices when the screen is sized correctly.

## Local Wi-Fi testing
1. Start the server from the project root:
   ```bash
   go run ./cmd/server
   ```
2. Find the laptop's local IP address.
   - Windows: `ipconfig`
   - Mac/Linux: `ip addr` or `ifconfig`
3. Connect all phones to the same Wi-Fi network as the laptop.
4. Open the game URL in each phone browser:
   ```text
   http://<laptop-ip>:8080/
   ```
5. On the host device, create a room and share the room code.
6. Other players enter the room code and a nickname, then join.
7. The host starts the game and everyone votes and answers from their phone.

## Recommended mobile checklist
- Keep the browser zoom at 100%.
- Use a short nickname so the lobby stays readable.
- Test with 2–3 players first before larger sessions.
- Have one person keep the host device visible for the lobby and category voting flow.
- Make sure the Wi-Fi signal is stable before running a timing-heavy test.

## Firewall
- Allow port `8080` through Windows Firewall.
- If needed, allow `go.exe` or the local server binary in the firewall rules.
- If you are testing from a phone on the same network, make sure the laptop is not blocking inbound connections for the local network profile.
- If local network discovery fails, temporarily pause VPN, antivirus firewall filtering, or mobile hotspot protection for the test.
- For a quick fallback, connect your phone to the same hotspot as the laptop and confirm the browser can reach `http://<laptop-ip>:8080/` before the game starts.
- When the default port is busy, start the server with a different port, such as `QUIZ_SERVER_PORT=8081 go run ./cmd/server`, and use that URL instead.

## Tuning for slow mobile networks
- Increase the category-vote window with `QUIZ_CATEGORY_VOTE_SECONDS=30`.
- Increase the question timer with `QUIZ_QUESTION_SECONDS=30` for slower devices or spotty connections.
- Reduce the round count during a live test with `QUIZ_MAX_ROUNDS=2` if you want faster loops.
- Keep the environment variables consistent across all participants so everyone experiences the same timing.

## After the game ends
- The host can tap `Restart Game` to reset the room and begin a fresh match.
- The host or any player can tap `Leave Room` to exit the current room and return to the connection screen.
- Use this flow during live tests to quickly rerun the same game without restarting the server.

## Remote testing
For testing from outside the local network, use a tunnel:

```bash
ngrok http 8080
```

Then open the public URL shown by ngrok on each device.

## Best practices during a live playtest
- Confirm everyone can join before starting a game.
- Check the lobby list and verify the host is assigned correctly.
- Test category voting, answer submission, round summaries, and the game-over screen.
- Watch for disconnects, stale timers, or missed answer submissions.
- Note any latency issues caused by slower mobile connections.

## Notes
- Local Wi-Fi is the fastest and most reliable option for playtesting.
- The browser test client is not the final production UI, but it is suitable for validating multiplayer flow on real devices.
- For a polished mobile experience, the next pass will likely be a dedicated app or a mobile-optimized frontend rather than a desktop-first test page.

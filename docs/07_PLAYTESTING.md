# 7. Playtesting Guide

## Local Wi-Fi testing
1. Run the server:
   ```bash
   go run ./cmd/server
   ```
2. Find the laptop IP address.
3. Connect phones to the same Wi-Fi.
4. Open:
   `http://<laptop-ip>:8080/`

## Firewall
- Allow port `8080` through Windows Firewall.
- Alternatively allow `go.exe` if needed.

## Remote testing
- Use `ngrok`:
  ```bash
  ngrok http 8080
  ```
- Share the public URL shown by ngrok.

## Best practices
- Start with 2–3 players.
- Confirm everyone can join before starting a game.
- Test category voting, answer submission, and round summary.
- Note any disconnects or latency issues.

## Notes
- Local Wi-Fi is easiest and fastest.
- For mobile remote access, tunneling services are the simplest option.

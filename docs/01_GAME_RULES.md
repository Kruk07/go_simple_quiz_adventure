# 1. Game Logic & Flow Specification

## 1.1. Core Game Loop
1. **Lobby:** The host creates a room (receives a 6-digit room code). Players join using the code and a display nickname.
2. **Start:** The host clicks "Start" (requires a minimum of 2 players).
3. **4-Round Loop:**
   - **Category Voting Phase (15s):** The server randomly picks 4 categories. Players cast their votes. The category with the most votes wins (in case of a tie, the server picks randomly among the tied ones).
   - **Question Phase (5 questions per round):**
     - Display ABCD question (20s to answer).
     - Question summary (5s): Shows the correct answer and points awarded for the question.
   - **Round Summary (10s):** Displays the current leaderboard after the completed round.
4. **Game Over Screen:** Announcement of the winner, podium display, and option to return to lobby / close room.

## 1.2. Scoring System
* **Standard question:** 100 base points for a correct answer.
* **Speed bonus:** Up to +50 additional points based on response speed:
  $$\text{Points} = 100 + \left(50 \times \frac{\text{Remaining Time}}{\text{Total Time}}\right)$$

## 1.3. Edge Cases
* **Player Disconnection:** If a player loses connection, their session remains active in RAM for 30 seconds. If they fail to reconnect, they are removed from the room.
* **Host Leaving:** If the host leaves during the lobby phase, the room is disbanded. If they leave mid-game, host privileges automatically transfer to the next connected player.
* **No Category Votes:** If no player votes within 15 seconds, the server randomly picks one of the 4 presented categories.
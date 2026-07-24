# 2. WebSocket Protocol Specification

## 2.1. Envelope Format
Every message transmitted over WebSockets (both directions) uses the following JSON envelope layout:

json
{
  "type": "EVENT_NAME",
  "payload": { ... }
}

2.2. Events: Client -> Server
1. VOTE_CATEGORY
JSON

{
  "type": "VOTE_CATEGORY",
  "payload": {
    "categoryId": "cat-go-123"
  }
}

2. SUBMIT_ANSWER
JSON

{
  "type": "SUBMIT_ANSWER",
  "payload": {
    "questionId": "q-99",
    "answer": "B"
  }
}

2.3. Events: Server -> Client (Broadcast)
1. LOBBY_UPDATE
JSON

{
  "type": "LOBBY_UPDATE",
  "payload": {
    "roomCode": "482910",
    "players": [
      { "id": "p-1", "nickname": "Alice", "isHost": true },
      { "id": "p-2", "nickname": "Bob", "isHost": false }
    ]
  }
}

2. CATEGORY_VOTE_START
JSON

{
  "type": "CATEGORY_VOTE_START",
  "payload": {
    "round": 1,
    "durationSeconds": 15,
    "categories": [
      { "id": "cat-1", "name": "Programming" },
      { "id": "cat-2", "name": "History" },
      { "id": "cat-3", "name": "Pop Culture" },
      { "id": "cat-4", "name": "Geography" }
    ]
  }
}

3. QUESTION_START
JSON

{
  "type": "QUESTION_START",
  "payload": {
    "questionIndex": 1,
    "totalQuestions": 5,
    "category": "Programming",
    "questionId": "q-99",
    "text": "Which data type in Go represents UTF-8 text?",
    "options": {
      "A": "char[]",
      "B": "string",
      "C": "Text",
      "D": "varchar"
    },
    "durationSeconds": 20
  }
}

4. QUESTION_RESULT
JSON

{
  "type": "QUESTION_RESULT",
  "payload": {
    "questionId": "q-99",
    "correctAnswer": "B",
    "leaderboard": [
      { "nickname": "Bob", "score": 140, "gained": 140 },
      { "nickname": "Alice", "score": 0, "gained": 0 }
    ]
  }
}

5. ROUND_SUMMARY
JSON

{
  "type": "ROUND_SUMMARY",
  "payload": {
    "round": 1,
    "category": "Science",
    "questionsAnswered": 5,
    "totalQuestions": 5,
    "leaderboard": [
      { "nickname": "Bob", "score": 620 },
      { "nickname": "Alice", "score": 480 }
    ]
  }
}

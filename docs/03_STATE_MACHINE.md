Markdown

# 3. State Machine & Go Architecture

## 3.1. State Machine Diagram (Mermaid)

```mermaid
stateDiagram-v2
    [*] --> LOBBY : Room Created
    LOBBY --> CATEGORY_VOTING : Host clicks START
    
    state RoundLoop {
        CATEGORY_VOTING --> QUESTION_ACTIVE : Timeout / All Voted
        QUESTION_ACTIVE --> QUESTION_RESULT : Timeout / All Answered
        QUESTION_RESULT --> QUESTION_ACTIVE : Next Question (< 5)
        QUESTION_RESULT --> ROUND_SUMMARY : 5 Questions Completed
    }
    
    ROUND_SUMMARY --> CATEGORY_VOTING : Next Round (< 4)
    ROUND_SUMMARY --> GAME_OVER : After Round 4


3.2. Room Event Loop Architecture (Hub Pattern in Go)

Instead of traditional locking (sync.Mutex), each room executes as an isolated goroutine processing state transitions through event channels:
Go

type Room struct {
    Code       string
    Players    map[string]*Client
    State      GameState
    Broadcast  chan Message
    Register   chan *Client
    Unregister chan *Client
    Action     chan ClientAction
}

func (r *Room) Run() {
    for {
        select {
        case client := <-r.Register:
            r.Players[client.ID] = client
        case action := <-r.Action:
            // Safe concurrent processing without explicit mutex locking!
            r.handleAction(action)
        case <-r.Timer.C:
            // State transition on timer tick
            r.advanceState()
        }
    }
}
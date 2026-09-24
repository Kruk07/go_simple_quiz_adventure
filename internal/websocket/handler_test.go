package websocket

import (
	"testing"
	"time"

	"go_simple_quiz_adventure/internal/game"

	"github.com/gorilla/websocket"
)

func TestWebSocketServerRemoveRoomDoesNotDeadlock(t *testing.T) {
	server := NewWebSocketServer(nil)
	room := game.NewRoom("ABC123")
	server.Rooms[room.Code] = room
	server.Connections[room.Code] = map[string]*websocket.Conn{
		"p-1": nil,
	}

	done := make(chan struct{})
	go func() {
		server.removeRoom(room.Code)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("removeRoom deadlocked while cleaning up an empty room")
	}

	if _, ok := server.Rooms[room.Code]; ok {
		t.Fatal("expected room to be removed")
	}
	if _, ok := server.Connections[room.Code]; ok {
		t.Fatal("expected room connections to be removed")
	}
}

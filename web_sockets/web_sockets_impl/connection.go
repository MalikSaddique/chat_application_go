package websocketsimpl

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var connMap = make(map[string]*websocket.Conn)
var connLock = sync.Mutex{}

func (w *WebSocketsImpl) AddConn(userID string, wsConn *websocket.Conn) error {
	connLock.Lock()
	connMap[userID] = wsConn
	connLock.Unlock()

	log.Println("User connected:", userID)

	defer func() {
		connLock.Lock()
		delete(connMap, userID)
		connLock.Unlock()
		wsConn.Close()
		log.Println("User disconnected:", userID)
	}()

	for {
		if _, _, err := wsConn.ReadMessage(); err != nil {
			break
		}
	}
	return nil
}

package websocketsimpl

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ConnMap = make(map[string]*websocket.Conn)
var ConnLock = sync.Mutex{}

func (w *WebSocketsImpl) AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) error {
	ConnLock.Lock()
	ConnMap[userID] = wsConn
	ConnLock.Unlock()

	log.Println("User connected:", userID)

	defer func() {
		ConnLock.Lock()
		delete(ConnMap, userID)
		ConnLock.Unlock()
		wsConn.Close()
		log.Println("User disconnected:", userID)
	}()

	for {
		_, msgData, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("Error reading message", err)
			break
		}

		log.Println("Received message from", userID, string(msgData))
	}
	return nil
}

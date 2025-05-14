package websocketsimpl

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ConnMap = make(map[string]*websocket.Conn)
var ConnLock = sync.Mutex{}

func (w *WebSocketsImpl) AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) {
	ConnLock.Lock()
	ConnMap[userID] = wsConn
	ConnLock.Unlock()

	log.Println("co", ConnMap)

	// defer func() {
	// 	ConnLock.Lock()
	// 	delete(ConnMap, userID)
	// 	ConnLock.Unlock()
	// 	wsConn.Close()
	// 	log.Println("User disconnected:", userID)
	// }()

	for {
		_, _, err := wsConn.ReadMessage()
		if err != nil {
			delete(ConnMap, userID)
			wsConn.Close()
			log.Printf("Client %s disconnected", userID)
			break
		}
	}
}

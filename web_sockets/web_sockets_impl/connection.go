package websocketsimpl

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var connMap = make(map[string]*websocket.Conn)
var connLock = sync.Mutex{}

func (w *WebSocketsImpl) AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) error {
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
		_, msgBytes, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var msg models.Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		log.Printf("Message received from user %s to user %d: %s", userID, msg.ReceiverID, msg.Message)
		intUserID, _ := strconv.Atoi(userID)
		msg.SenderID = int64(intUserID)
		w.SendMessage(c, msg)
	}

	return nil
}

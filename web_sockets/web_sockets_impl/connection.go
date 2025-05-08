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

var ConnMap = make(map[string]*websocket.Conn)
var ConnLock = sync.Mutex{}

func (w *WebSocketsImpl) AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) error {
	ConnLock.Lock()
	ConnMap[userID] = wsConn
	ConnLock.Unlock()

	log.Println("User connected:", userID)

	userIDstr, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("Error: ", err)
	}

	pendingMessages, err := w.MessageService.GetUndeliveredMessages(int64(userIDstr))
	if err == nil {
		for _, m := range pendingMessages {
			msgResp := models.MessageResponse{
				SenderID:   m.SenderID,
				ReceiverID: m.ReceiverID,
				Message:    m.Message,
				Timestamp:  m.Timestamp,
			}

			msgBytes, _ := json.Marshal(msgResp)
			wsConn.WriteMessage(websocket.TextMessage, msgBytes)
		}
	}
	defer func() {
		ConnLock.Lock()
		delete(ConnMap, userID)
		ConnLock.Unlock()
		wsConn.Close()
		log.Println("User disconnected:", userID)
	}()

	for {
		_, _, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
	}

	return nil
}

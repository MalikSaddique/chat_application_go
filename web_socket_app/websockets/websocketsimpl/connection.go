package websocketsimpl

import (
	"log"
	"strconv"
	"sync"
	"time"

	messageserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/message_service/message_service_impl"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ConnLock = sync.Mutex{}

const (
	pingPeriod = 30 * time.Second
	pongWait   = 60 * time.Second
)

func (w *WebSocketsImpl) AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) error {
	uid, _ := strconv.Atoi(userID)
	ConnLock.Lock()
	if existingConn, ok := w.Clients[uid]; ok {
		existingConn.Close()
	}

	w.Clients[uid] = wsConn
	ConnLock.Unlock()

	log.Println("Conn", w.Clients)

	log.Println("User connected:", uid)

	defer func() {
		ConnLock.Lock()
		delete(w.Clients, uid)
		ConnLock.Unlock()
		wsConn.Close()
		log.Println("User disconnected:", uid)
	}()

	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()

		for range ticker.C {
			if err := wsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping error:", err)
				return
			}
		}
	}()

	for {
		var incoming messageserviceimpl.ServerMesageToSocket
		err := wsConn.ReadJSON(&incoming)
		if err != nil {
			log.Println("Error reading JSON:", err)
			break
		}

		log.Println("Received JSON from", userID, incoming)

		action := incoming.Action

		if action == "send" {
			// fmt.Println(47)
			receiverIDFloat := incoming.DestinationID
			receiverID := int(receiverIDFloat)

			message := incoming.Message
			if conn, ok := w.Clients[receiverID]; ok {

				err := conn.WriteJSON(map[string]any{
					"receiverID": receiverID,
					"message":    message,
				})

				if err != nil {
					log.Println("Error writing to receiver:", err)
				}
			} else {
				log.Println("Receiver not connected:", receiverID)
			}
		}
	}

	return nil
}

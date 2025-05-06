package websocketsimpl

import (
	"log"

	"github.com/MalikSaddique/chat_application_go/models"
)

func SendUser(userID string, msg models.Message) {
	connLock.Lock()
	conn, ok := connMap[userID]
	connLock.Unlock()

	if ok {
		if err := conn.WriteJSON(msg); err != nil {
			log.Println("Failed to send message to user:", userID, err)
		}
	}
}

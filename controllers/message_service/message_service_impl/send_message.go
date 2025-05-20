package messageserviceimpl

import (
	"log"

	"github.com/MalikSaddique/chat_application_go/models"
	websocketclient "github.com/MalikSaddique/chat_application_go/web_socket_client"
	"github.com/gin-gonic/gin"
)

type ServerMesageToSocket struct {
	Action        string
	DestinationID int
	Message       string
}

func (m *MessageServiceImpl) SendMessage(c *gin.Context, msg models.Message) error {

	err := m.MessageAuth.SaveMessage(c, msg)
	if err != nil {
		return err
	}

	messageToSend := ServerMesageToSocket{
		Action:        "send",
		DestinationID: int(msg.ReceiverID),
		Message:       msg.Message,
	}

	err = websocketclient.Conn.WriteJSON(messageToSend)
	if err != nil {
		log.Println("Error writing to WebSocket server:", err)
	}

	return nil
}

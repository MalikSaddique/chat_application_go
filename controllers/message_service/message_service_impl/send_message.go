package messageserviceimpl

import (
	"log"

	"github.com/MalikSaddique/chat_application_go/models"
	modelsWeb "github.com/MalikSaddique/socket/models"
	"github.com/gin-gonic/gin"
)

func (m *MessageServiceImpl) SendMessage(c *gin.Context, msg models.Message) error {
	err := m.MessageAuth.SaveMessage(c, msg)
	if err != nil {
		return err
	}
	msgRes := &modelsWeb.MessageRes{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Message:    msg.Message,
		Timestamp:  msg.Timestamp,
	}

	m.WebSocket.SendMessage(msgRes, c)
	log.Println("Message in msg service", msgRes)
	return nil
}

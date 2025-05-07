package messageserviceimpl

import (
	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (m *MessageServiceImpl) SendMessage(c *gin.Context, msg models.Message) error {
	err := m.MessageAuth.SaveMessage(c, msg)
	if err != nil {
		return err
	}
	return nil
}

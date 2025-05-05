package messageservice

import (
	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

type MessageService interface {
	SendMessage(senderID string, receiverID string, msg models.Message) error
	GetMessages(senderID string, receiverID string, pageStr string, limitStr string) ([]models.Message, error)
	UpdateMessage(c *gin.Context, idStr string, updatedMsg *models.Message) (*models.Message, error)
	DeleteMessage(c *gin.Context, idStr string) error
}

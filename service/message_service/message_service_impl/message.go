package messageserviceimpl

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (m *MessageServiceImpl) SendMessage(c *gin.Context) error {
	senderID := c.MustGet("userID").(string) // UserID extracted from token

	var msg models.SendMessageRequest
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil
	}

	err := m.UserAuth.SaveMessage(senderID, msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message"})
		return nil
	}

	return err

}

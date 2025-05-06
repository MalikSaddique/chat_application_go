package router

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) SendMessage(c *gin.Context) {

	var msg models.Message
	userID := c.MustGet("userID").(int64)

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	receiverID := msg.ReceiverID

	err := r.MessageService.SendMessage(userID, receiverID, msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access granted", "id": userID})
}

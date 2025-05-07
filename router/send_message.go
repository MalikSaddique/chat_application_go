package router

import (
	"net/http"
	"strconv"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) SendMessage(c *gin.Context) {

	var msg models.Message
	userIDstr := c.MustGet("userID").(string)
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// receiverID := msg.ReceiverID

	err = r.MessageService.SendMessage(c, msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access granted", "id": userID})
}

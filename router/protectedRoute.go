package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) SendMessage(c *gin.Context) {
	var msg models.Message
	userID := c.MustGet("userID").(string)

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	receiverID := strconv.FormatInt(msg.ReceiverID, 10)

	err := r.MessageService.SendMessage(userID, receiverID, msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access granted", "id": userID})
}

func (r *Router) GetMessages(c *gin.Context) {
	senderID := c.MustGet("userID").(string)
	receiverID := c.Query("receiver_id")

	if receiverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receiver_id is required"})
		return
	}

	messages, err := r.MessageService.GetMessages(senderID, receiverID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

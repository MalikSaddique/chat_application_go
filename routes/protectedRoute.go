package routes

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

// func (r *Router) GetMessage(c *gin.Context) {

// 	userID := c.MustGet("userID").(string)

// 	messages, err := h.messageService.FetchMessages(userID)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch messages"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, messages)

// }

func (r *Router) SendMessage(c *gin.Context) {
	// userID := c.MustGet("userID").(string)

	var msg models.SendMessageRequest
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := r.MessageService.SendMessage(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}

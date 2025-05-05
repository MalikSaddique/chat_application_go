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
	senderIDStr := c.Query("sender_id")
	receiverIDStr := c.Query("receiver_id")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	messages, err := r.MessageService.GetMessages(senderIDStr, receiverIDStr, pageStr, limitStr)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":     pageStr,
		"limit":    limitStr,
		"messages": messages,
	})
}

func (r *Router) UpdateMessage(c *gin.Context) {
	msgID := c.Param("_id")

	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	updatedMsg, err := r.MessageService.UpdateMessage(c, msgID, &msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message Updated Successfully": updatedMsg,
	})
}

func (r *Router) DeleteMessage(c *gin.Context) {
	msgID := c.Param("_id")

	err := r.MessageService.DeleteMessage(c, msgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

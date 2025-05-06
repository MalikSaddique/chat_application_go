package router

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

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

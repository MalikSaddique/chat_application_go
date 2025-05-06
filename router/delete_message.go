package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) DeleteMessage(c *gin.Context) {
	msgID := c.Param("_id")

	err := r.MessageService.DeleteMessage(c, msgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

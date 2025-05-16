package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteMessage godoc
// @Summary      Delete a message
// @Description  Delete a message by ID
// @Tags         messages
// @Produce      json
// @Param        _id  path      string  true  "Message ID"
// @Success      200
// @Failure      401
// @Failure      500
// @Security     BearerAuth
// @Router       /delete/{_id} [post]
func (r *Router) DeleteMessage(c *gin.Context) {
	msgID := c.Param("_id")

	err := r.MessageService.DeleteMessage(c, msgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

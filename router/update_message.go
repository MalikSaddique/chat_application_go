package router

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

// UpdateMessage godoc
// @Summary      Update a message
// @Description  Update the content of an existing message
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        _id      path      string         true  "Message ID"
// @Param        message  body      models.Message  true  "Updated message data"
// @Success      200
// @Failure      400
// @Failure      401
// @Failure      500
// @Security     BearerAuth
// @Router       /update/{_id} [get]
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

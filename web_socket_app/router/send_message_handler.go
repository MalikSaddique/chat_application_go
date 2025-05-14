package router

import (
	"net/http"
	"strconv"

	modelsWeb "github.com/MalikSaddique/socket/models"
	"github.com/gin-gonic/gin"
)

func (r *Router) SendMessageHTTPHandler(c *gin.Context) {
	receiverIDStr := c.Query("receiver_id")
	message := c.Query("message")

	receiverID, err := strconv.ParseInt(receiverIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver ID"})
		return
	}

	msg := &modelsWeb.MessageRes{
		ReceiverID: receiverID,
		Message:    message,
	}

	r.WebSocketSvc.SendMessage(msg)
	c.JSON(http.StatusOK, gin.H{"status": "message delivered"})
}

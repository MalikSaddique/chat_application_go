package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MalikSaddique/chat_application_go/models"
	websocketsimpl "github.com/MalikSaddique/chat_application_go/web_sockets/web_sockets_impl"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	msgResponse := models.MessageResponse{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Message:    msg.Message,
		Timestamp:  msg.Timestamp,
	}

	receiverID := strconv.FormatInt(msg.ReceiverID, 10)
	conn, ok := websocketsimpl.ConnMap[receiverID]

	if ok {
		msgBytes, err := json.Marshal(msgResponse)
		if err == nil {
			err = conn.WriteMessage(websocket.TextMessage, msgBytes)
			if err != nil {
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message sent", "from": userID, "to": msg.ReceiverID})

}

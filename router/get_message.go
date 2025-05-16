package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMessages godoc
// @Summary      Get messages
// @Description  Retrieve messages between two users
// @Tags         messages
// @Produce      json
// @Param        sender_id    query  string  true  "Sender ID"
// @Param        reciever_id  query  string  true  "Receiver ID"
// @Success      200          {array}   models.Message
// @Failure      400
// @Failure      500
// @Security     BearerAuth
// @Router       /getmessage [get]
func (r *Router) GetMessages(c *gin.Context) {

	senderIDStr := c.Query("sender_id")
	sID, err := strconv.ParseInt(senderIDStr, 10, 64)
	if err != nil {
		return
	}
	receiverIDStr := c.Query("receiver_id")
	rID, err := strconv.ParseInt(receiverIDStr, 10, 64)
	if err != nil {
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	messages, err := r.MessageService.GetMessages(sID, rID, pageStr, limitStr)
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

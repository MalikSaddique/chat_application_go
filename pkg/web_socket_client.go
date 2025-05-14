package pkg

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	modelsWeb "github.com/MalikSaddique/socket/models"
	"github.com/MalikSaddique/socket/websockets"
	"github.com/gin-gonic/gin"
)

type MessagePusher interface {
	SendMessage(msg *modelsWeb.MessageRes, c *gin.Context)
}

type WebSocketHTTPClient struct {
	ServerURL string
	WebSocket websockets.WebSockets
}

func (w *WebSocketHTTPClient) SendMessage(msg *modelsWeb.MessageRes, c *gin.Context) {
	userIDstr := c.MustGet("userID").(string)
	userID, err := strconv.ParseInt(userIDstr, 10, 64)
	endpoint := fmt.Sprintf("%s/protected/send?receiver_id=%s&message=%s",
		w.ServerURL,
		url.QueryEscape(strconv.FormatInt(userID, 10)),
		url.QueryEscape(msg.Message),
	)

	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println("Failed to send to WebSocket server:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("WebSocket server returned %s\n", resp.Status)
	}
}

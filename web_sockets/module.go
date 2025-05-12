package websockets

import (
	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSockets interface {
	AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) error
	SendMessage(msg *models.Message)
}

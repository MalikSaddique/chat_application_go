package websockets

import (
	modelsWeb "github.com/MalikSaddique/socket/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSockets interface {
	AddConn(userID string, wsConn *websocket.Conn, c *gin.Context)
	SendMessage(msg *modelsWeb.MessageRes)
}

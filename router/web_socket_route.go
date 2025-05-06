package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (r *Router) ConnectionUpgrade(c *gin.Context) error {
	userID := c.Query("userID")

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade failed:", err)
		return err
	}
	err, _ = r.WebSocketSvc.AddConn(userID, wsConn)

	return err

}

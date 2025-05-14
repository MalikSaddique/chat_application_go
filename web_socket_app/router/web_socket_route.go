package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (r *Router) StartWebSocketServer(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade failed:", err)
		return
	}
	go func() {
		r.WebSocketSvc.AddConn(userID, wsConn, c)
		// if err != nil {
		// 	log.Println("AddConn error:", err)
		// }
	}()
}

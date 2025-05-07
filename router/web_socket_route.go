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

func (r *Router) ConnectionUpgrade(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	log.Println("Route userid", userID)

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade failed:", err)
		return
	}

	err = r.WebSocketSvc.AddConn(userID, wsConn, c)
	if err != nil {
		log.Println("Error in AddConn:", err)
		return
	}
}

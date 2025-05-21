package router

import (
	"time"

	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	"github.com/MalikSaddique/socket/websockets"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine         *gin.Engine
	MessageService messageservice.MessageService
	WebSocketSvc   websockets.WebSockets
}

func NewRouter(messageService messageservice.MessageService, websocket websockets.WebSockets) *Router {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes := &Router{
		Engine:         engine,
		MessageService: messageService,
		WebSocketSvc:   websocket,
	}
	routes.defineWebSocketRouter()
	return routes
}

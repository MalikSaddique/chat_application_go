package router

import (
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	"github.com/MalikSaddique/socket/websockets"

	// websockets "github.com/MalikSaddique/chat_application_go/web_sockets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine         *gin.Engine
	MessageService messageservice.MessageService
	WebSocketSvc   websockets.WebSockets
}

func NewRouter(messageService messageservice.MessageService, websocket websockets.WebSockets) *Router {
	engine := gin.Default()
	routes := &Router{
		Engine:         engine,
		MessageService: messageService,
		WebSocketSvc:   websocket,
	}
	routes.defineWebSocketRouter()
	return routes
}

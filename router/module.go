package router

import (
	authservice "github.com/MalikSaddique/chat_application_go/controllers/auth_service"
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	websockets "github.com/MalikSaddique/chat_application_go/web_sockets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine         *gin.Engine
	AuthService    authservice.AuthService
	MessageService messageservice.MessageService
	WebSocketSvc   websockets.WebSockets
}

func NewRouter(authService authservice.AuthService, messageService messageservice.MessageService, webSocket websockets.WebSockets, onlyWs bool) *Router {
	engine := gin.Default()
	routes := &Router{
		Engine:         engine,
		AuthService:    authService,
		MessageService: messageService,
		WebSocketSvc:   webSocket,
	}
	if onlyWs {
		routes.defineWebSocketRouter()
	} else {
		routes.defineRoutes()
	}
	return routes
}

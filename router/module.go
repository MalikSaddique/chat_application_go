package router

import (
	authservice "github.com/MalikSaddique/chat_application_go/controllers/auth_service"
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine         *gin.Engine
	AuthService    authservice.AuthService
	MessageService messageservice.MessageService
}

func NewRouter(authService authservice.AuthService, messageService messageservice.MessageService) *Router {
	engine := gin.Default()
	routes := &Router{
		Engine:         engine,
		AuthService:    authService,
		MessageService: messageService,
	}
	routes.defineRoutes()
	return routes
}

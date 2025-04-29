package routes

import (
	authservice "github.com/MalikSaddique/chat_application_go/service/auth_service"
	messageservice "github.com/MalikSaddique/chat_application_go/service/message_service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService authservice.AuthService
	// UserService userservice.UserService
	MessageService messageservice.MessageService
}

func NewRouter(authService authservice.AuthService, messageService messageservice.MessageService) *Router {
	engine := gin.Default()
	router := &Router{
		Engine:      engine,
		AuthService: authService,
		// UserService: userService,
		MessageService: messageService,
	}
	router.defineRoutes()
	return router
}

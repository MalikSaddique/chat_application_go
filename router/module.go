package router

import (
	"time"

	authservice "github.com/MalikSaddique/chat_application_go/controllers/auth_service"
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine         *gin.Engine
	AuthService    authservice.AuthService
	MessageService messageservice.MessageService
}

func NewRouter(authService authservice.AuthService, messageService messageservice.MessageService) *Router {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes := &Router{
		Engine:         engine,
		AuthService:    authService,
		MessageService: messageService,
	}
	routes.defineRoutes()
	return routes
}

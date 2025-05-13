package router

import "github.com/MalikSaddique/chat_application_go/middleware"

func (r *Router) defineWebSocketRouter() {
	protected := r.Engine.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/ws", r.StartWebSocketServer)
	}
}

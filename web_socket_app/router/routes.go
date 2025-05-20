package router

import "github.com/MalikSaddique/chat_application_go/middleware"

func (r *Router) defineWebSocketRouter() {
	protected := r.Engine.Group("/protected")
	protected.Use(middleware.WSMiddleware())
	{
		protected.GET("/ws", r.StartWebSocketServer)
	}
	backend := r.Engine.Group("/backend")
	backend.Use(middleware.BackendWSMiddleware())
	{
		backend.GET("/ws", r.SaveBackendConnection)
	}

}

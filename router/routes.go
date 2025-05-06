package router

import (
	"github.com/MalikSaddique/chat_application_go/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) defineRoutes() {
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Engine.POST("/signup", r.SignUp)
	r.Engine.POST("/login", r.Login)
	r.Engine.GET("/refresh", r.RefreshKey)
	protected := r.Engine.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/send", r.SendMessage)
		protected.GET("/message", r.GetMessages)
		protected.GET("/update/:_id", r.UpdateMessage)
		protected.GET("/delete/:_id", r.DeleteMessage)
		protected.GET("/ws", r.ConnectionUpgrade)

	}
}

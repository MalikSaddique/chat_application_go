package routes

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) defineRoutes() {
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Engine.POST("/signup", r.SignUp)
	r.Engine.POST("/login", r.Login)
	r.Engine.GET("/protected", r.SendMessage)
	r.Engine.GET("/refresh", r.RefreshKey)
	// r.Engine.GET("getdata/:user_id", r.GetResult)
}

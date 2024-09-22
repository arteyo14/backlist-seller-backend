package routes

import (
	"blacklist-backend/controllers"
	"blacklist-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Header())

	//auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("logout", controllers.Logout)

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
	}

	return r
}

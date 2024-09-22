package routes

import (
	"blacklist-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/user", controllers.GetUsers)

	return r
}

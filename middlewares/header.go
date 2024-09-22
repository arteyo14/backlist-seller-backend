package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Header middleware is running") // เพิ่ม log เพื่อตรวจสอบ
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

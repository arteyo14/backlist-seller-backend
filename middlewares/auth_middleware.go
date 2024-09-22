package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ฟังก์ชันการตรวจสอบ JWT Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"error":  gin.H{"message": "Authorization header is required"},
				"status": false,
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// ตรวจสอบ token ที่นี่ (เพิ่มการตรวจสอบจริงหากต้องการ)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"error":  gin.H{"message": "Invalid token"},
				"status": false,
			})
			c.Abort()
			return
		}

		// หากทุกอย่างถูกต้อง
		c.Next()
	}
}

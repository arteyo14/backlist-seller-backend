package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // อนุญาตทุก Origin สามารถเปลี่ยนเป็น domain ที่ต้องการ
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// จัดการกับ OPTIONS method
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// ดำเนินการต่อกับคำขอ
		c.Next()
	}
}

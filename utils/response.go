package utils

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, code int, status bool, data interface{}) {
	response := gin.H{
		"status": status,
		"code":   code,
	}

	if message, ok := data.(string); ok {
		response["message"] = message
	} else {
		response["data"] = data
	}

	c.JSON(code, response)
}

func JSONErrorResponse(c *gin.Context, code int, err string) {
	if err != "" {
		JSONResponse(c, code, false, gin.H{"message": err})
	}
}

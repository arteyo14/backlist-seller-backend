package controllers

import (
	"blacklist-backend/config"
	"blacklist-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "code": http.StatusInternalServerError, "message": err})
		return
	}

	c.JSON(http.StatusOK, users)
}

package controllers

import (
	"blacklist-backend/config"
	"blacklist-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	// ตรวจสอบการเชื่อมต่อฐานข้อมูล
	if config.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error": gin.H{
				"message": "Database connection is not established",
			},
		})
		return
	}

	// ตรวจสอบว่าตาราง users มีอยู่ในฐานข้อมูลหรือไม่
	if !config.DB.Migrator().HasTable(&models.User{}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error": gin.H{
				"message": "Table 'users' does not exist in the database",
			},
		})
		return
	}

	// ดึงข้อมูลผู้ใช้ทั้งหมด
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error": gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"code":   http.StatusOK,
		"data":   users,
	})
}

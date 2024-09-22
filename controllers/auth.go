package controllers

import (
	"blacklist-backend/config"
	"blacklist-backend/models"
	"blacklist-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("UpJ2g,r9H/:>qERc!hn`SK")

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	// ตรวจสอบว่า JSON ที่ส่งมาถูกต้องหรือไม่
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	// เช็คว่ามี username อยู่แล้วหรือไม่
	var existingUser models.User
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		// พบผู้ใช้ที่มี username นี้อยู่แล้ว
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"code":   http.StatusConflict,
			"error": gin.H{
				"message": "Username already exists",
			},
		})
		return
	}

	// เช็คว่ามี email อยู่แล้วหรือไม่
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		// พบผู้ใช้ที่มี email นี้อยู่แล้ว
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"code":   http.StatusConflict,
			"error": gin.H{
				"message": "Email already exists",
			},
		})
		return
	}

	// เข้ารหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error": gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	// สร้างผู้ใช้ใหม่
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Email:    input.Email,
	}

	// บันทึกผู้ใช้ใหม่ลงในฐานข้อมูล
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error": gin.H{
				"meesage": err.Error(),
			},
		})
		return
	}

	// ตอบกลับเมื่อการลงทะเบียนสำเร็จ
	utils.JSONResponse(c, http.StatusOK, true, "Register Successfully")
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	var user models.User

	//เช็ค username จาก DB
	if err := config.DB.Where("username =?", input.Username).First(&user).Error; err != nil {
		utils.JSONErrorResponse(c, http.StatusNotFound, "Invalid username or password")
		return
	}

	//ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.JSONErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	//สร้าง JWT
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString((jwtSecretKey))
	if err != nil {
		utils.JSONErrorResponse(c, http.StatusInternalServerError, "Failed to create token ")
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, gin.H{"token": tokenString})
}

func Logout(c *gin.Context) {
	// ลบ Token ใน Cookie (ถ้ามีการเก็บไว้ใน Cookie)
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"code":    http.StatusOK,
		"message": "Logged out successfully!",
	})
}

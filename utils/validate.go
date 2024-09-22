package utils

import (
	"regexp"
	"unicode"
)

func IsValidUsername(username string) bool {
	// ตรวจสอบความยาวและตัวอักษรที่ใช้ได้
	if len(username) < 3 || len(username) > 50 {
		return false
	}
	// ตรวจสอบว่า username ประกอบด้วยตัวอักษร, ตัวเลข, _ และ - เท่านั้น
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return validUsername.MatchString(username)
}

// ฟังก์ชันตรวจสอบความถูกต้องของ Password
func IsValidPassword(password string) bool {
	var (
		hasMinLen      = false
		hasUpper       = false
		hasLower       = false
		hasNumber      = false
		hasSpecialChar = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecialChar
}

// ฟังก์ชันตรวจสอบความถูกต้องของ Email
func IsValidEmail(email string) bool {
	// ใช้ regex สำหรับตรวจสอบรูปแบบอีเมลที่ถูกต้อง
	validEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return validEmail.MatchString(email)
}

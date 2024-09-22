package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=blacklist-seller port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = database

	// ใช้ log.Println แทน log.Fatal เพื่อแสดงข้อความโดยไม่หยุดโปรแกรม
	log.Println("Successfully connected to database")
}

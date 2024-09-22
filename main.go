package main

import (
	"blacklist-backend/config"
	"blacklist-backend/middlewares"
	"blacklist-backend/models"
	"blacklist-backend/routes"
	"log"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRoutes()

	if config.DB == nil {
		log.Fatal("Database connection is not established")
	}

	err := config.DB.AutoMigrate(&models.User{}, &models.Seller{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully.")

	r.Use(middlewares.Header())

	r.Run(":8080")
}

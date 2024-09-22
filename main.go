package main

import (
	"blacklist-backend/middlewares"
	"blacklist-backend/routes"
)

func main() {
	r := routes.SetupRoutes()

	r.Use(middlewares.Header())

	r.Run(":8080")
}

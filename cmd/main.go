package main

import (
	"github.com/agungputrap/devconnect-backend/config"
	"github.com/agungputrap/devconnect-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.ConnectDatabase()
	app := fiber.New()
	routes.SetupRoutes(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to start application: ", err)
	}
}

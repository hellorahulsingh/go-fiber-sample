package main

import (
	"go-fiber-app/internal/config"
	"go-fiber-app/internal/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load environment config
	config.LoadEnv()

	// connect to MongoDb
	config.ConnectMongo()

	// Register routes
	route.SetupRoutes(app)

	app.Listen(":3000")
}


package main

import (
	"resume_builder_backend/src/pkg/configs"
	"resume_builder_backend/src/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title User API documentation
// @version 1.0.0
// @host localhost:3000
func main() {
	initialize()

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3001")
}

func initialize() {
	configs.SetupEnv()
	database.SetupPostgres()
}

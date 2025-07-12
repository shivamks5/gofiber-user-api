package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivamks5/gofiber-user-api/routes"
)

func main() {
	app := fiber.New()
	routes.SetupUserRoutes(app)
	app.Listen(":3000")
}

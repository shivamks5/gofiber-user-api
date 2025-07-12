package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivamks5/gofiber-user-api/handler"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", handler.GetUsers)
	api.Post("/users", handler.CreateUser)
	api.Get("/users/:id", handler.GetUserByID)
	api.Put("/users/:id", handler.UpdateUser)
	api.Delete("/users/:id", handler.DeleteUser)
}

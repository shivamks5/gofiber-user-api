package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shivamks5/gofiber-user-api/model"
)

var users []model.User = []model.User{}

func CreateUser(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json request",
		})
	}
	user.ID = uuid.New().String()
	users = append(users, user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, user := range users {
		if user.ID == id {
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "user not found",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedUser model.User
	err := c.BodyParser(&updatedUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json request",
		})
	}
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = user.ID
			users[i] = updatedUser
			return c.JSON(updatedUser)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "user not found",
	})
}

func PatchUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var patchedUser map[string]interface{}
	err := c.BodyParser(&patchedUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json request",
		})
	}
	for i, user := range users {
		if user.ID == id {
			if name, ok := patchedUser["name"].(string); ok {
				user.Name = name
			}
			if email, ok := patchedUser["email"].(string); ok {
				user.Email = email
			}
			if age, ok := patchedUser["age"].(float64); ok {
				user.Age = int(age)
			}
			users[i] = user
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "user not found",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(fiber.Map{
				"message": "user deleted successfully",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "user not found",
	})
}

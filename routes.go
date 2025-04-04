package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Server is running",
		})
	})

	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Get("/", handleGetUsers)
	users.Get("/:id", handleGetUser)
	users.Post("/", handleCreateUser)
	users.Put("/:id", handleUpdateUser)
	users.Delete("/:id", handleDeleteUser)

}

func handleGetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get all users"})
}

func handleGetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get user by ID: " + c.Params("id")})
}

func handleCreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create user"})
}

func handleUpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update user: " + c.Params("id")})
}

func handleDeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete user: " + c.Params("id")})
}

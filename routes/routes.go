package routes

import (
	"go-postgre/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userController := controllers.UserController{}

	// Welcome route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API")
	})

	// Users route with pagination support
	app.Get("/users", func(c *fiber.Ctx) error {
		// Call the controller method to fetch the list of users with pagination
		err := userController.GetUsers(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return nil
	})
}

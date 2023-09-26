package controllers

import (
	"go-postgre/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserModel *models.UserModel
}

// GetUsers retrieves a list of users with pagination.
func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	// Get the page and limit parameters from the query string.
	page := ctx.QueryInt("page", 1)

	limit := ctx.QueryInt("limit", 100)

	// Call the FindAll method from the UserModel with the page and limit parameters.
	users, err := c.UserModel.FindAll(page, limit)
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return the list of users as JSON response.
	return ctx.JSON(users)
}

// Implement controller methods for CRUD operations here.

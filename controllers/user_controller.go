package controllers

import (
	"go-postgre/models"
)

type UserController struct {
	UserModel *models.UserModel
}

// GetUsers retrieves a list of all users.
func (c *UserController) GetUsers() ([]models.User, error) {
	// Call the FindAll method from the UserModel
	users, err := c.UserModel.FindAll()
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		return nil, err
	}
	return users, nil
}

// Implement controller methods for CRUD operations here.

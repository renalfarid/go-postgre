package models

import (
	"database/sql"
)

var db *sql.DB

type User struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
}

// UserModel represents the User model and its database operations.
type UserModel struct{}

// Create inserts a new user into the database.
func (m *UserModel) Create(user *User) error {
	// Implement the insert logic here using db.Exec
	return nil
}

// FindAll fetches all users from the database.
func (m *UserModel) FindAll() ([]User, error) {
	// Implement the fetch all users logic here using db.Query
	// Query the database to retrieve all users
	db = ConnectDb()
	rows, err := db.Query("SELECT id, user_name, email FROM users")
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			// Handle the error, e.g., log it or return an error response
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		// Handle the error, e.g., log it or return an error response
		return nil, err
	}

	return users, nil
}

// FindByID retrieves a user by ID from the database.
func (m *UserModel) FindByID(id int) (*User, error) {
	// Implement the fetch user by ID logic here using db.QueryRow
	return nil, nil
}

// Update updates an existing user in the database.
func (m *UserModel) Update(user *User) error {
	// Implement the update logic here using db.Exec
	return nil
}

// Delete removes a user from the database by ID.
func (m *UserModel) Delete(id int) error {
	// Implement the delete logic here using db.Exec
	return nil
}

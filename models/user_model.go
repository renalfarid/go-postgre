package models

import (
	"database/sql"
	"encoding/json"
)

var db *sql.DB

type User struct {
	ID       int            `json:"id"`
	Username sql.NullString `json:"username"`
	Email    sql.NullString `json:"email"`
}

// UserModel represents the User model and its database operations.
type UserModel struct{}

// MarshalJSON implements custom JSON marshaling for the User struct.
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	if u.Username.Valid {
		return json.Marshal(&struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{
			ID:       u.ID,
			Username: u.Username.String,
			Email:    u.Email.String,
		})
	} else {
		return json.Marshal(&struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{
			ID:       u.ID,
			Username: "",
			Email:    "",
		})
	}
}

// Create inserts a new user into the database.
func (m *UserModel) Create(user *User) error {
	// Implement the insert logic here using db.Exec
	return nil
}

// FindAll fetches users from the database with pagination support.
func (m *UserModel) FindAll(page, limit int) ([]User, error) {
	db = ConnectDb()
	// Calculate the offset based on the page and limit values.
	offset := (page - 1) * limit

	// Define the SQL query for fetching users with pagination.
	query := `
		SELECT id, username, email
		FROM users
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	// Execute the SQL query with the specified limit and offset.
	rows, err := db.Query(query, limit, offset)
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

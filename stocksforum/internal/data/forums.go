// Filename: internal/data/forums.go

package data

import (
	"database/sql"
	"time"

	"stocksforum.renesanchez.net/internal/validator"
)

type Forum struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
	Version   int32     `json:"version"`
}

func ValidateForum(v *validator.Validator, forum *Forum) {
	// Use the Check() method to execute our validation checks
	v.Check(forum.Name != "", "name", "must be provided")
	v.Check(len(forum.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(forum.Message != "", "message", "must be provided")
	v.Check(len(forum.Message) <= 2000, "message", "must not be more than 2000 bytes long")
}

// Define a ForumModel which wraps a sql.DB connection pool
type ForumModel struct {
	DB *sql.DB
}

// Insert() allows us  to create a new Forum
func (m ForumModel) Insert(forum *Forum) error {
	query := `
		INSERT INTO forums (name, message)
		VALUES ($1, $2)
		RETURNING id, created_at, version
	`
	// Collect the data fields into a slice
	args := []interface{}{
		forum.Name, forum.Message,
	}
	return m.DB.QueryRow(query, args...).Scan(&forum.ID, &forum.CreatedAt, &forum.Version)
}

// Get() allows us to retrieve a specific Forum
func (m ForumModel) Get(id int64) (*Forum, error) {
	return nil, nil
}

// Update() allows us to edit/alter a specific Forum
func (m ForumModel) Update(forum *Forum) error {
	return nil
}

// Delete() removes a specific Forum
func (m ForumModel) Delete(id int64) error {
	return nil
}

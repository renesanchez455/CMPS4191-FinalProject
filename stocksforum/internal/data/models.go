// Filename: internal/data/models.go

package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// A wrapper for our data models
type Models struct {
	Forums ForumModel
	Users  UserModel
}

// NewModels() allows us to create a new Models
func NewModels(db *sql.DB) Models {
	return Models{
		Forums: ForumModel{DB: db},
		Users:  UserModel{DB: db},
	}
}

// Filename: internal/data/forums.go

package data

import (
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

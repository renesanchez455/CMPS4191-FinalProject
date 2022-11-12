package main

import (
	"fmt"
	"net/http"
	"time"

	"stocksforum.renesanchez.net/internal/data"
	"stocksforum.renesanchez.net/internal/validator"
)

// createForumHandler for the "POST /v1/forums" endpoint
func (app *application) createForumHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination
	var input struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}
	// Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Forum struct
	forum := &data.Forum{
		Name:    input.Name,
		Message: input.Message,
	}

	// Initialize a new Validator instance
	v := validator.New()

	// Check the map to determine if there were any validation errors
	if data.ValidateForum(v, forum); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Create a Forum
	err = app.models.Forums.Insert(forum)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// Create a Location header for the newly created resource/Forum
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/forums/%d", forum.ID))
	// Write the JSON response with 201 - Created status code with the body
	// being the Forum data and the header being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"forum": forum}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// showForumHandler for the "GET /v1/forums/:id" endpoint
func (app *application) showForumHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the Forum struct containing the ID we extracted
	// from our URL and some sample data
	forum := data.Forum{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "The stock market",
		Message:   "This is a post about the stock market.",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"forum": forum}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

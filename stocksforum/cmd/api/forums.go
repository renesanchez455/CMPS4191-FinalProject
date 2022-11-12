package main

import (
	"errors"
	"fmt"
	"net/http"

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

	// Fetch the specific forum
	forum, err := app.models.Forums.Get(id)
	// Handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"forum": forum}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

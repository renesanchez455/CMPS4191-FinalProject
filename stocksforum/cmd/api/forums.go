package main

import (
	"fmt"
	"net/http"
	"time"

	"stocksforum.renesanchez.net/internal/data"
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
	// Display the request
	fmt.Fprintf(w, "%+v\n", input)
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

package main

import (
	"fmt"
	"net/http"
)

// createSchoolHandler for the "POST /v1/schools" endpoint
func (app *application) createForumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new forum..")
}

// showSchoolHandler for the "GET /v1/schools/:id" endpoint
func (app *application) showForumHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Display the school id
	fmt.Fprintf(w, "show the details for forum %d\n", id)
}

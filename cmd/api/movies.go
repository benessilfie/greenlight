package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.essilfie.co.uk/internal/data"
)

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovie(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		Title:     "Avengers Endgame",
		Year:      2019,
		Genres:    []string{"action", "science fiction", "drama", "adventure"},
		Runtime:   182,
		Version:   1,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

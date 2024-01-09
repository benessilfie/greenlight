package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheck)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.listMovies)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovie)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovie)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovie)

	return app.recoverPanic(router)
}

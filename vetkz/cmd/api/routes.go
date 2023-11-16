package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/cat", app.listCatsHandler)

	router.HandlerFunc(http.MethodGet, "/v1/catalog", app.catalogHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cat/:id", app.showCatHandler)
	router.HandlerFunc(http.MethodPost, "/v1/cat", app.createCatHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/cat/:id", app.updateCatHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteCatHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))

}

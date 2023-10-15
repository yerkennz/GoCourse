package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/catalog", app.catalogHandler)
	router.HandlerFunc(http.MethodGet, "/v1/cat/:id", app.showCatHandler)

	return router
}

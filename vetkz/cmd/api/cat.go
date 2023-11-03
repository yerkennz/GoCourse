package main

import (
	"fmt"
	"net/http"
	"time"
	"vetkz.yerkennz.net/internal/data"
)

func (app *application) createCatHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title         string `json:"title"`
		Product       string `json:"product"`
		Price         int64  `json:"price"`
		AgeCat        string `json:"age_cat"`
		SizeCat       string `json:"size_cat"`
		Breed         string `json:"breed"`
		CountryOrigin string `json:"country_origin"`
		Description   string `json:"description"`
		Quantity      int64  `json:"quantity"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showCatHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}
	cat := data.Cat{
		ID:            id,
		CreatedAt:     time.Now(),
		Title:         "СУХОЙ КОРМ",
		Price:         44410,
		Product:       "Corm",
		AgeCat:        "Взрослые (1 - 7 лет)",
		SizeCat:       "Породы любого размера",
		Breed:         "Любая порода",
		CountryOrigin: "Россия",
		Description:   "Void",
		Quantity:      64,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"cat": cat}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

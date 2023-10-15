package main

import (
	"fmt"
	"net/http"
	"time"
	"vetkz.yerkennz.net/internal/data"
)

func (app *application) createCatHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title             string   `json:"title"`
		Product           string   `json:"product"`
		Packing           []string `json:"packing"`
		Price             int64    `json:"price"`
		TypePreparation   string   `json:"type_preparation"`
		TypeFeed          string   `json:"type_feed"`
		AgeCat            string   `json:"age_cat"`
		SizeCat           string   `json:"size_cat"`
		ActivityLevel     string   `json:"activity_level"`
		Breed             string   `json:"breed"`
		TypeProtection    string   `json:"type_protection"`
		SpecialIndication string   `json:"special_indication"`
		Taste             []string `json:"taste"`
		TypeTool          string   `json:"type_tool"`
		CountryOrigin     string   `json:"country_origin"`
		Description       string   `json:"description"`
		Quantity          int64    `json:"quantity"`
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
		ID:                id,
		CreatedAt:         time.Now(),
		Title:             "СУХОЙ КОРМ",
		Price:             44410,
		Product:           "Corm",
		TypeFeed:          "Сухой корм",
		AgeCat:            "Взрослые (1 - 7 лет)",
		SizeCat:           "Породы любого размера",
		ActivityLevel:     "Нормальный",
		Breed:             "Любая порода",
		SpecialIndication: "Кастрация и стерилизация",
		CountryOrigin:     "Россия",
		Description:       "Void",
		Quantity:          64,
		Taste:             []string{"Кролик", "Лосось", "Индейка"},
		Packing:           []string{"400 гр", "1,5 кг"},
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"cat": cat}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

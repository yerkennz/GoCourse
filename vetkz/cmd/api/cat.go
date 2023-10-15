package main

import (
	"net/http"
	"time"
	"vetkz.yerkennz.net/internal/data"
)

func (app *application) showCatHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
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
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

package main

import (
	"fmt"
	_ "github.com/lib/pq"
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

	cat := &data.Cat{
		Title:         input.Title,
		Product:       input.Product,
		Price:         input.Price,
		AgeCat:        input.AgeCat,
		SizeCat:       input.SizeCat,
		Breed:         input.Breed,
		CountryOrigin: input.CountryOrigin,
		Description:   input.Description,
		Quantity:      input.Quantity,
	}
	//v := validator.New()
	//if data.ValidateMovie(v, movie); !v.Valid() {
	//	app.failedValidationResponse(w, r, v.Errors)
	//	return
	//}
	// Call the Insert() method on our movies model, passing in a pointer to the
	// validated movie struct. This will create a record in the database and update the
	// movie struct with the system-generated information.
	err = app.models.Cats.Insert(cat)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// When sending a HTTP response, we want to include a Location header to let the
	// client know which URL they can find the newly-created resource at. We make an
	// empty http.Header map and then use the Set() method to add a new Location header,
	// interpolating the system-generated ID for our new movie in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/cat/%d", cat.ID))
	// Write a JSON response with a 201 Created status code, the movie data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"cat": cat}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

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

package main

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
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
	cat, err := app.models.Cats.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"cat": cat}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateCatHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Fetch the existing movie record from the database, sending a 404 Not Found
	// response to the client if we couldn't find a matching record.
	cat, err := app.models.Cats.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Declare an input struct to hold the expected data from the client.
	var input struct {
		Title       *string `json:"title"`
		Product     *string `json:"product"`
		Price       *int64  `json:"price"`
		Description *string `json:"description"`
		Quantity    *int64  `json:"quantity"`
	}
	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the request body to the appropriate fields of the movie
	// record.
	if input.Title != nil {
		cat.Title = *input.Title
	}
	if input.Product != nil {
		cat.Product = *input.Product
	}
	if input.Price != nil {
		cat.Price = *input.Price
	}
	if input.Description != nil {
		cat.Description = *input.Description
	}
	if input.Quantity != nil {
		cat.Quantity = *input.Quantity
	}

	// Validate the updated movie record, sending the client a 422 Unprocessable Entity
	// response if any checks fail.
	//v := validator.New()
	//if data.ValidateMovie(v, movie); !v.Valid() {
	//app.failedValidationResponse(w, r, v.Errors)
	//return
	//}
	// Pass the updated movie record to our new Update() method.
	err = app.models.Cats.Update(cat)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Write the updated movie record in a JSON response.
	err = app.writeJSON(w, http.StatusOK, envelope{"cat": cat}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteCatHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Delete the movie from the database, sending a 404 Not Found response to the
	// client if there isn't a matching record.
	err = app.models.Cats.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listCatsHandler(w http.ResponseWriter, r *http.Request) {
	// To keep things consistent with our other handlers, we'll define an input struct
	// to hold the expected values from the request query string.

	var input struct {
		Title       string
		Product     string
		Description string
		data.Filters
	}
	// Initialize a new Validator instance.
	// Call r.URL.Query() to get the url.Values map containing the query string data.
	qs := r.URL.Query()
	// Use our helpers to extract the title and genres query string values, falling back
	// to defaults of an empty string and an empty slice respectively if they are not
	// provided by the client.
	input.Title = app.readString(qs, "title", "")
	input.Product = app.readString(qs, "product", "")
	// Get the page and page_size query string values as integers. Notice that we set
	// the default page value to 1 and default page_size to 20, and that we pass the
	// validator instance as the final argument here.
	input.Filters.Price = app.readInt(qs, "price", 1)
	input.Description = app.readString(qs, "product", "")
	input.Filters.Quantity = app.readInt(qs, "quantity", 20)
	// Extract the sort query string value, falling back to "id" if it is not provided
	// by the client (which will imply a ascending sort on movie ID).
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "-id", "-title"}

	// Check the Validator instance for any errors and use the failedValidationResponse()
	// helper to send the client a response if necessary.
	//if !v.Valid() {
	//	app.failedValidationResponse(w, r, v.Errors)
	//	return
	//}
	// Dump the contents of the input struct in a HTTP response.
	cat, err := app.models.Cats.GetAll(input.Title, input.Product, input.Description, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"cats": cat}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

package data

import (
	"database/sql"
	"errors"
	"time"
)

type Cat struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"-"`
	Title         string    `json:"title"`
	Product       string    `json:"product"`
	Price         int64     `json:"price"`
	AgeCat        string    `json:"age_cat,omitempty"`
	SizeCat       string    `json:"size_cat,omitempty"`
	Breed         string    `json:"breed,omitempty"`
	CountryOrigin string    `json:"country_origin"`
	Description   string    `json:"description,omitempty"`
	Quantity      int64     `json:"quantity"`
}

type CatModel struct {
	DB *sql.DB
}

func (m CatModel) Get(id int64) (*Cat, error) {
	// The PostgreSQL bigserial type that we're using for the movie ID starts
	// auto-incrementing at 1 by default, so we know that no movies will have ID values
	// less than that. To avoid making an unnecessary database call, we take a shortcut
	// and return an ErrRecordNotFound error straight away.
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the movie data.
	query := `
			SELECT id, title, product, price, quantity
			FROM cats
			WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var cat Cat
	// Execute the query using the QueryRow() method, passing in the provided id value
	// as a placeholder parameter, and scan the response data into the fields of the
	// Movie struct. Importantly, notice that we need to convert the scan target for the
	// genres column using the pq.Array() adapter function again.
	err := m.DB.QueryRow(query, id).Scan(
		&cat.ID,
		&cat.Title,
		&cat.Product,
		&cat.Price,
		&cat.Quantity,
	)
	// Handle any errors. If there was no matching movie found, Scan() will return
	// a sql.ErrNoRows error. We check for this and return our custom ErrRecordNotFound
	// error instead.
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Otherwise, return a pointer to the Movie struct.
	return &cat, nil
}

// Add a placeholder method for inserting a new record in the movies table.
func (m CatModel) Insert(cat *Cat) error {
	query := `
			INSERT INTO cats (title,product, price, age_cat, size_cat, breed, country_origin, description, quantity)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id, created_at, title`
	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []interface{}{cat.Title, cat.Product, cat.Price, cat.AgeCat, cat.SizeCat, cat.Breed, cat.CountryOrigin, cat.Description, cat.Quantity}
	// Use the QueryRow() method to execute the SQL query on our connection pool,
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the movie struct.
	return m.DB.QueryRow(query, args...).Scan(&cat.ID, &cat.CreatedAt, &cat.Title)
}

// Add a placeholder method for fetching a specific record from the movies table.

// Add a placeholder method for updating a specific record in the movies table.
func (m CatModel) Update(cat *Cat) error {
	query := `
			UPDATE cats
			SET title = $1, product = $2, price = $3, description = $4, quantity = $5
			WHERE id = $6
			RETURNING title`
	// Create an args slice containing the values for the placeholder parameters.
	args := []interface{}{
		cat.Title,
		cat.Product,
		cat.Price,
		cat.Description,
		cat.Quantity,
		cat.ID,
	}
	// Use the QueryRow() method to execute the query, passing in the args slice as a
	// variadic parameter and scanning the new version value into the movie struct.
	return m.DB.QueryRow(query, args...).Scan(&cat.Title)
}

// Add a placeholder method for deleting a specific record from the movies table.
func (m CatModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	// Construct the SQL query to delete the record.
	query := `
			DELETE FROM cats
			WHERE id = $1`
	// Execute the SQL query using the Exec() method, passing in the id variable as
	// the value for the placeholder parameter. The Exec() method returns a sql.Result
	// object.
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	// Call the RowsAffected() method on the sql.Result object to get the number of rows
	// affected by the query.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// If no rows were affected, we know that the movies table didn't contain a record
	// with the provided ID at the moment we tried to delete it. In that case we
	// return an ErrRecordNotFound error.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil

}

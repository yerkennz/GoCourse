package data

import (
	"database/sql"
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
func (m CatModel) Get(id int64) (*Cat, error) {
	return nil, nil
}

// Add a placeholder method for updating a specific record in the movies table.
func (m CatModel) Update(movie *Cat) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the movies table.
func (m CatModel) Delete(id int64) error {
	return nil
}

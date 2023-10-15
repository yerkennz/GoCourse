package data

import "time"

type Cat struct {
	ID                int64     `json:"id"`
	CreatedAt         time.Time `json:"-"`
	Title             string    `json:"title"`
	Product           string    `json:"product"`
	Packing           []string  `json:"packing"`
	Price             int64     `json:"price"`
	TypePreparation   string    `json:"type_preparation"`
	TypeFeed          string    `json:"type_feed,omitempty"`
	AgeCat            string    `json:"age_cat,omitempty"`
	SizeCat           string    `json:"size_cat,omitempty"`
	ActivityLevel     string    `json:"activity_level,omitempty"`
	Breed             string    `json:"breed,omitempty"`
	TypeProtection    string    `json:"type_protection,omitempty"`
	SpecialIndication string    `json:"special_indication,omitempty"`
	Taste             []string  `json:"taste,omitempty"`
	TypeTool          string    `json:"type_tool"`
	CountryOrigin     string    `json:"country_origin"`
	Description       string    `json:"description,omitempty"`
	Quantity          int64     `json:"quantity"`
}

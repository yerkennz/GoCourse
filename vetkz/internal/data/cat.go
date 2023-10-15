package data

import "time"

type Cat struct {
	ID                int64
	CreatedAt         time.Time
	Title             string
	Product           string
	Packing           []string
	Price             int64
	TypePreparation   string
	TypeFeed          string
	AgeCat            string
	SizeCat           string
	ActivityLevel     string
	Breed             string
	TypeProtection    string
	SpecialIndication string
	Taste             []string
	TypeTool          string
	CountryOrigin     string
	Description       string
	Quantity          int64
}

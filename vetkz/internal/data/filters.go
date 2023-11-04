package data

import (
	"strings"
	"vetkz.yerkennz.net/internal/validator"
)

type Filters struct {
	Price        int
	Quantity     int
	Sort         string
	SortSafelist []string
}

func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

// Return the sort direction ("ASC" or "DESC") depending on the prefix character of the
// Sort field.
func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filters) limit() int {
	return f.Price
}
func (f Filters) offset() int {
	return (f.Price - 1) * f.Quantity
}

func ValidateFilters(v *validator.Validator, f Filters) {
	// Check that the page and page_size parameters contain sensible values.
	v.Check(f.Price > 0, "price", "must be greater than zero")
	v.Check(f.Price <= 10_000_000, "price", "must be a maximum of 10 million")
	v.Check(f.Quantity > 0, "quantity", "must be greater than zero")
	v.Check(f.Quantity <= 100, "quantity", "must be a maximum of 100")
	// Check that the sort parameter matches a value in the safelist.
	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}

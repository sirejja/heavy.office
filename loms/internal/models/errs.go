package models

import "errors"

var (
	ErrEmptySKU     = errors.New("empty sku")
	ErrEmptyUser    = errors.New("empty user")
	ErrEmptyCount   = errors.New("empty count")
	ErrEmptyOrderID = errors.New("empty orderID")

	ErrNoFiltersProvided  = errors.New("No filters provided")
	ErrNotFound           = errors.New("There is no row for provided filters")
	ErrNoDataProvided     = errors.New("No data provided")
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

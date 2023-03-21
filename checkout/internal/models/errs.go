package models

import "errors"

var (
	ErrEmptySKU           = errors.New("empty sku")
	ErrEmptyUser          = errors.New("empty user")
	ErrEmptyCount         = errors.New("empty count")
	ErrInsufficientStocks = errors.New("insufficient stocks_handler")

	ErrDBEmptySKU      = errors.New("empty sku in db query")
	ErrInsertFailed    = errors.New("Insert failed")
	ErrNothingToDelete = errors.New("User has no such products in cart")
	ErrNoDataProvided  = errors.New("No data provided")
)

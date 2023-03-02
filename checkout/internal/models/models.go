package models

import "github.com/pkg/errors"

var (
	ErrEmptySKU           = errors.New("empty sku")
	ErrEmptyUser          = errors.New("empty user")
	ErrEmptyCount         = errors.New("empty count")
	ErrInsufficientStocks = errors.New("insufficient stocks_handler")
)

type Product struct {
	SKU   uint32
	Count uint16
	Name  string
	Price uint32
}

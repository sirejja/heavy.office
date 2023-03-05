package models

import "errors"

var (
	ErrEmptySKU     = errors.New("empty sku")
	ErrEmptyUser    = errors.New("empty user")
	ErrEmptyCount   = errors.New("empty count")
	ErrEmptyOrderID = errors.New("empty orderID")
)

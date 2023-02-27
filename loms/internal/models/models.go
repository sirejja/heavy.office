package models

import "github.com/pkg/errors"

var (
	ErrEmptyOrder = errors.New("empty orders")
	ErrEmptySKU   = errors.New("empty sku")
	ErrEmptyUser  = errors.New("empty user")
	ErrEmptyCount = errors.New("empty count")
)

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Stock struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

package stocks_handler

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
)

type IService interface {
	Stocks(ctx context.Context, SKU uint32) (*[]models.Stock, error)
}
type Handler struct {
	warehouse IService
}

func New(warehouse IService) *Handler {
	return &Handler{
		warehouse: warehouse,
	}
}

type Request struct {
	SKU uint32 `json:"sku"`
}

func (r Request) Validate() error {
	if r.SKU == 0 {
		return models.ErrEmptySKU
	}
	return nil
}

type Item struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type Response struct {
	Stocks []models.Stock `json:"stocks"`
}

func (h *Handler) Handle(ctx context.Context, request Request) (*Response, error) {
	op := "Handler.Handle"
	log.Printf("stocks_handler: %+v", request)

	stocks, err := h.warehouse.Stocks(ctx, request.SKU)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Response{*stocks}, nil
}

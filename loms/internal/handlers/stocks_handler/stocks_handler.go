package stocks_handler

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"route256/loms/internal/repo/warehouse"
	"route256/loms/pkg/models"
)

type Handler struct {
	warehouseRepo warehouse.Repo
}

func New(warehouseRepo warehouse.Repo) *Handler {
	return &Handler{
		warehouseRepo: warehouseRepo,
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
	log.Printf("stocks_handler: %+v", request)

	stocks, err := h.warehouseRepo.GetStocks(request.SKU)
	if err != nil {
		return nil, errors.WithMessage(err, "checking stocks_handler")
	}
	return &Response{*stocks}, nil
}

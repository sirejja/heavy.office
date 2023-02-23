package create_order_handler

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
	User  int64         `json:"user"`
	Items []models.Item `json:"items"`
}

func (r Request) Validate() error {
	if r.User == 0 {
		return models.ErrEmptyUser
	}
	for _, item := range r.Items {
		if item.SKU == 0 {
			return models.ErrEmptySKU
		}
		if item.Count == 0 {
			return models.ErrEmptyCount
		}
	}
	return nil
}

type Response struct {
	OrderId *uint64 `json:"orderID"`
}

func (h *Handler) Handle(ctx context.Context, request Request) (*Response, error) {
	log.Printf("create_order_handler: %+v", request)

	orderId, err := h.warehouseRepo.BookProducts(request.User, request.Items)
	if err != nil {
		return nil, errors.WithMessage(err, "checking stocks_handler")
	}
	return &Response{OrderId: orderId}, nil
}

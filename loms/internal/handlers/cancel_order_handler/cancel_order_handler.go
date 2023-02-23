package cancel_order_handler

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
	OrderID int64 `json:"orderID"`
}

func (r Request) Validate() error {
	if r.OrderID == 0 {
		return models.ErrEmptyOrder
	}
	return nil
}

type Response struct {
}

func (h *Handler) Handle(ctx context.Context, request Request) (*Response, error) {
	log.Printf("payed_order_handler: %+v", request)

	err := h.warehouseRepo.CancellOrder(request.OrderID)
	if err != nil {
		return nil, errors.WithMessage(err, "CancellOrder")
	}
	return &Response{}, nil
}

package cancel_order_handler

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
)

type IService interface {
	CancelOrder(ctx context.Context, orderID int64) error
}

type Handler struct {
	orders IService
}

func New(orders IService) *Handler {
	return &Handler{
		orders: orders,
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
	op := "Handler.Handle"
	log.Printf("payed_order_handler: %+v", request)

	err := h.orders.CancelOrder(ctx, request.OrderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Response{}, nil
}

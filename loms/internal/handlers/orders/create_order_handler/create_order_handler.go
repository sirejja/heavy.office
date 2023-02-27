package create_order_handler

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
)

type IService interface {
	CreateOrder(ctx context.Context, user int64, items []models.Item) (*uint64, error)
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
	op := "Handler.Handle"
	log.Printf("create_order_handler: %+v", request)

	orderId, err := h.orders.CreateOrder(ctx, request.User, request.Items)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Response{OrderId: orderId}, nil
}

package purchase_handler

import (
	"context"
	"fmt"
	"log"
	"route256/checkout/internal/models"
)

type IService interface {
	PurchaseCart(ctx context.Context, user int64) (*uint64, error)
}

type Handler struct {
	model IService
}

func New(service IService) *Handler {
	return &Handler{
		model: service,
	}
}

type Request struct {
	User int64 `json:"user"`
}

func (r Request) Validate() error {
	if r.User == 0 {
		return models.ErrEmptyUser
	}
	return nil
}

type Response struct {
	OrderId *uint64 `json:"orderID"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (*Response, error) {
	op := "Handler.Handle"
	log.Printf("purchase_handler: %+v", req)

	orderID, err := h.model.PurchaseCart(ctx, req.User)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Response{orderID}, nil
}

package delete_from_cart_handler

import (
	"context"
	"fmt"
	"log"
	"route256/checkout/internal/models"
)

type IService interface {
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
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
	User  int64  `json:"user"`
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (r Request) Validate() error {
	if r.User == 0 {
		return models.ErrEmptyUser
	}
	if r.SKU == 0 {
		return models.ErrEmptySKU
	}
	if r.Count == 0 {
		return models.ErrEmptyCount
	}
	return nil
}

type Response struct {
}

func (h *Handler) Handle(ctx context.Context, req Request) (*Response, error) {
	op := "Handler.Handle"
	log.Printf("delete_from_cart_handler: %+v", req)
	var response Response
	if err := h.model.DeleteFromCart(ctx, req.User, req.SKU, req.Count); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &response, nil
}

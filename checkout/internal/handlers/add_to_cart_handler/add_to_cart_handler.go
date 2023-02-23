package add_to_cart_handler

import (
	"context"
	"log"
	"route256/checkout/internal/usecase"
	"route256/loms/pkg/models"
)

type Handler struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type Request struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (r Request) Validate() error {
	if r.User == 0 {
		return models.ErrEmptyUser
	}
	if r.Sku == 0 {
		return models.ErrEmptySKU
	}
	if r.Count == 0 {
		return models.ErrEmptyCount
	}
	return nil
}

type Response struct{}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("addToCart: %+v", req)

	var response Response

	err := h.usecase.Cart.AddToCart(ctx, req.User, req.Sku, req.Count)
	if err != nil {
		return response, err
	}

	return response, nil
}

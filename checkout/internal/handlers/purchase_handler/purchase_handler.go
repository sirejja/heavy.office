package purchase_handler

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
	return nil
}

type Response struct{}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("addToCart: %+v", req)

	var response Response

	err := h.usecase.Cart.PurchaseCart(ctx, req.User)
	if err != nil {
		return response, err
	}

	return response, nil
}

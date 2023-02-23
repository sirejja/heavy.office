package list_cart_handler

import (
	"context"
	"log"
	"route256/checkout/internal/usecase"
	"route256/checkout/internal/usecase/cart_processor"
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
	User int64 `json:"user"`
}

func (r Request) Validate() error {
	if r.User == 0 {
		return models.ErrEmptyUser
	}
	return nil
}

type Response struct {
	Items      []cart_processor.Item `json:"items"`
	TotalPrice uint32                `json:"totalPrice"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("addToCart: %+v", req)

	var response Response

	items, totalPrice, err := h.usecase.Cart.ListCart(ctx, req.User)
	if err != nil {
		return response, err
	}

	return Response{Items: items, TotalPrice: totalPrice}, nil
}

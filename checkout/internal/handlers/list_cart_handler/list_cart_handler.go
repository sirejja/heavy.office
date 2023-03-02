package list_cart_handler

import (
	"context"
	"fmt"
	"log"
	"route256/checkout/internal/models"
)

type IService interface {
	ListCart(ctx context.Context, user int64) (*[]models.Product, *uint32, error)
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

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Response struct {
	Items      []Item `json:"items"`
	TotalPrice uint32 `json:"totalPrice"`
}

func (h *Handler) buildResponse(ctx context.Context, products []models.Product, totalPrice uint32) *Response {
	var response Response
	for _, product := range products {
		response.Items = append(response.Items, Item{
			SKU:   product.SKU,
			Count: product.Count,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	response.TotalPrice = totalPrice
	return &response
}

func (h *Handler) Handle(ctx context.Context, req Request) (*Response, error) {
	op := "Handler.Handle"
	log.Printf("list_cart_handler: %+v", req)

	products, totalPrice, err := h.model.ListCart(ctx, req.User)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return h.buildResponse(ctx, *products, *totalPrice), nil
}

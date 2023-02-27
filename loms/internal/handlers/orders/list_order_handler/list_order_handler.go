package list_order_handler

import (
	"context"
	"fmt"
	"log"
	"route256/loms/internal/models"
	"route256/loms/internal/repo/order_repo"
)

type IService interface {
	ListOrder(ctx context.Context, orderID int64) (*order_repo.Order, error)
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

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	Status models.OrderStatus `json:"status"`
	User   uint64             `json:"user"`
	Items  []Item             `json:"items"`
}

func (h *Handler) Handle(ctx context.Context, request Request) (*Response, error) {
	op := "Handler.Handle"
	log.Printf("create_order_handler: %+v", request)

	order, err := h.orders.ListOrder(ctx, request.OrderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	items := make([]Item, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, Item{SKU: item.SKU, Count: item.Count})
	}
	return &Response{User: order.User, Status: order.Status, Items: items}, nil
}

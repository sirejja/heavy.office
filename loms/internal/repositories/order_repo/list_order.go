package order_repo

import "route256/loms/internal/models"

func (o *OrderRepo) ListOrder(orderID int64) (*models.Order, error) {
	order := models.Order{User: 111, Status: "new", Items: []models.Item{{SKU: 111, Count: 10}}}
	return &order, nil
}

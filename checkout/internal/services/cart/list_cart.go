package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (c *Cart) ListCart(ctx context.Context, user int64) (*[]models.Product, *uint32, error) {
	op := "Cart.ListCart"

	// TODO get cart
	var products []models.Product
	var totalPrice uint32
	// TODO убрать мок
	productCount := uint32(5)
	// TODO убрать мок
	for _, productSku := range []uint32{773297411} {
		product, err := c.productsClient.GetProduct(ctx, productSku)
		if err != nil {
			return nil, nil, fmt.Errorf("%s: %w", op, err)
		}
		products = append(products, models.Product{
			SKU:   productSku,
			Count: uint16(productCount),
			Name:  product.Name,
			Price: product.Price,
		})
		totalPrice += product.Price * productCount
	}

	return &products, &totalPrice, nil
}

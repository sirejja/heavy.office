package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (c *Cart) ListCart(ctx context.Context, user int64) ([]models.CartProduct, uint32, error) {
	op := "Cart.ListCart"

	// TODO get cart
	var products []models.CartProduct
	var totalPrice uint32
	// TODO убрать мок
	productCount := uint32(5)
	// TODO убрать мок
	for _, productSku := range []uint32{1625903, 2618151, 4487693, 773297411} {
		product, err := c.productsClient.GetProduct(ctx, productSku)
		if err != nil {
			return nil, uint32(0), fmt.Errorf("%s: %w", op, err)
		}
		products = append(products, models.CartProduct{
			SKU:   productSku,
			Count: productCount,
			Name:  product.Name,
			Price: product.Price,
		})
		totalPrice += product.Price * productCount
	}

	return products, totalPrice, nil
}

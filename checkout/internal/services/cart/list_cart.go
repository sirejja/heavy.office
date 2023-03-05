package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (c *Cart) ListCart(ctx context.Context, user int64) ([]models.CartProduct, uint32, error) {
	op := "Cart.ListCart"

	var totalPrice uint32

	products := []uint32{1625903, 2618151, 4487693, 773297411}
	cartProducts := make([]models.CartProduct, 0, len(products))
	productCount := uint32(5)

	for _, productSku := range products {

		product, err := c.productsClient.GetProduct(ctx, productSku)
		if err != nil {
			return nil, uint32(0), fmt.Errorf("%s: %w", op, err)
		}

		cartProducts = append(cartProducts, models.CartProduct{
			SKU:   productSku,
			Count: productCount,
			Name:  product.Name,
			Price: product.Price,
		})
		totalPrice += product.Price * productCount
	}

	return cartProducts, totalPrice, nil
}

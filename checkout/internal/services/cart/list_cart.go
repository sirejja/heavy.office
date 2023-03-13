package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/carts_products_repo"
)

func (c *Cart) ListCart(ctx context.Context, user int64) ([]models.CartProduct, uint32, error) {
	op := "Cart.ListCart"

	var totalPrice uint32

	cartProducts, err := c.cartsProductsRepo.GetCartsProducts(ctx, &carts_products_repo.GetCartProductsFilter{UserID: user, IsDeleted: false})
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	resultCartProducts := make([]models.CartProduct, 0)

	for _, cartProduct := range cartProducts {
		product, err := c.productsClient.GetProduct(ctx, (*cartProduct).SKU)
		if err != nil {
			return nil, 0, fmt.Errorf("%s: %w", op, err)
		}

		resultCartProducts = append(resultCartProducts, models.CartProduct{
			SKU:   (*cartProduct).SKU,
			Count: (*cartProduct).Count,
			Name:  product.Name,
			Price: product.Price,
		})
		totalPrice += product.Price * (*cartProduct).Count
	}

	return resultCartProducts, totalPrice, nil
}

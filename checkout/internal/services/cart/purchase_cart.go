package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/carts_products_repo"
	"route256/checkout/internal/repositories/carts_repo"
)

func (c *Cart) PurchaseCart(ctx context.Context, user int64) (int64, error) {
	op := "Cart.PurchaseCart"

	cartProducts, err := c.cartsProductsRepo.GetCartsProducts(ctx, &carts_products_repo.GetCartProductsFilter{UserID: user, IsDeleted: false})
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}

	purchaseCart := make([]models.Item, 0, len(cartProducts))
	for _, item := range cartProducts {
		purchaseCart = append(purchaseCart, models.Item{SKU: item.SKU, Count: item.Count})
	}

	orderID, err := c.lomsClient.CreateOrder(ctx, user, purchaseCart)
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}

	_, err = c.cartsRepo.UpdateCart(ctx,
		&carts_repo.UpdatCartValues{IsPurchased: true},
		&carts_repo.UpdatCartFilter{UserID: user, IsDeleted: false, IsPurchased: false})
	if err != nil {
		return int64(0), fmt.Errorf("%s: %w", op, err)
	}

	return orderID, nil
}

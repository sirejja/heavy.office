package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/carts_products_repo"
	"route256/checkout/internal/repositories/carts_repo"
)

func (c *Cart) AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	op := "Cart.AddToCart"

	_, err := c.checkStocks(ctx, sku, count)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	cart, err := c.cartsRepo.GetCarts(ctx, &carts_repo.GetCartFilter{UserId: user, IsDeleted: false})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if len(cart) > 1 {
		return fmt.Errorf("%s: %w", op, models.ErrUserHasMoreThanOneCart)
	}

	var cartID uint64
	if len(cart) == 0 {
		cartID, err = c.cartsRepo.CreateCart(ctx, &carts_repo.CreateCartIns{UserID: user})
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	} else {
		cartID = (*cart[0]).Id
		cartsProducts, err := c.cartsProductsRepo.GetCartsProducts(ctx,
			&carts_products_repo.GetCartProductsFilter{SKU: sku, IsDeleted: false})
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		if len(cartsProducts) == 1 {
			increasedCount := (*cartsProducts[0]).Count + count
			_, err = c.cartsProductsRepo.UpdateCartProduct(ctx,
				&carts_products_repo.UpdateProductCartValues{SKU: sku, Count: increasedCount},
				&carts_products_repo.UpdateProductCartFilter{CartID: cartID, SKU: sku, IsDeleted: false})
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			return nil
		}
	}

	id, err := c.cartsProductsRepo.AddProductToCart(ctx,
		&carts_products_repo.AddProductToCartInsert{CartID: cartID, SKU: sku, Count: count})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if id == 0 {
		return fmt.Errorf("%s: %w", op, models.ErrInsertFailed)
	}

	return nil
}

func (c *Cart) checkStocks(ctx context.Context, sku uint32, count uint32) ([]models.Stock, error) {
	op := "Cart.checkStocks"

	stocks, err := c.lomsClient.Stocks(ctx, sku)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			break
		}
	}
	if counter > 0 {
		return nil, fmt.Errorf("%s: %w", op, models.ErrInsufficientStocks)
	}

	return stocks, nil
}

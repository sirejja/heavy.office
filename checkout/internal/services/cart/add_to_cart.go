package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
)

func (c *Cart) AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	op := "Cart.AddToCart"

	_, err := c.checkStocks(ctx, sku, count)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = c.txManager.RunRepeatableRead(ctx, func(ctxTX context.Context) error {

		cartID, err := c.cartsRepo.GetCartID(ctxTX, user)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		if cartID == 0 {
			cartID, err = c.cartsRepo.CreateCart(ctxTX, user)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
		} else {
			productCount, err := c.cartsProductsRepo.GetCartProductCount(ctxTX, sku, cartID)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			if productCount != 0 {
				increasedCount := productCount + count
				_, err = c.cartsProductsRepo.UpdateCartProduct(ctxTX, uint64(sku), increasedCount, uint32(cartID))
				if err != nil {
					return fmt.Errorf("%s: %w", op, err)
				}
				return nil
			}
		}

		id, err := c.cartsProductsRepo.AddProductToCart(ctxTX, cartID, sku, count)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		if id == 0 {
			return fmt.Errorf("%s: %w", op, models.ErrInsertFailed)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
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

package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/libs/worker_pool"
	"sync"
)

func (c *Cart) ListCart(ctx context.Context, user int64) ([]models.CartProduct, uint32, error) {
	op := "Cart.ListCart"

	resultCartProducts, totalPrice, err := c.processGetProductBatch(ctx, user)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	return resultCartProducts, totalPrice, nil
}

func (c *Cart) processGetProductBatch(ctx context.Context, user int64) ([]models.CartProduct, uint32, error) {
	op := "Cart.processGetProductBatch"

	cartProducts, err := c.cartsProductsRepo.GetCartsProducts(ctx, user)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	var totalPrice uint32
	var mutex sync.Mutex

	callbacks := make([]func(struct{}) map[string]any, 0, len(cartProducts))
	for _, cartProduct := range cartProducts {
		cartProduct := cartProduct
		callbacks = append(callbacks, func(struct{}) map[string]any {
			err = c.productsLimiter.Wait(ctx)
			if err != nil {
				return map[string]any{"error": fmt.Errorf("%s: %w", op, err)}
			}

			product, err := c.productsClient.GetProduct(ctx, cartProduct.SKU)
			if err != nil {
				return map[string]any{"error": fmt.Errorf("%s: %w", op, err)}
			}
			mutex.Lock()
			totalPrice += product.Price * cartProduct.Count
			mutex.Unlock()
			return map[string]any{
				"data": &models.CartProduct{
					SKU:   cartProduct.SKU,
					Count: cartProduct.Count,
					Name:  product.Name,
					Price: product.Price,
				},
				"error": nil,
			}
		})
	}

	amountWorkers := 5
	batchingPool, workerCh := worker_pool.NewPool[struct{}, map[string]any](ctx, amountWorkers)

	var wg sync.WaitGroup
	tasks := make([]worker_pool.Task[struct{}, map[string]any], 0, len(cartProducts))
	for _, callback := range callbacks {
		wg.Add(1)
		tasks = append(tasks, worker_pool.Task[struct{}, map[string]any]{
			Callback: callback,
			InArgs:   struct{}{},
		})
	}

	resultCartProducts := make([]models.CartProduct, 0, len(cartProducts))
	batchingPool.Submit(ctx, tasks)

	go func() {
		for res := range workerCh {
			wg.Done()
			if res["error"] != nil {
				err = fmt.Errorf("%s: %w", op, err)
			}
			resultCartProducts = append(resultCartProducts, *(res["data"].(*models.CartProduct)))
		}
	}()
	wg.Wait()
	batchingPool.Close()
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	return resultCartProducts, totalPrice, nil
}

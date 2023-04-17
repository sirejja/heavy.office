package cart

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/libs/cache/inmemory"
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

	callbacks := make([]func(struct{}) *worker_pool.OutSink[*models.CartProduct], 0, len(cartProducts))
	for _, cartProduct := range cartProducts {
		cartProduct := cartProduct
		callbacks = append(callbacks, func(struct{}) *worker_pool.OutSink[*models.CartProduct] {

			inmemory.CacheRequestsTotal.Inc()
			product, err := c.productsClient.GetProductCached(cartProduct.SKU)
			if err != nil {
				err = nil
				err = c.productsLimiter.Wait(ctx)
				if err != nil {
					return &worker_pool.OutSink[*models.CartProduct]{Res: nil, Err: fmt.Errorf("%s: %w", op, err)}
				}
				product, err = c.productsClient.GetProduct(ctx, cartProduct.SKU)
			}
			if err != nil {
				return &worker_pool.OutSink[*models.CartProduct]{Res: nil, Err: fmt.Errorf("%s: %w", op, err)}
			}
			mutex.Lock()
			totalPrice += product.Price * cartProduct.Count
			mutex.Unlock()
			return &worker_pool.OutSink[*models.CartProduct]{
				Res: &models.CartProduct{
					SKU:   cartProduct.SKU,
					Count: cartProduct.Count,
					Name:  product.Name,
					Price: product.Price,
				},
				Err: nil,
			}
		})
	}

	amountWorkers := 5
	batchingPool, workerCh := worker_pool.NewPool[struct{}, *models.CartProduct](ctx, amountWorkers)

	var wg sync.WaitGroup
	tasks := make([]worker_pool.Task[struct{}, *models.CartProduct], 0, len(cartProducts))
	for _, callback := range callbacks {
		wg.Add(1)
		tasks = append(tasks, worker_pool.Task[struct{}, *models.CartProduct]{
			Callback: callback,
			InArgs:   struct{}{},
		})
	}

	resultCartProducts := make([]models.CartProduct, 0, len(cartProducts))
	batchingPool.Submit(ctx, tasks)

	// TODO check goroutine leak
	go func() {
		for res := range workerCh {
			wg.Done()
			if res.Err != nil {
				err = fmt.Errorf("%s: %w", op, res.Err)
				continue
			}
			resultCartProducts = append(resultCartProducts, *res.Res)
		}
	}()
	wg.Wait()
	batchingPool.Close()
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	return resultCartProducts, totalPrice, nil
}

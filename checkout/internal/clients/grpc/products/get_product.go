package products

import (
	"context"
	"fmt"
	"route256/checkout/internal/models"
	"route256/libs/cache/inmemory"
	"route256/libs/logger"
	product_service "route256/product_service/pkg/v1/api"
	"strconv"
	"time"
)

func (c *Client) GetProduct(ctx context.Context, Sku uint32) (*models.ProductAttrs, error) {
	op := "Client.GetProduct"

	response, err := c.client.GetProduct(ctx, &product_service.GetProductRequest{Sku: Sku, Token: c.token})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	product := models.ProductAttrs{Name: response.GetName(), Price: response.GetPrice()}

	go func() {
		if ok := c.cache.Set(strconv.Itoa(int(Sku)), &product); !ok {
			inmemory.CacheErrorsTotal.Inc()
			return
		}
	}()

	return &product, nil
}

func (c *Client) GetProductCached(Sku uint32) (*models.ProductAttrs, error) {
	op := "Client.GetProductCached"
	timeStart := time.Now()

	res, err := c.cache.Get(strconv.Itoa(int(Sku)))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("retrieved from cache")

	inmemory.CacheHitsTotal.Inc()
	elapsed := time.Since(timeStart)
	inmemory.HistogramResponseTimeCache.Observe(elapsed.Seconds())
	return res.(*models.ProductAttrs), nil
}

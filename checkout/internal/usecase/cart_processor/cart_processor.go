package cart_processor

import (
	"context"
	"github.com/pkg/errors"
	"route256/libs/clients/loms_client"
	"route256/libs/clients/products_client"
	"route256/loms/pkg/models"
)

type CartUsecase interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error
	ListCart(ctx context.Context, user int64) ([]Item, uint32, error)
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error
	PurchaseCart(ctx context.Context, user int64) error
}

type CartProcessor struct {
	lomsClient     *loms_client.Client
	productsClient *products_client.Client
}

func New(lomsClient *loms_client.Client, productsClient *products_client.Client) CartUsecase {
	return &CartProcessor{
		lomsClient:     lomsClient,
		productsClient: productsClient,
	}
}

func (c *CartProcessor) AddToCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	stocks, err := c.lomsClient.Stocks(ctx, sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks_handler")
	}
	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			// TODO set cart to user
			return nil
		}
	}

	return models.ErrInsufficientStocks
}

func (c *CartProcessor) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	return nil
}

func (c *CartProcessor) PurchaseCart(ctx context.Context, user int64) error {
	if err := c.lomsClient.CreateOrder(ctx, user); err != nil {
		return errors.WithMessage(err, "PurchaseCart")
	}
	return nil
}

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *CartProcessor) ListCart(ctx context.Context, user int64) ([]Item, uint32, error) {
	// TODO get cart
	var products []Item
	var totalPrice uint32
	for _, productSku := range []uint32{773297411} {
		productCount := productSku
		product, err := c.productsClient.GetProduct(ctx, productSku)
		if err != nil {
			return products, totalPrice, errors.WithMessage(err, "ListCart")
		}
		products = append(products, Item{productSku, uint16(productCount), product.Name, product.Price})
		totalPrice += product.Price
	}

	return products, totalPrice, nil
}

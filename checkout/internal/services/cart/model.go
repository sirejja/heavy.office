package cart

import (
	"context"
	"route256/checkout/internal/clients/grpc/loms"
	"route256/checkout/internal/clients/grpc/products"
	"route256/checkout/internal/models"
	"route256/checkout/internal/repositories/carts_products_repo"
	"route256/checkout/internal/repositories/carts_repo"
	"route256/libs/transactor"

	"golang.org/x/time/rate"
)

type ICartProcessor interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error
	ListCart(ctx context.Context, user int64) ([]models.CartProduct, uint32, error)
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error
	PurchaseCart(ctx context.Context, user int64) (int64, error)
}

var _ ICartProcessor = (*Cart)(nil)

type Cart struct {
	lomsClient        loms.ILOMSClient
	productsClient    products.IProductServiceClient
	cartsRepo         carts_repo.ICartsRepo
	cartsProductsRepo carts_products_repo.ICartsProductsRepo
	txManager         transactor.ITransactor
	productsLimiter   *rate.Limiter
}

func New(
	lomsClient loms.ILOMSClient,
	productsClient products.IProductServiceClient,
	cartsRepo carts_repo.ICartsRepo,
	cartsProductsRepo carts_products_repo.ICartsProductsRepo,
	txManager transactor.ITransactor,
	productsLimiter *rate.Limiter,
) *Cart {
	return &Cart{
		lomsClient:        lomsClient,
		productsClient:    productsClient,
		cartsRepo:         cartsRepo,
		cartsProductsRepo: cartsProductsRepo,
		txManager:         txManager,
		productsLimiter:   productsLimiter,
	}
}

func NewMockService(deps ...interface{}) *Cart {
	ns := Cart{}

	for _, v := range deps {
		switch s := v.(type) {
		case loms.ILOMSClient:
			ns.lomsClient = s
		case products.IProductServiceClient:
			ns.productsClient = s
		case carts_repo.ICartsRepo:
			ns.cartsRepo = s
		case carts_products_repo.ICartsProductsRepo:
			ns.cartsProductsRepo = s
		case transactor.ITransactor:
			ns.txManager = s
		case *rate.Limiter:
			ns.productsLimiter = s
		}
	}

	return &ns
}

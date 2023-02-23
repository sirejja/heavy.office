package usecase

import (
	"route256/checkout/internal/usecase/cart_processor"
	"route256/libs/clients/loms_client"
	"route256/libs/clients/products_client"
)

type Usecase struct {
	Cart cart_processor.CartUsecase
}

func New(lomsClient *loms_client.Client, productsClient *products_client.Client) Usecase {
	return Usecase{
		Cart: cart_processor.New(lomsClient, productsClient),
	}
}

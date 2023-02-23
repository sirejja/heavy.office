package main

import (
	"log"
	"net/http"
	"route256/checkout/internal/config"
	"route256/checkout/internal/handlers/add_to_cart_handler"
	"route256/checkout/internal/handlers/delete_from_cart_handler"
	"route256/checkout/internal/handlers/list_cart_handler"
	"route256/checkout/internal/handlers/purchase_handler"
	"route256/checkout/internal/usecase"
	"route256/libs/clients/loms_client"
	"route256/libs/clients/products_client"
	"route256/libs/server_wrapper"
)

const port = ":8080"

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	lomsClient := loms_client.New(config.ConfigData.Services.Loms.URL)
	productsClient := products_client.New(config.ConfigData.Services.Products.URL, config.ConfigData.Services.Products.Token)
	Usecase := usecase.New(lomsClient, productsClient)
	// TODO прокинуть репу
	addToCartHandler := add_to_cart_handler.New(Usecase)
	deleteFromCart := delete_from_cart_handler.New(Usecase)
	purchase := purchase_handler.New(Usecase)
	listCart := list_cart_handler.New(Usecase)

	http.Handle("/addToCart", server_wrapper.New(addToCartHandler.Handle))
	http.Handle("/deleteFromCart", server_wrapper.New(deleteFromCart.Handle))
	http.Handle("/listCart", server_wrapper.New(listCart.Handle))
	http.Handle("/purchase", server_wrapper.New(purchase.Handle))

	log.Println("listening http at", port)
	err = http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}

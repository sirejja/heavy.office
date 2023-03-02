package main

import (
	"log"
	"net/http"
	"route256/checkout/internal/clients/loms"
	"route256/checkout/internal/clients/products"
	"route256/checkout/internal/config"
	"route256/checkout/internal/handlers/add_to_cart_handler"
	"route256/checkout/internal/handlers/delete_from_cart_handler"
	"route256/checkout/internal/handlers/list_cart_handler"
	"route256/checkout/internal/handlers/purchase_handler"
	"route256/checkout/internal/services/cart"
	"route256/libs/server_wrapper"
)

const port = ":8080"

func main() {
	cfg := config.New()
	err := cfg.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	lomsClient := loms.New(cfg.Services.Loms.URL)
	productsClient := products.New(cfg.Services.Products.URL, cfg.Services.Products.Token)
	cartProcessor := cart.New(lomsClient, productsClient)

	addToCartHandler := add_to_cart_handler.New(cartProcessor)
	deleteFromCart := delete_from_cart_handler.New(cartProcessor)
	purchase := purchase_handler.New(cartProcessor)
	listCart := list_cart_handler.New(cartProcessor)

	http.Handle("/addToCart", server_wrapper.New(addToCartHandler.Handle))
	http.Handle("/deleteFromCart", server_wrapper.New(deleteFromCart.Handle))
	http.Handle("/listCart", server_wrapper.New(listCart.Handle))
	http.Handle("/purchase", server_wrapper.New(purchase.Handle))

	log.Println("listening http at", port)
	err = http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}

package main

import (
	"log"
	"net/http"
	"route256/libs/server_wrapper"
	"route256/loms/internal/config"
	"route256/loms/internal/handlers/cancel_order_handler"
	"route256/loms/internal/handlers/create_order_handler"
	"route256/loms/internal/handlers/list_order_handler"
	"route256/loms/internal/handlers/payed_order_handler"
	"route256/loms/internal/handlers/stocks_handler"
	"route256/loms/internal/repo/warehouse"
)

const port = ":8081"

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}
	// TODO прокинуть юзкейсы бизнес логики
	warehouseRepo := warehouse.New()

	stocksHandler := stocks_handler.New(warehouseRepo)
	createOrderHandler := create_order_handler.New(warehouseRepo)
	listOrderHandler := list_order_handler.New(warehouseRepo)
	payedOrderHandler := payed_order_handler.New(warehouseRepo)
	cancelOrderHandler := cancel_order_handler.New(warehouseRepo)

	http.Handle("/createOrder", server_wrapper.New(createOrderHandler.Handle))
	http.Handle("/listOrder", server_wrapper.New(listOrderHandler.Handle))
	http.Handle("/orderPayed", server_wrapper.New(payedOrderHandler.Handle))
	http.Handle("/cancelOrder", server_wrapper.New(cancelOrderHandler.Handle))
	http.Handle("/stocks", server_wrapper.New(stocksHandler.Handle))

	log.Println("listening http at", port)
	err = http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}

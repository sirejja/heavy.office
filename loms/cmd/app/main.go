package main

import (
	"log"
	"net/http"
	"route256/libs/server_wrapper"
	"route256/loms/internal/config"
	"route256/loms/internal/handlers/orders/cancel_order_handler"
	"route256/loms/internal/handlers/orders/create_order_handler"
	"route256/loms/internal/handlers/orders/list_order_handler"
	"route256/loms/internal/handlers/orders/payed_order_handler"
	"route256/loms/internal/handlers/warehouse/stocks_handler"
	"route256/loms/internal/repo/order_repo"
	"route256/loms/internal/repo/warehouse_repo"
	"route256/loms/internal/services/orders"
	"route256/loms/internal/services/warehouse"
)

const port = ":8081"

func main() {
	cfg := config.New()
	err := cfg.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	warehouseRepo := warehouse_repo.New()
	ordersRepo := order_repo.New()

	ordersProcessor := orders.New(ordersRepo, warehouseRepo)
	warehouseProcessor := warehouse.New(warehouseRepo)

	// warehouse handlers
	stocksHandler := stocks_handler.New(warehouseProcessor)

	// orders handlers
	createOrderHandler := create_order_handler.New(ordersProcessor)
	listOrderHandler := list_order_handler.New(ordersProcessor)
	payedOrderHandler := payed_order_handler.New(ordersProcessor)
	cancelOrderHandler := cancel_order_handler.New(ordersProcessor)

	http.Handle("/createOrder", server_wrapper.New(createOrderHandler.Handle))
	http.Handle("/listOrder", server_wrapper.New(listOrderHandler.Handle))
	http.Handle("/orderPayed", server_wrapper.New(payedOrderHandler.Handle))
	http.Handle("/cancelOrder", server_wrapper.New(cancelOrderHandler.Handle))
	http.Handle("/stocks", server_wrapper.New(stocksHandler.Handle))

	log.Println("listening http at", port)
	err = http.ListenAndServe(port, nil)
	log.Fatal("cannot listen http", err)
}

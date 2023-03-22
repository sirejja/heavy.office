package cronjob

import (
	"context"
	"log"
	"route256/loms/internal/cronjob/cancel_orders_cron"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/services/orders"

	"github.com/robfig/cron"
)

type CronJob struct {
	cronjob          *cron.Cron
	orders           *orders.Order
	orderRepo        *order_repo.OrderRepo
	specCancelOrders string
}

func New(orders *orders.Order, orderRepo *order_repo.OrderRepo, specCancelOrders string) CronJob {
	cronjob := cron.New()

	return CronJob{cronjob: cronjob, orders: orders, orderRepo: orderRepo, specCancelOrders: specCancelOrders}
}

func (c *CronJob) Start(ctx context.Context) {
	op := "CronJob.Start"

	cancelOrdersCron := cancel_orders_cron.New(ctx, c.orders, c.orderRepo)

	err := c.cronjob.AddJob(c.specCancelOrders, cancelOrdersCron)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}

	c.cronjob.Start()
}

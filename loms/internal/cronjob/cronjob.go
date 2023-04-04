package cronjob

import (
	"log"
	"route256/loms/internal/services/cancel_orders_cron"

	"github.com/robfig/cron"
)

type CronJob struct {
	cronjob         *cron.Cron
	cancelOrdersJob *cancel_orders_cron.CancelOrdersJob
}

func New(cancelOrdersCron *cancel_orders_cron.CancelOrdersJob) CronJob {
	return CronJob{
		cronjob:         cron.New(),
		cancelOrdersJob: cancelOrdersCron,
	}
}

func (c *CronJob) Start() {
	op := "CronJob.Start"

	err := c.cronjob.AddJob(c.cancelOrdersJob.SpecCancelOrders, c.cancelOrdersJob)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}

	c.cronjob.Start()
}

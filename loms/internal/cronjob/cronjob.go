package cronjob

import (
	"fmt"
	"route256/libs/logger"
	"route256/loms/internal/services/cron/cancel_orders"
	"route256/loms/internal/services/cron/outbox"

	"github.com/robfig/cron"
	"go.uber.org/zap"
)

type CronJob struct {
	cronjob         *cron.Cron
	cancelOrdersJob *cancel_orders.CancelOrdersJob
	outboxJob       *outbox.OutboxJob
}

func New(cancelOrdersJob *cancel_orders.CancelOrdersJob, outboxJob *outbox.OutboxJob) CronJob {
	return CronJob{
		cronjob:         cron.New(),
		cancelOrdersJob: cancelOrdersJob,
		outboxJob:       outboxJob,
	}
}

func (c *CronJob) Start() {
	op := "CronJob.Start"

	err := c.cronjob.AddJob(c.cancelOrdersJob.SpecCancelOrders, c.cancelOrdersJob)
	if err != nil {
		logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
	}
	err = c.cronjob.AddJob(c.outboxJob.SpecOutbox, c.outboxJob)
	if err != nil {
		logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
	}

	c.cronjob.Start()
}

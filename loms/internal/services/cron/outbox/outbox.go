package outbox

import (
	"context"
	"fmt"
	"route256/libs/logger"
	"route256/loms/internal/kafka/outbox_producer"
	"route256/loms/internal/repositories/outbox_repo"

	"github.com/robfig/cron"
	"go.uber.org/zap"
)

type OutboxJob struct {
	outboxRepo outbox_repo.IOutboxRepo
	ctx        context.Context
	producer   outbox_producer.IProducerHandler
	SpecOutbox string
}

type IOutboxCron interface {
	cron.Job
}

var _ IOutboxCron = (*OutboxJob)(nil)

func New(
	ctx context.Context,
	producer outbox_producer.IProducerHandler,
	outboxRepo outbox_repo.IOutboxRepo,
	specOutbox string,
) *OutboxJob {
	return &OutboxJob{
		ctx:        ctx,
		producer:   producer,
		outboxRepo: outboxRepo,
		SpecOutbox: specOutbox,
	}
}

func (o *OutboxJob) Run() {
	op := "OutboxJob.Run"

	tasks, err := o.outboxRepo.GetTasks(o.ctx)
	if err != nil {
		logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
		return
	}

	for _, task := range tasks {
		if err = o.producer.ResolveProducerHandler(task.Topic, task.Args); err != nil {
			logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
			return
		}
		// TODO добавить функционал дропа пропущенных из-за ошибок таск
		if _, err = o.outboxRepo.DeleteTask(o.ctx, task.ID); err != nil {
			logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
			return
		}

	}

}

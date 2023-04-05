package outbox

import (
	"context"
	"log"
	"route256/loms/internal/kafka/outbox_producer"
	"route256/loms/internal/repositories/outbox_repo"
	"time"

	"github.com/robfig/cron"
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
	log.Printf("%s started at %s", op, time.Now().Format("2006-01-02 15:04"))

	tasks, err := o.outboxRepo.GetTasks(o.ctx)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return
	}

	for _, task := range tasks {
		if err = o.producer.ResolveProducerHandler(task.Topic, task.Args); err != nil {
			log.Printf("%s: %v", op, err)
			return
		}
		// TODO добавить функционал дропа пропущенных из-за ошибок таск
		if _, err = o.outboxRepo.DeleteTask(o.ctx, task.ID); err != nil {
			log.Printf("%s: %v", op, err)
			return
		}

	}

	log.Printf("%s finished at %s", op, time.Now().Format("2006-01-02 15:04"))
}

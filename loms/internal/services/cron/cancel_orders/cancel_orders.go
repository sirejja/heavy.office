package cancel_orders

import (
	"context"
	"fmt"
	"route256/libs/logger"
	"route256/libs/worker_pool"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/services/orders"
	"sync"
	"time"

	"github.com/robfig/cron"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type CancelOrdersJob struct {
	orders           orders.IOrdersService
	orderRepo        order_repo.IOrderRepo
	ctx              context.Context
	SpecCancelOrders string
}

type ICancelOrdersCron interface {
	cron.Job
}

var _ ICancelOrdersCron = (*CancelOrdersJob)(nil)

func New(ctx context.Context, orders orders.IOrdersService, orderRepo order_repo.IOrderRepo, specCancelOrders string) *CancelOrdersJob {
	return &CancelOrdersJob{
		ctx:              ctx,
		orders:           orders,
		orderRepo:        orderRepo,
		SpecCancelOrders: specCancelOrders,
	}
}

func (o *CancelOrdersJob) Run() {
	op := "CancelOrdersJob.Run"
	logger.Info("cron started", zap.String("op", op), zap.Time("timestamp", time.Now()))

	orderIDs, err := o.orderRepo.GetOrdersForCancel(o.ctx)
	if err != nil {
		logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
		return
	}
	limiter := rate.NewLimiter(rate.Every(time.Second/100), 40)

	callbacks := make([]func(struct{}) *worker_pool.OutSink[struct{}], 0, len(orderIDs))
	for _, orderID := range orderIDs {
		orderID := orderID
		callbacks = append(callbacks, func(struct{}) *worker_pool.OutSink[struct{}] {
			err = limiter.Wait(o.ctx)
			if err != nil {
				logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
				return &worker_pool.OutSink[struct{}]{Res: struct{}{}, Err: fmt.Errorf("%s: %w", op, err)}
			}

			err = o.orders.CancelOrder(o.ctx, orderID)
			if err != nil {
				logger.Error(op, zap.Error(fmt.Errorf("%s: %v", op, err)))
				return &worker_pool.OutSink[struct{}]{Res: struct{}{}, Err: fmt.Errorf("%s: %w", op, err)}
			}
			return &worker_pool.OutSink[struct{}]{Res: struct{}{}, Err: nil}
		})
	}

	amountWorkers := 20
	batchingPool, workerCh := worker_pool.NewPool[struct{}, struct{}](o.ctx, amountWorkers)

	var wg sync.WaitGroup
	tasks := make([]worker_pool.Task[struct{}, struct{}], 0, len(orderIDs))
	for _, callback := range callbacks {
		wg.Add(1)
		tasks = append(tasks, worker_pool.Task[struct{}, struct{}]{
			Callback: callback,
			InArgs:   struct{}{},
		})
	}

	batchingPool.Submit(o.ctx, tasks)

	go func() {
		for range workerCh {
			wg.Done()
		}
	}()
	wg.Wait()
	batchingPool.Close()
	logger.Info("cron finished", zap.String("op", op), zap.Time("timestamp", time.Now()))
}

package cancel_orders_cron

import (
	"context"
	"fmt"
	"log"
	"route256/libs/worker_pool"
	"route256/loms/internal/repositories/order_repo"
	"route256/loms/internal/services/orders"
	"sync"
	"time"

	"github.com/robfig/cron"
	"golang.org/x/time/rate"
)

type CancelOrdersCronService struct {
	orders    orders.IOrdersService
	orderRepo order_repo.IOrderRepo
	ctx       context.Context
}

type ICancelOrdersCron interface {
	cron.Job
}

var _ ICancelOrdersCron = (*CancelOrdersCronService)(nil)

func New(ctx context.Context, orders orders.IOrdersService, orderRepo order_repo.IOrderRepo) *CancelOrdersCronService {
	return &CancelOrdersCronService{
		ctx:       ctx,
		orders:    orders,
		orderRepo: orderRepo,
	}
}

func (r *CancelOrdersCronService) Run() {
	op := "CancelOrdersCronService.Run"
	log.Printf("%s started at %s", op, time.Now().Format("2006-01-02 15:04"))

	orderIDs, err := r.orderRepo.GetOrdersForCancel(r.ctx)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return
	}
	limiter := rate.NewLimiter(rate.Every(time.Second/100), 40)

	callbacks := make([]func(struct{}) *worker_pool.OutSink[struct{}], 0, len(orderIDs))
	for _, orderID := range orderIDs {
		orderID := orderID
		callbacks = append(callbacks, func(struct{}) *worker_pool.OutSink[struct{}] {
			err = limiter.Wait(r.ctx)
			if err != nil {
				log.Printf("%s: %v", op, err)
				return &worker_pool.OutSink[struct{}]{Res: struct{}{}, Err: fmt.Errorf("%s: %w", op, err)}
			}

			err = r.orders.CancelOrder(r.ctx, orderID)
			if err != nil {
				log.Printf("%s: %v", op, err)
				return &worker_pool.OutSink[struct{}]{Res: struct{}{}, Err: fmt.Errorf("%s: %w", op, err)}
			}
			return &worker_pool.OutSink[struct{}]{Res: struct{}{}, Err: nil}
		})
	}

	amountWorkers := 20
	batchingPool, workerCh := worker_pool.NewPool[struct{}, struct{}](r.ctx, amountWorkers)

	var wg sync.WaitGroup
	tasks := make([]worker_pool.Task[struct{}, struct{}], 0, len(orderIDs))
	for _, callback := range callbacks {
		wg.Add(1)
		tasks = append(tasks, worker_pool.Task[struct{}, struct{}]{
			Callback: callback,
			InArgs:   struct{}{},
		})
	}

	batchingPool.Submit(r.ctx, tasks)

	go func() {
		for range workerCh {
			wg.Done()
		}
	}()
	wg.Wait()
	batchingPool.Close()
	log.Printf("%s finished at %s", op, time.Now().Format("2006-01-02 15:04"))
}

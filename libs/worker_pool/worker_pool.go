package worker_pool

import (
	"context"
	"sync"
)

type Task[In, Out any] struct {
	Callback func(In) Out
	InArgs   In
}

type IPool[In, Out any] interface {
	Submit(ctx context.Context, tasks []Task[In, Out])
	//SplitTasksByChunks(in []Task[In, Out], chunkSize int) [][]Task[In, Out]
	Close()
}

var _ IPool[any, any] = &poolInstance[any, any]{}

type poolInstance[In, Out any] struct {
	amountWorkers int

	wg sync.WaitGroup

	taskSource chan Task[In, Out]
	outSink    chan Out
}

func NewPool[In, Out any](ctx context.Context, amountWorkers int) (IPool[In, Out], <-chan Out) {
	pool := &poolInstance[In, Out]{
		amountWorkers: amountWorkers,
	}

	pool.bootstrap(ctx)

	return pool, pool.outSink
}

// Close implements IPool
func (p *poolInstance[In, Out]) Close() {
	// Больше задач не будет
	close(p.taskSource)

	// Дожидаемся, пока все воркеры закончат работы
	p.wg.Wait()

	// Закрываем канал на выход, чтобы потребители могли выйти из := range
	close(p.outSink)
}

// Submit implements IPool
func (p *poolInstance[In, Out]) Submit(ctx context.Context, tasks []Task[In, Out]) {

	go func() {
		for _, task := range tasks {
			select {
			case <-ctx.Done():
				return

			case p.taskSource <- task:
			}
		}
	}()
}

func (p *poolInstance[In, Out]) bootstrap(ctx context.Context) {
	p.taskSource = make(chan Task[In, Out], p.amountWorkers)
	p.outSink = make(chan Out, p.amountWorkers)

	for i := 0; i < p.amountWorkers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			worker(ctx, p.taskSource, p.outSink)
		}()
	}
}

func worker[In, Out any](
	ctx context.Context,
	taskSource <-chan Task[In, Out],
	resultSink chan<- Out,
) {
	for task := range taskSource {
		select {
		case <-ctx.Done():
			return
		case resultSink <- task.Callback(task.InArgs):
		}
	}
	return
}

func splitTasksByChunks[In, Out any](tasks []Task[In, Out], chunkSize int) [][]Task[In, Out] {
	if len(tasks) == 0 {
		return nil
	}

	if len(tasks) <= chunkSize {
		return [][]Task[In, Out]{tasks}
	}

	var result [][]Task[In, Out]
	chunk := make([]Task[In, Out], 0, len(tasks))
	for _, task := range tasks {
		chunk = append(chunk, task)
		if len(chunk) == chunkSize {
			result = append(result, chunk)
			chunk = make([]Task[In, Out], 0, len(tasks))
		}
	}

	return result
}

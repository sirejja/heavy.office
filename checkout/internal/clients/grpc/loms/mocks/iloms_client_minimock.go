package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/clients/loms.ILOMSClient -o ./mocks/iloms_client_minimock.go -n ILOMSClientMock

import (
	"context"
	"route256/checkout/internal/models"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ILOMSClientMock implements loms.ILOMSClient
type ILOMSClientMock struct {
	t minimock.Tester

	funcCreateOrder          func(ctx context.Context, user int64, items []models.Item) (i1 int64, err error)
	inspectFuncCreateOrder   func(ctx context.Context, user int64, items []models.Item)
	afterCreateOrderCounter  uint64
	beforeCreateOrderCounter uint64
	CreateOrderMock          mILOMSClientMockCreateOrder

	funcStocks          func(ctx context.Context, sku uint32) (sa1 []models.Stock, err error)
	inspectFuncStocks   func(ctx context.Context, sku uint32)
	afterStocksCounter  uint64
	beforeStocksCounter uint64
	StocksMock          mILOMSClientMockStocks
}

// NewILOMSClientMock returns a mock for loms.ILOMSClient
func NewILOMSClientMock(t minimock.Tester) *ILOMSClientMock {
	m := &ILOMSClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateOrderMock = mILOMSClientMockCreateOrder{mock: m}
	m.CreateOrderMock.callArgs = []*ILOMSClientMockCreateOrderParams{}

	m.StocksMock = mILOMSClientMockStocks{mock: m}
	m.StocksMock.callArgs = []*ILOMSClientMockStocksParams{}

	return m
}

type mILOMSClientMockCreateOrder struct {
	mock               *ILOMSClientMock
	defaultExpectation *ILOMSClientMockCreateOrderExpectation
	expectations       []*ILOMSClientMockCreateOrderExpectation

	callArgs []*ILOMSClientMockCreateOrderParams
	mutex    sync.RWMutex
}

// ILOMSClientMockCreateOrderExpectation specifies expectation struct of the ILOMSClient.CreateOrder
type ILOMSClientMockCreateOrderExpectation struct {
	mock    *ILOMSClientMock
	params  *ILOMSClientMockCreateOrderParams
	results *ILOMSClientMockCreateOrderResults
	Counter uint64
}

// ILOMSClientMockCreateOrderParams contains parameters of the ILOMSClient.CreateOrder
type ILOMSClientMockCreateOrderParams struct {
	ctx   context.Context
	user  int64
	items []models.Item
}

// ILOMSClientMockCreateOrderResults contains results of the ILOMSClient.CreateOrder
type ILOMSClientMockCreateOrderResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for ILOMSClient.CreateOrder
func (mmCreateOrder *mILOMSClientMockCreateOrder) Expect(ctx context.Context, user int64, items []models.Item) *mILOMSClientMockCreateOrder {
	if mmCreateOrder.mock.funcCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("ILOMSClientMock.CreateOrder mock is already set by Set")
	}

	if mmCreateOrder.defaultExpectation == nil {
		mmCreateOrder.defaultExpectation = &ILOMSClientMockCreateOrderExpectation{}
	}

	mmCreateOrder.defaultExpectation.params = &ILOMSClientMockCreateOrderParams{ctx, user, items}
	for _, e := range mmCreateOrder.expectations {
		if minimock.Equal(e.params, mmCreateOrder.defaultExpectation.params) {
			mmCreateOrder.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateOrder.defaultExpectation.params)
		}
	}

	return mmCreateOrder
}

// Inspect accepts an inspector function that has same arguments as the ILOMSClient.CreateOrder
func (mmCreateOrder *mILOMSClientMockCreateOrder) Inspect(f func(ctx context.Context, user int64, items []models.Item)) *mILOMSClientMockCreateOrder {
	if mmCreateOrder.mock.inspectFuncCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("Inspect function is already set for ILOMSClientMock.CreateOrder")
	}

	mmCreateOrder.mock.inspectFuncCreateOrder = f

	return mmCreateOrder
}

// Return sets up results that will be returned by ILOMSClient.CreateOrder
func (mmCreateOrder *mILOMSClientMockCreateOrder) Return(i1 int64, err error) *ILOMSClientMock {
	if mmCreateOrder.mock.funcCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("ILOMSClientMock.CreateOrder mock is already set by Set")
	}

	if mmCreateOrder.defaultExpectation == nil {
		mmCreateOrder.defaultExpectation = &ILOMSClientMockCreateOrderExpectation{mock: mmCreateOrder.mock}
	}
	mmCreateOrder.defaultExpectation.results = &ILOMSClientMockCreateOrderResults{i1, err}
	return mmCreateOrder.mock
}

// Set uses given function f to mock the ILOMSClient.CreateOrder method
func (mmCreateOrder *mILOMSClientMockCreateOrder) Set(f func(ctx context.Context, user int64, items []models.Item) (i1 int64, err error)) *ILOMSClientMock {
	if mmCreateOrder.defaultExpectation != nil {
		mmCreateOrder.mock.t.Fatalf("Default expectation is already set for the ILOMSClient.CreateOrder method")
	}

	if len(mmCreateOrder.expectations) > 0 {
		mmCreateOrder.mock.t.Fatalf("Some expectations are already set for the ILOMSClient.CreateOrder method")
	}

	mmCreateOrder.mock.funcCreateOrder = f
	return mmCreateOrder.mock
}

// When sets expectation for the ILOMSClient.CreateOrder which will trigger the result defined by the following
// Then helper
func (mmCreateOrder *mILOMSClientMockCreateOrder) When(ctx context.Context, user int64, items []models.Item) *ILOMSClientMockCreateOrderExpectation {
	if mmCreateOrder.mock.funcCreateOrder != nil {
		mmCreateOrder.mock.t.Fatalf("ILOMSClientMock.CreateOrder mock is already set by Set")
	}

	expectation := &ILOMSClientMockCreateOrderExpectation{
		mock:   mmCreateOrder.mock,
		params: &ILOMSClientMockCreateOrderParams{ctx, user, items},
	}
	mmCreateOrder.expectations = append(mmCreateOrder.expectations, expectation)
	return expectation
}

// Then sets up ILOMSClient.CreateOrder return parameters for the expectation previously defined by the When method
func (e *ILOMSClientMockCreateOrderExpectation) Then(i1 int64, err error) *ILOMSClientMock {
	e.results = &ILOMSClientMockCreateOrderResults{i1, err}
	return e.mock
}

// CreateOrder implements loms.ILOMSClient
func (mmCreateOrder *ILOMSClientMock) CreateOrder(ctx context.Context, user int64, items []models.Item) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreateOrder.beforeCreateOrderCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateOrder.afterCreateOrderCounter, 1)

	if mmCreateOrder.inspectFuncCreateOrder != nil {
		mmCreateOrder.inspectFuncCreateOrder(ctx, user, items)
	}

	mm_params := &ILOMSClientMockCreateOrderParams{ctx, user, items}

	// Record call args
	mmCreateOrder.CreateOrderMock.mutex.Lock()
	mmCreateOrder.CreateOrderMock.callArgs = append(mmCreateOrder.CreateOrderMock.callArgs, mm_params)
	mmCreateOrder.CreateOrderMock.mutex.Unlock()

	for _, e := range mmCreateOrder.CreateOrderMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreateOrder.CreateOrderMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateOrder.CreateOrderMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateOrder.CreateOrderMock.defaultExpectation.params
		mm_got := ILOMSClientMockCreateOrderParams{ctx, user, items}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateOrder.t.Errorf("ILOMSClientMock.CreateOrder got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateOrder.CreateOrderMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateOrder.t.Fatal("No results are set for the ILOMSClientMock.CreateOrder")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreateOrder.funcCreateOrder != nil {
		return mmCreateOrder.funcCreateOrder(ctx, user, items)
	}
	mmCreateOrder.t.Fatalf("Unexpected call to ILOMSClientMock.CreateOrder. %v %v %v", ctx, user, items)
	return
}

// CreateOrderAfterCounter returns a count of finished ILOMSClientMock.CreateOrder invocations
func (mmCreateOrder *ILOMSClientMock) CreateOrderAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateOrder.afterCreateOrderCounter)
}

// CreateOrderBeforeCounter returns a count of ILOMSClientMock.CreateOrder invocations
func (mmCreateOrder *ILOMSClientMock) CreateOrderBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateOrder.beforeCreateOrderCounter)
}

// Calls returns a list of arguments used in each call to ILOMSClientMock.CreateOrder.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateOrder *mILOMSClientMockCreateOrder) Calls() []*ILOMSClientMockCreateOrderParams {
	mmCreateOrder.mutex.RLock()

	argCopy := make([]*ILOMSClientMockCreateOrderParams, len(mmCreateOrder.callArgs))
	copy(argCopy, mmCreateOrder.callArgs)

	mmCreateOrder.mutex.RUnlock()

	return argCopy
}

// MinimockCreateOrderDone returns true if the count of the CreateOrder invocations corresponds
// the number of defined expectations
func (m *ILOMSClientMock) MinimockCreateOrderDone() bool {
	for _, e := range m.CreateOrderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateOrderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateOrder != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateOrderInspect logs each unmet expectation
func (m *ILOMSClientMock) MinimockCreateOrderInspect() {
	for _, e := range m.CreateOrderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ILOMSClientMock.CreateOrder with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateOrderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		if m.CreateOrderMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ILOMSClientMock.CreateOrder")
		} else {
			m.t.Errorf("Expected call to ILOMSClientMock.CreateOrder with params: %#v", *m.CreateOrderMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateOrder != nil && mm_atomic.LoadUint64(&m.afterCreateOrderCounter) < 1 {
		m.t.Error("Expected call to ILOMSClientMock.CreateOrder")
	}
}

type mILOMSClientMockStocks struct {
	mock               *ILOMSClientMock
	defaultExpectation *ILOMSClientMockStocksExpectation
	expectations       []*ILOMSClientMockStocksExpectation

	callArgs []*ILOMSClientMockStocksParams
	mutex    sync.RWMutex
}

// ILOMSClientMockStocksExpectation specifies expectation struct of the ILOMSClient.Stocks
type ILOMSClientMockStocksExpectation struct {
	mock    *ILOMSClientMock
	params  *ILOMSClientMockStocksParams
	results *ILOMSClientMockStocksResults
	Counter uint64
}

// ILOMSClientMockStocksParams contains parameters of the ILOMSClient.Stocks
type ILOMSClientMockStocksParams struct {
	ctx context.Context
	sku uint32
}

// ILOMSClientMockStocksResults contains results of the ILOMSClient.Stocks
type ILOMSClientMockStocksResults struct {
	sa1 []models.Stock
	err error
}

// Expect sets up expected params for ILOMSClient.Stocks
func (mmStocks *mILOMSClientMockStocks) Expect(ctx context.Context, sku uint32) *mILOMSClientMockStocks {
	if mmStocks.mock.funcStocks != nil {
		mmStocks.mock.t.Fatalf("ILOMSClientMock.Stocks mock is already set by Set")
	}

	if mmStocks.defaultExpectation == nil {
		mmStocks.defaultExpectation = &ILOMSClientMockStocksExpectation{}
	}

	mmStocks.defaultExpectation.params = &ILOMSClientMockStocksParams{ctx, sku}
	for _, e := range mmStocks.expectations {
		if minimock.Equal(e.params, mmStocks.defaultExpectation.params) {
			mmStocks.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStocks.defaultExpectation.params)
		}
	}

	return mmStocks
}

// Inspect accepts an inspector function that has same arguments as the ILOMSClient.Stocks
func (mmStocks *mILOMSClientMockStocks) Inspect(f func(ctx context.Context, sku uint32)) *mILOMSClientMockStocks {
	if mmStocks.mock.inspectFuncStocks != nil {
		mmStocks.mock.t.Fatalf("Inspect function is already set for ILOMSClientMock.Stocks")
	}

	mmStocks.mock.inspectFuncStocks = f

	return mmStocks
}

// Return sets up results that will be returned by ILOMSClient.Stocks
func (mmStocks *mILOMSClientMockStocks) Return(sa1 []models.Stock, err error) *ILOMSClientMock {
	if mmStocks.mock.funcStocks != nil {
		mmStocks.mock.t.Fatalf("ILOMSClientMock.Stocks mock is already set by Set")
	}

	if mmStocks.defaultExpectation == nil {
		mmStocks.defaultExpectation = &ILOMSClientMockStocksExpectation{mock: mmStocks.mock}
	}
	mmStocks.defaultExpectation.results = &ILOMSClientMockStocksResults{sa1, err}
	return mmStocks.mock
}

// Set uses given function f to mock the ILOMSClient.Stocks method
func (mmStocks *mILOMSClientMockStocks) Set(f func(ctx context.Context, sku uint32) (sa1 []models.Stock, err error)) *ILOMSClientMock {
	if mmStocks.defaultExpectation != nil {
		mmStocks.mock.t.Fatalf("Default expectation is already set for the ILOMSClient.Stocks method")
	}

	if len(mmStocks.expectations) > 0 {
		mmStocks.mock.t.Fatalf("Some expectations are already set for the ILOMSClient.Stocks method")
	}

	mmStocks.mock.funcStocks = f
	return mmStocks.mock
}

// When sets expectation for the ILOMSClient.Stocks which will trigger the result defined by the following
// Then helper
func (mmStocks *mILOMSClientMockStocks) When(ctx context.Context, sku uint32) *ILOMSClientMockStocksExpectation {
	if mmStocks.mock.funcStocks != nil {
		mmStocks.mock.t.Fatalf("ILOMSClientMock.Stocks mock is already set by Set")
	}

	expectation := &ILOMSClientMockStocksExpectation{
		mock:   mmStocks.mock,
		params: &ILOMSClientMockStocksParams{ctx, sku},
	}
	mmStocks.expectations = append(mmStocks.expectations, expectation)
	return expectation
}

// Then sets up ILOMSClient.Stocks return parameters for the expectation previously defined by the When method
func (e *ILOMSClientMockStocksExpectation) Then(sa1 []models.Stock, err error) *ILOMSClientMock {
	e.results = &ILOMSClientMockStocksResults{sa1, err}
	return e.mock
}

// Stocks implements loms.ILOMSClient
func (mmStocks *ILOMSClientMock) Stocks(ctx context.Context, sku uint32) (sa1 []models.Stock, err error) {
	mm_atomic.AddUint64(&mmStocks.beforeStocksCounter, 1)
	defer mm_atomic.AddUint64(&mmStocks.afterStocksCounter, 1)

	if mmStocks.inspectFuncStocks != nil {
		mmStocks.inspectFuncStocks(ctx, sku)
	}

	mm_params := &ILOMSClientMockStocksParams{ctx, sku}

	// Record call args
	mmStocks.StocksMock.mutex.Lock()
	mmStocks.StocksMock.callArgs = append(mmStocks.StocksMock.callArgs, mm_params)
	mmStocks.StocksMock.mutex.Unlock()

	for _, e := range mmStocks.StocksMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sa1, e.results.err
		}
	}

	if mmStocks.StocksMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStocks.StocksMock.defaultExpectation.Counter, 1)
		mm_want := mmStocks.StocksMock.defaultExpectation.params
		mm_got := ILOMSClientMockStocksParams{ctx, sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStocks.t.Errorf("ILOMSClientMock.Stocks got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmStocks.StocksMock.defaultExpectation.results
		if mm_results == nil {
			mmStocks.t.Fatal("No results are set for the ILOMSClientMock.Stocks")
		}
		return (*mm_results).sa1, (*mm_results).err
	}
	if mmStocks.funcStocks != nil {
		return mmStocks.funcStocks(ctx, sku)
	}
	mmStocks.t.Fatalf("Unexpected call to ILOMSClientMock.Stocks. %v %v", ctx, sku)
	return
}

// StocksAfterCounter returns a count of finished ILOMSClientMock.Stocks invocations
func (mmStocks *ILOMSClientMock) StocksAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStocks.afterStocksCounter)
}

// StocksBeforeCounter returns a count of ILOMSClientMock.Stocks invocations
func (mmStocks *ILOMSClientMock) StocksBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStocks.beforeStocksCounter)
}

// Calls returns a list of arguments used in each call to ILOMSClientMock.Stocks.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStocks *mILOMSClientMockStocks) Calls() []*ILOMSClientMockStocksParams {
	mmStocks.mutex.RLock()

	argCopy := make([]*ILOMSClientMockStocksParams, len(mmStocks.callArgs))
	copy(argCopy, mmStocks.callArgs)

	mmStocks.mutex.RUnlock()

	return argCopy
}

// MinimockStocksDone returns true if the count of the Stocks invocations corresponds
// the number of defined expectations
func (m *ILOMSClientMock) MinimockStocksDone() bool {
	for _, e := range m.StocksMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StocksMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStocksCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStocks != nil && mm_atomic.LoadUint64(&m.afterStocksCounter) < 1 {
		return false
	}
	return true
}

// MinimockStocksInspect logs each unmet expectation
func (m *ILOMSClientMock) MinimockStocksInspect() {
	for _, e := range m.StocksMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ILOMSClientMock.Stocks with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StocksMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStocksCounter) < 1 {
		if m.StocksMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ILOMSClientMock.Stocks")
		} else {
			m.t.Errorf("Expected call to ILOMSClientMock.Stocks with params: %#v", *m.StocksMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStocks != nil && mm_atomic.LoadUint64(&m.afterStocksCounter) < 1 {
		m.t.Error("Expected call to ILOMSClientMock.Stocks")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ILOMSClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateOrderInspect()

		m.MinimockStocksInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ILOMSClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ILOMSClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateOrderDone() &&
		m.MinimockStocksDone()
}

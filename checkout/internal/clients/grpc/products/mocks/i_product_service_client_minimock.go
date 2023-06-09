package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/clients/grpc/products.IProductServiceClient -o ./mocks/i_product_service_client_minimock.go -n IProductServiceClientMock

import (
	"context"
	"route256/checkout/internal/models"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// IProductServiceClientMock implements products.IProductServiceClient
type IProductServiceClientMock struct {
	t minimock.Tester

	funcGetProduct          func(ctx context.Context, Sku uint32) (pp1 *models.ProductAttrs, err error)
	inspectFuncGetProduct   func(ctx context.Context, Sku uint32)
	afterGetProductCounter  uint64
	beforeGetProductCounter uint64
	GetProductMock          mIProductServiceClientMockGetProduct
}

// NewIProductServiceClientMock returns a mock for products.IProductServiceClient
func NewIProductServiceClientMock(t minimock.Tester) *IProductServiceClientMock {
	m := &IProductServiceClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetProductMock = mIProductServiceClientMockGetProduct{mock: m}
	m.GetProductMock.callArgs = []*IProductServiceClientMockGetProductParams{}

	return m
}

type mIProductServiceClientMockGetProduct struct {
	mock               *IProductServiceClientMock
	defaultExpectation *IProductServiceClientMockGetProductExpectation
	expectations       []*IProductServiceClientMockGetProductExpectation

	callArgs []*IProductServiceClientMockGetProductParams
	mutex    sync.RWMutex
}

// IProductServiceClientMockGetProductExpectation specifies expectation struct of the IProductServiceClient.GetProduct
type IProductServiceClientMockGetProductExpectation struct {
	mock    *IProductServiceClientMock
	params  *IProductServiceClientMockGetProductParams
	results *IProductServiceClientMockGetProductResults
	Counter uint64
}

// IProductServiceClientMockGetProductParams contains parameters of the IProductServiceClient.GetProduct
type IProductServiceClientMockGetProductParams struct {
	ctx context.Context
	Sku uint32
}

// IProductServiceClientMockGetProductResults contains results of the IProductServiceClient.GetProduct
type IProductServiceClientMockGetProductResults struct {
	pp1 *models.ProductAttrs
	err error
}

// Expect sets up expected params for IProductServiceClient.GetProduct
func (mmGetProduct *mIProductServiceClientMockGetProduct) Expect(ctx context.Context, Sku uint32) *mIProductServiceClientMockGetProduct {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("IProductServiceClientMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &IProductServiceClientMockGetProductExpectation{}
	}

	mmGetProduct.defaultExpectation.params = &IProductServiceClientMockGetProductParams{ctx, Sku}
	for _, e := range mmGetProduct.expectations {
		if minimock.Equal(e.params, mmGetProduct.defaultExpectation.params) {
			mmGetProduct.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetProduct.defaultExpectation.params)
		}
	}

	return mmGetProduct
}

// Inspect accepts an inspector function that has same arguments as the IProductServiceClient.GetProduct
func (mmGetProduct *mIProductServiceClientMockGetProduct) Inspect(f func(ctx context.Context, Sku uint32)) *mIProductServiceClientMockGetProduct {
	if mmGetProduct.mock.inspectFuncGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("Inspect function is already set for IProductServiceClientMock.GetProduct")
	}

	mmGetProduct.mock.inspectFuncGetProduct = f

	return mmGetProduct
}

// Return sets up results that will be returned by IProductServiceClient.GetProduct
func (mmGetProduct *mIProductServiceClientMockGetProduct) Return(pp1 *models.ProductAttrs, err error) *IProductServiceClientMock {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("IProductServiceClientMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &IProductServiceClientMockGetProductExpectation{mock: mmGetProduct.mock}
	}
	mmGetProduct.defaultExpectation.results = &IProductServiceClientMockGetProductResults{pp1, err}
	return mmGetProduct.mock
}

// Set uses given function f to mock the IProductServiceClient.GetProduct method
func (mmGetProduct *mIProductServiceClientMockGetProduct) Set(f func(ctx context.Context, Sku uint32) (pp1 *models.ProductAttrs, err error)) *IProductServiceClientMock {
	if mmGetProduct.defaultExpectation != nil {
		mmGetProduct.mock.t.Fatalf("Default expectation is already set for the IProductServiceClient.GetProduct method")
	}

	if len(mmGetProduct.expectations) > 0 {
		mmGetProduct.mock.t.Fatalf("Some expectations are already set for the IProductServiceClient.GetProduct method")
	}

	mmGetProduct.mock.funcGetProduct = f
	return mmGetProduct.mock
}

// When sets expectation for the IProductServiceClient.GetProduct which will trigger the result defined by the following
// Then helper
func (mmGetProduct *mIProductServiceClientMockGetProduct) When(ctx context.Context, Sku uint32) *IProductServiceClientMockGetProductExpectation {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("IProductServiceClientMock.GetProduct mock is already set by Set")
	}

	expectation := &IProductServiceClientMockGetProductExpectation{
		mock:   mmGetProduct.mock,
		params: &IProductServiceClientMockGetProductParams{ctx, Sku},
	}
	mmGetProduct.expectations = append(mmGetProduct.expectations, expectation)
	return expectation
}

// Then sets up IProductServiceClient.GetProduct return parameters for the expectation previously defined by the When method
func (e *IProductServiceClientMockGetProductExpectation) Then(pp1 *models.ProductAttrs, err error) *IProductServiceClientMock {
	e.results = &IProductServiceClientMockGetProductResults{pp1, err}
	return e.mock
}

// GetProduct implements products.IProductServiceClient
func (mmGetProduct *IProductServiceClientMock) GetProduct(ctx context.Context, Sku uint32) (pp1 *models.ProductAttrs, err error) {
	mm_atomic.AddUint64(&mmGetProduct.beforeGetProductCounter, 1)
	defer mm_atomic.AddUint64(&mmGetProduct.afterGetProductCounter, 1)

	if mmGetProduct.inspectFuncGetProduct != nil {
		mmGetProduct.inspectFuncGetProduct(ctx, Sku)
	}

	mm_params := &IProductServiceClientMockGetProductParams{ctx, Sku}

	// Record call args
	mmGetProduct.GetProductMock.mutex.Lock()
	mmGetProduct.GetProductMock.callArgs = append(mmGetProduct.GetProductMock.callArgs, mm_params)
	mmGetProduct.GetProductMock.mutex.Unlock()

	for _, e := range mmGetProduct.GetProductMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp1, e.results.err
		}
	}

	if mmGetProduct.GetProductMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetProduct.GetProductMock.defaultExpectation.Counter, 1)
		mm_want := mmGetProduct.GetProductMock.defaultExpectation.params
		mm_got := IProductServiceClientMockGetProductParams{ctx, Sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetProduct.t.Errorf("IProductServiceClientMock.GetProduct got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetProduct.GetProductMock.defaultExpectation.results
		if mm_results == nil {
			mmGetProduct.t.Fatal("No results are set for the IProductServiceClientMock.GetProduct")
		}
		return (*mm_results).pp1, (*mm_results).err
	}
	if mmGetProduct.funcGetProduct != nil {
		return mmGetProduct.funcGetProduct(ctx, Sku)
	}
	mmGetProduct.t.Fatalf("Unexpected call to IProductServiceClientMock.GetProduct. %v %v", ctx, Sku)
	return
}

// GetProductAfterCounter returns a count of finished IProductServiceClientMock.GetProduct invocations
func (mmGetProduct *IProductServiceClientMock) GetProductAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.afterGetProductCounter)
}

// GetProductBeforeCounter returns a count of IProductServiceClientMock.GetProduct invocations
func (mmGetProduct *IProductServiceClientMock) GetProductBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.beforeGetProductCounter)
}

// Calls returns a list of arguments used in each call to IProductServiceClientMock.GetProduct.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetProduct *mIProductServiceClientMockGetProduct) Calls() []*IProductServiceClientMockGetProductParams {
	mmGetProduct.mutex.RLock()

	argCopy := make([]*IProductServiceClientMockGetProductParams, len(mmGetProduct.callArgs))
	copy(argCopy, mmGetProduct.callArgs)

	mmGetProduct.mutex.RUnlock()

	return argCopy
}

// MinimockGetProductDone returns true if the count of the GetProduct invocations corresponds
// the number of defined expectations
func (m *IProductServiceClientMock) MinimockGetProductDone() bool {
	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetProduct != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetProductInspect logs each unmet expectation
func (m *IProductServiceClientMock) MinimockGetProductInspect() {
	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IProductServiceClientMock.GetProduct with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		if m.GetProductMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IProductServiceClientMock.GetProduct")
		} else {
			m.t.Errorf("Expected call to IProductServiceClientMock.GetProduct with params: %#v", *m.GetProductMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetProduct != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		m.t.Error("Expected call to IProductServiceClientMock.GetProduct")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *IProductServiceClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetProductInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *IProductServiceClientMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *IProductServiceClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetProductDone()
}

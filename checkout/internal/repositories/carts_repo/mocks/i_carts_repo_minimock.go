package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/repositories/carts_repo.ICartsRepo -o ./mocks/i_carts_repo_minimock.go -n ICartsRepoMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ICartsRepoMock implements carts_repo.ICartsRepo
type ICartsRepoMock struct {
	t minimock.Tester

	funcCreateCart          func(ctx context.Context, UserID int64) (u1 uint64, err error)
	inspectFuncCreateCart   func(ctx context.Context, UserID int64)
	afterCreateCartCounter  uint64
	beforeCreateCartCounter uint64
	CreateCartMock          mICartsRepoMockCreateCart

	funcGetCartID          func(ctx context.Context, userID int64) (u1 uint64, err error)
	inspectFuncGetCartID   func(ctx context.Context, userID int64)
	afterGetCartIDCounter  uint64
	beforeGetCartIDCounter uint64
	GetCartIDMock          mICartsRepoMockGetCartID

	funcPurchaseCart          func(ctx context.Context, userID int64) (u1 uint64, err error)
	inspectFuncPurchaseCart   func(ctx context.Context, userID int64)
	afterPurchaseCartCounter  uint64
	beforePurchaseCartCounter uint64
	PurchaseCartMock          mICartsRepoMockPurchaseCart
}

// NewICartsRepoMock returns a mock for carts_repo.ICartsRepo
func NewICartsRepoMock(t minimock.Tester) *ICartsRepoMock {
	m := &ICartsRepoMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateCartMock = mICartsRepoMockCreateCart{mock: m}
	m.CreateCartMock.callArgs = []*ICartsRepoMockCreateCartParams{}

	m.GetCartIDMock = mICartsRepoMockGetCartID{mock: m}
	m.GetCartIDMock.callArgs = []*ICartsRepoMockGetCartIDParams{}

	m.PurchaseCartMock = mICartsRepoMockPurchaseCart{mock: m}
	m.PurchaseCartMock.callArgs = []*ICartsRepoMockPurchaseCartParams{}

	return m
}

type mICartsRepoMockCreateCart struct {
	mock               *ICartsRepoMock
	defaultExpectation *ICartsRepoMockCreateCartExpectation
	expectations       []*ICartsRepoMockCreateCartExpectation

	callArgs []*ICartsRepoMockCreateCartParams
	mutex    sync.RWMutex
}

// ICartsRepoMockCreateCartExpectation specifies expectation struct of the ICartsRepo.CreateCart
type ICartsRepoMockCreateCartExpectation struct {
	mock    *ICartsRepoMock
	params  *ICartsRepoMockCreateCartParams
	results *ICartsRepoMockCreateCartResults
	Counter uint64
}

// ICartsRepoMockCreateCartParams contains parameters of the ICartsRepo.CreateCart
type ICartsRepoMockCreateCartParams struct {
	ctx    context.Context
	UserID int64
}

// ICartsRepoMockCreateCartResults contains results of the ICartsRepo.CreateCart
type ICartsRepoMockCreateCartResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for ICartsRepo.CreateCart
func (mmCreateCart *mICartsRepoMockCreateCart) Expect(ctx context.Context, UserID int64) *mICartsRepoMockCreateCart {
	if mmCreateCart.mock.funcCreateCart != nil {
		mmCreateCart.mock.t.Fatalf("ICartsRepoMock.CreateCart mock is already set by Set")
	}

	if mmCreateCart.defaultExpectation == nil {
		mmCreateCart.defaultExpectation = &ICartsRepoMockCreateCartExpectation{}
	}

	mmCreateCart.defaultExpectation.params = &ICartsRepoMockCreateCartParams{ctx, UserID}
	for _, e := range mmCreateCart.expectations {
		if minimock.Equal(e.params, mmCreateCart.defaultExpectation.params) {
			mmCreateCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateCart.defaultExpectation.params)
		}
	}

	return mmCreateCart
}

// Inspect accepts an inspector function that has same arguments as the ICartsRepo.CreateCart
func (mmCreateCart *mICartsRepoMockCreateCart) Inspect(f func(ctx context.Context, UserID int64)) *mICartsRepoMockCreateCart {
	if mmCreateCart.mock.inspectFuncCreateCart != nil {
		mmCreateCart.mock.t.Fatalf("Inspect function is already set for ICartsRepoMock.CreateCart")
	}

	mmCreateCart.mock.inspectFuncCreateCart = f

	return mmCreateCart
}

// Return sets up results that will be returned by ICartsRepo.CreateCart
func (mmCreateCart *mICartsRepoMockCreateCart) Return(u1 uint64, err error) *ICartsRepoMock {
	if mmCreateCart.mock.funcCreateCart != nil {
		mmCreateCart.mock.t.Fatalf("ICartsRepoMock.CreateCart mock is already set by Set")
	}

	if mmCreateCart.defaultExpectation == nil {
		mmCreateCart.defaultExpectation = &ICartsRepoMockCreateCartExpectation{mock: mmCreateCart.mock}
	}
	mmCreateCart.defaultExpectation.results = &ICartsRepoMockCreateCartResults{u1, err}
	return mmCreateCart.mock
}

// Set uses given function f to mock the ICartsRepo.CreateCart method
func (mmCreateCart *mICartsRepoMockCreateCart) Set(f func(ctx context.Context, UserID int64) (u1 uint64, err error)) *ICartsRepoMock {
	if mmCreateCart.defaultExpectation != nil {
		mmCreateCart.mock.t.Fatalf("Default expectation is already set for the ICartsRepo.CreateCart method")
	}

	if len(mmCreateCart.expectations) > 0 {
		mmCreateCart.mock.t.Fatalf("Some expectations are already set for the ICartsRepo.CreateCart method")
	}

	mmCreateCart.mock.funcCreateCart = f
	return mmCreateCart.mock
}

// When sets expectation for the ICartsRepo.CreateCart which will trigger the result defined by the following
// Then helper
func (mmCreateCart *mICartsRepoMockCreateCart) When(ctx context.Context, UserID int64) *ICartsRepoMockCreateCartExpectation {
	if mmCreateCart.mock.funcCreateCart != nil {
		mmCreateCart.mock.t.Fatalf("ICartsRepoMock.CreateCart mock is already set by Set")
	}

	expectation := &ICartsRepoMockCreateCartExpectation{
		mock:   mmCreateCart.mock,
		params: &ICartsRepoMockCreateCartParams{ctx, UserID},
	}
	mmCreateCart.expectations = append(mmCreateCart.expectations, expectation)
	return expectation
}

// Then sets up ICartsRepo.CreateCart return parameters for the expectation previously defined by the When method
func (e *ICartsRepoMockCreateCartExpectation) Then(u1 uint64, err error) *ICartsRepoMock {
	e.results = &ICartsRepoMockCreateCartResults{u1, err}
	return e.mock
}

// CreateCart implements carts_repo.ICartsRepo
func (mmCreateCart *ICartsRepoMock) CreateCart(ctx context.Context, UserID int64) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmCreateCart.beforeCreateCartCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateCart.afterCreateCartCounter, 1)

	if mmCreateCart.inspectFuncCreateCart != nil {
		mmCreateCart.inspectFuncCreateCart(ctx, UserID)
	}

	mm_params := &ICartsRepoMockCreateCartParams{ctx, UserID}

	// Record call args
	mmCreateCart.CreateCartMock.mutex.Lock()
	mmCreateCart.CreateCartMock.callArgs = append(mmCreateCart.CreateCartMock.callArgs, mm_params)
	mmCreateCart.CreateCartMock.mutex.Unlock()

	for _, e := range mmCreateCart.CreateCartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmCreateCart.CreateCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateCart.CreateCartMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateCart.CreateCartMock.defaultExpectation.params
		mm_got := ICartsRepoMockCreateCartParams{ctx, UserID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateCart.t.Errorf("ICartsRepoMock.CreateCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateCart.CreateCartMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateCart.t.Fatal("No results are set for the ICartsRepoMock.CreateCart")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmCreateCart.funcCreateCart != nil {
		return mmCreateCart.funcCreateCart(ctx, UserID)
	}
	mmCreateCart.t.Fatalf("Unexpected call to ICartsRepoMock.CreateCart. %v %v", ctx, UserID)
	return
}

// CreateCartAfterCounter returns a count of finished ICartsRepoMock.CreateCart invocations
func (mmCreateCart *ICartsRepoMock) CreateCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateCart.afterCreateCartCounter)
}

// CreateCartBeforeCounter returns a count of ICartsRepoMock.CreateCart invocations
func (mmCreateCart *ICartsRepoMock) CreateCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateCart.beforeCreateCartCounter)
}

// Calls returns a list of arguments used in each call to ICartsRepoMock.CreateCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateCart *mICartsRepoMockCreateCart) Calls() []*ICartsRepoMockCreateCartParams {
	mmCreateCart.mutex.RLock()

	argCopy := make([]*ICartsRepoMockCreateCartParams, len(mmCreateCart.callArgs))
	copy(argCopy, mmCreateCart.callArgs)

	mmCreateCart.mutex.RUnlock()

	return argCopy
}

// MinimockCreateCartDone returns true if the count of the CreateCart invocations corresponds
// the number of defined expectations
func (m *ICartsRepoMock) MinimockCreateCartDone() bool {
	for _, e := range m.CreateCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateCart != nil && mm_atomic.LoadUint64(&m.afterCreateCartCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateCartInspect logs each unmet expectation
func (m *ICartsRepoMock) MinimockCreateCartInspect() {
	for _, e := range m.CreateCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ICartsRepoMock.CreateCart with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCartCounter) < 1 {
		if m.CreateCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ICartsRepoMock.CreateCart")
		} else {
			m.t.Errorf("Expected call to ICartsRepoMock.CreateCart with params: %#v", *m.CreateCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateCart != nil && mm_atomic.LoadUint64(&m.afterCreateCartCounter) < 1 {
		m.t.Error("Expected call to ICartsRepoMock.CreateCart")
	}
}

type mICartsRepoMockGetCartID struct {
	mock               *ICartsRepoMock
	defaultExpectation *ICartsRepoMockGetCartIDExpectation
	expectations       []*ICartsRepoMockGetCartIDExpectation

	callArgs []*ICartsRepoMockGetCartIDParams
	mutex    sync.RWMutex
}

// ICartsRepoMockGetCartIDExpectation specifies expectation struct of the ICartsRepo.GetCartID
type ICartsRepoMockGetCartIDExpectation struct {
	mock    *ICartsRepoMock
	params  *ICartsRepoMockGetCartIDParams
	results *ICartsRepoMockGetCartIDResults
	Counter uint64
}

// ICartsRepoMockGetCartIDParams contains parameters of the ICartsRepo.GetCartID
type ICartsRepoMockGetCartIDParams struct {
	ctx    context.Context
	userID int64
}

// ICartsRepoMockGetCartIDResults contains results of the ICartsRepo.GetCartID
type ICartsRepoMockGetCartIDResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for ICartsRepo.GetCartID
func (mmGetCartID *mICartsRepoMockGetCartID) Expect(ctx context.Context, userID int64) *mICartsRepoMockGetCartID {
	if mmGetCartID.mock.funcGetCartID != nil {
		mmGetCartID.mock.t.Fatalf("ICartsRepoMock.GetCartID mock is already set by Set")
	}

	if mmGetCartID.defaultExpectation == nil {
		mmGetCartID.defaultExpectation = &ICartsRepoMockGetCartIDExpectation{}
	}

	mmGetCartID.defaultExpectation.params = &ICartsRepoMockGetCartIDParams{ctx, userID}
	for _, e := range mmGetCartID.expectations {
		if minimock.Equal(e.params, mmGetCartID.defaultExpectation.params) {
			mmGetCartID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetCartID.defaultExpectation.params)
		}
	}

	return mmGetCartID
}

// Inspect accepts an inspector function that has same arguments as the ICartsRepo.GetCartID
func (mmGetCartID *mICartsRepoMockGetCartID) Inspect(f func(ctx context.Context, userID int64)) *mICartsRepoMockGetCartID {
	if mmGetCartID.mock.inspectFuncGetCartID != nil {
		mmGetCartID.mock.t.Fatalf("Inspect function is already set for ICartsRepoMock.GetCartID")
	}

	mmGetCartID.mock.inspectFuncGetCartID = f

	return mmGetCartID
}

// Return sets up results that will be returned by ICartsRepo.GetCartID
func (mmGetCartID *mICartsRepoMockGetCartID) Return(u1 uint64, err error) *ICartsRepoMock {
	if mmGetCartID.mock.funcGetCartID != nil {
		mmGetCartID.mock.t.Fatalf("ICartsRepoMock.GetCartID mock is already set by Set")
	}

	if mmGetCartID.defaultExpectation == nil {
		mmGetCartID.defaultExpectation = &ICartsRepoMockGetCartIDExpectation{mock: mmGetCartID.mock}
	}
	mmGetCartID.defaultExpectation.results = &ICartsRepoMockGetCartIDResults{u1, err}
	return mmGetCartID.mock
}

// Set uses given function f to mock the ICartsRepo.GetCartID method
func (mmGetCartID *mICartsRepoMockGetCartID) Set(f func(ctx context.Context, userID int64) (u1 uint64, err error)) *ICartsRepoMock {
	if mmGetCartID.defaultExpectation != nil {
		mmGetCartID.mock.t.Fatalf("Default expectation is already set for the ICartsRepo.GetCartID method")
	}

	if len(mmGetCartID.expectations) > 0 {
		mmGetCartID.mock.t.Fatalf("Some expectations are already set for the ICartsRepo.GetCartID method")
	}

	mmGetCartID.mock.funcGetCartID = f
	return mmGetCartID.mock
}

// When sets expectation for the ICartsRepo.GetCartID which will trigger the result defined by the following
// Then helper
func (mmGetCartID *mICartsRepoMockGetCartID) When(ctx context.Context, userID int64) *ICartsRepoMockGetCartIDExpectation {
	if mmGetCartID.mock.funcGetCartID != nil {
		mmGetCartID.mock.t.Fatalf("ICartsRepoMock.GetCartID mock is already set by Set")
	}

	expectation := &ICartsRepoMockGetCartIDExpectation{
		mock:   mmGetCartID.mock,
		params: &ICartsRepoMockGetCartIDParams{ctx, userID},
	}
	mmGetCartID.expectations = append(mmGetCartID.expectations, expectation)
	return expectation
}

// Then sets up ICartsRepo.GetCartID return parameters for the expectation previously defined by the When method
func (e *ICartsRepoMockGetCartIDExpectation) Then(u1 uint64, err error) *ICartsRepoMock {
	e.results = &ICartsRepoMockGetCartIDResults{u1, err}
	return e.mock
}

// GetCartID implements carts_repo.ICartsRepo
func (mmGetCartID *ICartsRepoMock) GetCartID(ctx context.Context, userID int64) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmGetCartID.beforeGetCartIDCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCartID.afterGetCartIDCounter, 1)

	if mmGetCartID.inspectFuncGetCartID != nil {
		mmGetCartID.inspectFuncGetCartID(ctx, userID)
	}

	mm_params := &ICartsRepoMockGetCartIDParams{ctx, userID}

	// Record call args
	mmGetCartID.GetCartIDMock.mutex.Lock()
	mmGetCartID.GetCartIDMock.callArgs = append(mmGetCartID.GetCartIDMock.callArgs, mm_params)
	mmGetCartID.GetCartIDMock.mutex.Unlock()

	for _, e := range mmGetCartID.GetCartIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmGetCartID.GetCartIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCartID.GetCartIDMock.defaultExpectation.Counter, 1)
		mm_want := mmGetCartID.GetCartIDMock.defaultExpectation.params
		mm_got := ICartsRepoMockGetCartIDParams{ctx, userID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetCartID.t.Errorf("ICartsRepoMock.GetCartID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetCartID.GetCartIDMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCartID.t.Fatal("No results are set for the ICartsRepoMock.GetCartID")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmGetCartID.funcGetCartID != nil {
		return mmGetCartID.funcGetCartID(ctx, userID)
	}
	mmGetCartID.t.Fatalf("Unexpected call to ICartsRepoMock.GetCartID. %v %v", ctx, userID)
	return
}

// GetCartIDAfterCounter returns a count of finished ICartsRepoMock.GetCartID invocations
func (mmGetCartID *ICartsRepoMock) GetCartIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCartID.afterGetCartIDCounter)
}

// GetCartIDBeforeCounter returns a count of ICartsRepoMock.GetCartID invocations
func (mmGetCartID *ICartsRepoMock) GetCartIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCartID.beforeGetCartIDCounter)
}

// Calls returns a list of arguments used in each call to ICartsRepoMock.GetCartID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetCartID *mICartsRepoMockGetCartID) Calls() []*ICartsRepoMockGetCartIDParams {
	mmGetCartID.mutex.RLock()

	argCopy := make([]*ICartsRepoMockGetCartIDParams, len(mmGetCartID.callArgs))
	copy(argCopy, mmGetCartID.callArgs)

	mmGetCartID.mutex.RUnlock()

	return argCopy
}

// MinimockGetCartIDDone returns true if the count of the GetCartID invocations corresponds
// the number of defined expectations
func (m *ICartsRepoMock) MinimockGetCartIDDone() bool {
	for _, e := range m.GetCartIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCartIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCartIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCartID != nil && mm_atomic.LoadUint64(&m.afterGetCartIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetCartIDInspect logs each unmet expectation
func (m *ICartsRepoMock) MinimockGetCartIDInspect() {
	for _, e := range m.GetCartIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ICartsRepoMock.GetCartID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCartIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCartIDCounter) < 1 {
		if m.GetCartIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ICartsRepoMock.GetCartID")
		} else {
			m.t.Errorf("Expected call to ICartsRepoMock.GetCartID with params: %#v", *m.GetCartIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCartID != nil && mm_atomic.LoadUint64(&m.afterGetCartIDCounter) < 1 {
		m.t.Error("Expected call to ICartsRepoMock.GetCartID")
	}
}

type mICartsRepoMockPurchaseCart struct {
	mock               *ICartsRepoMock
	defaultExpectation *ICartsRepoMockPurchaseCartExpectation
	expectations       []*ICartsRepoMockPurchaseCartExpectation

	callArgs []*ICartsRepoMockPurchaseCartParams
	mutex    sync.RWMutex
}

// ICartsRepoMockPurchaseCartExpectation specifies expectation struct of the ICartsRepo.PurchaseCart
type ICartsRepoMockPurchaseCartExpectation struct {
	mock    *ICartsRepoMock
	params  *ICartsRepoMockPurchaseCartParams
	results *ICartsRepoMockPurchaseCartResults
	Counter uint64
}

// ICartsRepoMockPurchaseCartParams contains parameters of the ICartsRepo.PurchaseCart
type ICartsRepoMockPurchaseCartParams struct {
	ctx    context.Context
	userID int64
}

// ICartsRepoMockPurchaseCartResults contains results of the ICartsRepo.PurchaseCart
type ICartsRepoMockPurchaseCartResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for ICartsRepo.PurchaseCart
func (mmPurchaseCart *mICartsRepoMockPurchaseCart) Expect(ctx context.Context, userID int64) *mICartsRepoMockPurchaseCart {
	if mmPurchaseCart.mock.funcPurchaseCart != nil {
		mmPurchaseCart.mock.t.Fatalf("ICartsRepoMock.PurchaseCart mock is already set by Set")
	}

	if mmPurchaseCart.defaultExpectation == nil {
		mmPurchaseCart.defaultExpectation = &ICartsRepoMockPurchaseCartExpectation{}
	}

	mmPurchaseCart.defaultExpectation.params = &ICartsRepoMockPurchaseCartParams{ctx, userID}
	for _, e := range mmPurchaseCart.expectations {
		if minimock.Equal(e.params, mmPurchaseCart.defaultExpectation.params) {
			mmPurchaseCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPurchaseCart.defaultExpectation.params)
		}
	}

	return mmPurchaseCart
}

// Inspect accepts an inspector function that has same arguments as the ICartsRepo.PurchaseCart
func (mmPurchaseCart *mICartsRepoMockPurchaseCart) Inspect(f func(ctx context.Context, userID int64)) *mICartsRepoMockPurchaseCart {
	if mmPurchaseCart.mock.inspectFuncPurchaseCart != nil {
		mmPurchaseCart.mock.t.Fatalf("Inspect function is already set for ICartsRepoMock.PurchaseCart")
	}

	mmPurchaseCart.mock.inspectFuncPurchaseCart = f

	return mmPurchaseCart
}

// Return sets up results that will be returned by ICartsRepo.PurchaseCart
func (mmPurchaseCart *mICartsRepoMockPurchaseCart) Return(u1 uint64, err error) *ICartsRepoMock {
	if mmPurchaseCart.mock.funcPurchaseCart != nil {
		mmPurchaseCart.mock.t.Fatalf("ICartsRepoMock.PurchaseCart mock is already set by Set")
	}

	if mmPurchaseCart.defaultExpectation == nil {
		mmPurchaseCart.defaultExpectation = &ICartsRepoMockPurchaseCartExpectation{mock: mmPurchaseCart.mock}
	}
	mmPurchaseCart.defaultExpectation.results = &ICartsRepoMockPurchaseCartResults{u1, err}
	return mmPurchaseCart.mock
}

// Set uses given function f to mock the ICartsRepo.PurchaseCart method
func (mmPurchaseCart *mICartsRepoMockPurchaseCart) Set(f func(ctx context.Context, userID int64) (u1 uint64, err error)) *ICartsRepoMock {
	if mmPurchaseCart.defaultExpectation != nil {
		mmPurchaseCart.mock.t.Fatalf("Default expectation is already set for the ICartsRepo.PurchaseCart method")
	}

	if len(mmPurchaseCart.expectations) > 0 {
		mmPurchaseCart.mock.t.Fatalf("Some expectations are already set for the ICartsRepo.PurchaseCart method")
	}

	mmPurchaseCart.mock.funcPurchaseCart = f
	return mmPurchaseCart.mock
}

// When sets expectation for the ICartsRepo.PurchaseCart which will trigger the result defined by the following
// Then helper
func (mmPurchaseCart *mICartsRepoMockPurchaseCart) When(ctx context.Context, userID int64) *ICartsRepoMockPurchaseCartExpectation {
	if mmPurchaseCart.mock.funcPurchaseCart != nil {
		mmPurchaseCart.mock.t.Fatalf("ICartsRepoMock.PurchaseCart mock is already set by Set")
	}

	expectation := &ICartsRepoMockPurchaseCartExpectation{
		mock:   mmPurchaseCart.mock,
		params: &ICartsRepoMockPurchaseCartParams{ctx, userID},
	}
	mmPurchaseCart.expectations = append(mmPurchaseCart.expectations, expectation)
	return expectation
}

// Then sets up ICartsRepo.PurchaseCart return parameters for the expectation previously defined by the When method
func (e *ICartsRepoMockPurchaseCartExpectation) Then(u1 uint64, err error) *ICartsRepoMock {
	e.results = &ICartsRepoMockPurchaseCartResults{u1, err}
	return e.mock
}

// PurchaseCart implements carts_repo.ICartsRepo
func (mmPurchaseCart *ICartsRepoMock) PurchaseCart(ctx context.Context, userID int64) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmPurchaseCart.beforePurchaseCartCounter, 1)
	defer mm_atomic.AddUint64(&mmPurchaseCart.afterPurchaseCartCounter, 1)

	if mmPurchaseCart.inspectFuncPurchaseCart != nil {
		mmPurchaseCart.inspectFuncPurchaseCart(ctx, userID)
	}

	mm_params := &ICartsRepoMockPurchaseCartParams{ctx, userID}

	// Record call args
	mmPurchaseCart.PurchaseCartMock.mutex.Lock()
	mmPurchaseCart.PurchaseCartMock.callArgs = append(mmPurchaseCart.PurchaseCartMock.callArgs, mm_params)
	mmPurchaseCart.PurchaseCartMock.mutex.Unlock()

	for _, e := range mmPurchaseCart.PurchaseCartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmPurchaseCart.PurchaseCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPurchaseCart.PurchaseCartMock.defaultExpectation.Counter, 1)
		mm_want := mmPurchaseCart.PurchaseCartMock.defaultExpectation.params
		mm_got := ICartsRepoMockPurchaseCartParams{ctx, userID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPurchaseCart.t.Errorf("ICartsRepoMock.PurchaseCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPurchaseCart.PurchaseCartMock.defaultExpectation.results
		if mm_results == nil {
			mmPurchaseCart.t.Fatal("No results are set for the ICartsRepoMock.PurchaseCart")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmPurchaseCart.funcPurchaseCart != nil {
		return mmPurchaseCart.funcPurchaseCart(ctx, userID)
	}
	mmPurchaseCart.t.Fatalf("Unexpected call to ICartsRepoMock.PurchaseCart. %v %v", ctx, userID)
	return
}

// PurchaseCartAfterCounter returns a count of finished ICartsRepoMock.PurchaseCart invocations
func (mmPurchaseCart *ICartsRepoMock) PurchaseCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPurchaseCart.afterPurchaseCartCounter)
}

// PurchaseCartBeforeCounter returns a count of ICartsRepoMock.PurchaseCart invocations
func (mmPurchaseCart *ICartsRepoMock) PurchaseCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPurchaseCart.beforePurchaseCartCounter)
}

// Calls returns a list of arguments used in each call to ICartsRepoMock.PurchaseCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPurchaseCart *mICartsRepoMockPurchaseCart) Calls() []*ICartsRepoMockPurchaseCartParams {
	mmPurchaseCart.mutex.RLock()

	argCopy := make([]*ICartsRepoMockPurchaseCartParams, len(mmPurchaseCart.callArgs))
	copy(argCopy, mmPurchaseCart.callArgs)

	mmPurchaseCart.mutex.RUnlock()

	return argCopy
}

// MinimockPurchaseCartDone returns true if the count of the PurchaseCart invocations corresponds
// the number of defined expectations
func (m *ICartsRepoMock) MinimockPurchaseCartDone() bool {
	for _, e := range m.PurchaseCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PurchaseCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPurchaseCartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPurchaseCart != nil && mm_atomic.LoadUint64(&m.afterPurchaseCartCounter) < 1 {
		return false
	}
	return true
}

// MinimockPurchaseCartInspect logs each unmet expectation
func (m *ICartsRepoMock) MinimockPurchaseCartInspect() {
	for _, e := range m.PurchaseCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ICartsRepoMock.PurchaseCart with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PurchaseCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPurchaseCartCounter) < 1 {
		if m.PurchaseCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ICartsRepoMock.PurchaseCart")
		} else {
			m.t.Errorf("Expected call to ICartsRepoMock.PurchaseCart with params: %#v", *m.PurchaseCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPurchaseCart != nil && mm_atomic.LoadUint64(&m.afterPurchaseCartCounter) < 1 {
		m.t.Error("Expected call to ICartsRepoMock.PurchaseCart")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ICartsRepoMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateCartInspect()

		m.MinimockGetCartIDInspect()

		m.MinimockPurchaseCartInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ICartsRepoMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ICartsRepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateCartDone() &&
		m.MinimockGetCartIDDone() &&
		m.MinimockPurchaseCartDone()
}

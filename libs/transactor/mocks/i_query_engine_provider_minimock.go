package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/libs/transactor.IQueryEngineProvider -o ./mocks/i_query_engine_provider_minimock.go -n IQueryEngineProviderMock

import (
	"context"
	mm_transactor "route256/libs/transactor"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// IQueryEngineProviderMock implements transactor.IQueryEngineProvider
type IQueryEngineProviderMock struct {
	t minimock.Tester

	funcGetQueryEngine          func(ctx context.Context) (i1 mm_transactor.IQueryEngine)
	inspectFuncGetQueryEngine   func(ctx context.Context)
	afterGetQueryEngineCounter  uint64
	beforeGetQueryEngineCounter uint64
	GetQueryEngineMock          mIQueryEngineProviderMockGetQueryEngine
}

// NewIQueryEngineProviderMock returns a mock for transactor.IQueryEngineProvider
func NewIQueryEngineProviderMock(t minimock.Tester) *IQueryEngineProviderMock {
	m := &IQueryEngineProviderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetQueryEngineMock = mIQueryEngineProviderMockGetQueryEngine{mock: m}
	m.GetQueryEngineMock.callArgs = []*IQueryEngineProviderMockGetQueryEngineParams{}

	return m
}

type mIQueryEngineProviderMockGetQueryEngine struct {
	mock               *IQueryEngineProviderMock
	defaultExpectation *IQueryEngineProviderMockGetQueryEngineExpectation
	expectations       []*IQueryEngineProviderMockGetQueryEngineExpectation

	callArgs []*IQueryEngineProviderMockGetQueryEngineParams
	mutex    sync.RWMutex
}

// IQueryEngineProviderMockGetQueryEngineExpectation specifies expectation struct of the IQueryEngineProvider.GetQueryEngine
type IQueryEngineProviderMockGetQueryEngineExpectation struct {
	mock    *IQueryEngineProviderMock
	params  *IQueryEngineProviderMockGetQueryEngineParams
	results *IQueryEngineProviderMockGetQueryEngineResults
	Counter uint64
}

// IQueryEngineProviderMockGetQueryEngineParams contains parameters of the IQueryEngineProvider.GetQueryEngine
type IQueryEngineProviderMockGetQueryEngineParams struct {
	ctx context.Context
}

// IQueryEngineProviderMockGetQueryEngineResults contains results of the IQueryEngineProvider.GetQueryEngine
type IQueryEngineProviderMockGetQueryEngineResults struct {
	i1 mm_transactor.IQueryEngine
}

// Expect sets up expected params for IQueryEngineProvider.GetQueryEngine
func (mmGetQueryEngine *mIQueryEngineProviderMockGetQueryEngine) Expect(ctx context.Context) *mIQueryEngineProviderMockGetQueryEngine {
	if mmGetQueryEngine.mock.funcGetQueryEngine != nil {
		mmGetQueryEngine.mock.t.Fatalf("IQueryEngineProviderMock.GetQueryEngine mock is already set by Set")
	}

	if mmGetQueryEngine.defaultExpectation == nil {
		mmGetQueryEngine.defaultExpectation = &IQueryEngineProviderMockGetQueryEngineExpectation{}
	}

	mmGetQueryEngine.defaultExpectation.params = &IQueryEngineProviderMockGetQueryEngineParams{ctx}
	for _, e := range mmGetQueryEngine.expectations {
		if minimock.Equal(e.params, mmGetQueryEngine.defaultExpectation.params) {
			mmGetQueryEngine.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetQueryEngine.defaultExpectation.params)
		}
	}

	return mmGetQueryEngine
}

// Inspect accepts an inspector function that has same arguments as the IQueryEngineProvider.GetQueryEngine
func (mmGetQueryEngine *mIQueryEngineProviderMockGetQueryEngine) Inspect(f func(ctx context.Context)) *mIQueryEngineProviderMockGetQueryEngine {
	if mmGetQueryEngine.mock.inspectFuncGetQueryEngine != nil {
		mmGetQueryEngine.mock.t.Fatalf("Inspect function is already set for IQueryEngineProviderMock.GetQueryEngine")
	}

	mmGetQueryEngine.mock.inspectFuncGetQueryEngine = f

	return mmGetQueryEngine
}

// Return sets up results that will be returned by IQueryEngineProvider.GetQueryEngine
func (mmGetQueryEngine *mIQueryEngineProviderMockGetQueryEngine) Return(i1 mm_transactor.IQueryEngine) *IQueryEngineProviderMock {
	if mmGetQueryEngine.mock.funcGetQueryEngine != nil {
		mmGetQueryEngine.mock.t.Fatalf("IQueryEngineProviderMock.GetQueryEngine mock is already set by Set")
	}

	if mmGetQueryEngine.defaultExpectation == nil {
		mmGetQueryEngine.defaultExpectation = &IQueryEngineProviderMockGetQueryEngineExpectation{mock: mmGetQueryEngine.mock}
	}
	mmGetQueryEngine.defaultExpectation.results = &IQueryEngineProviderMockGetQueryEngineResults{i1}
	return mmGetQueryEngine.mock
}

// Set uses given function f to mock the IQueryEngineProvider.GetQueryEngine method
func (mmGetQueryEngine *mIQueryEngineProviderMockGetQueryEngine) Set(f func(ctx context.Context) (i1 mm_transactor.IQueryEngine)) *IQueryEngineProviderMock {
	if mmGetQueryEngine.defaultExpectation != nil {
		mmGetQueryEngine.mock.t.Fatalf("Default expectation is already set for the IQueryEngineProvider.GetQueryEngine method")
	}

	if len(mmGetQueryEngine.expectations) > 0 {
		mmGetQueryEngine.mock.t.Fatalf("Some expectations are already set for the IQueryEngineProvider.GetQueryEngine method")
	}

	mmGetQueryEngine.mock.funcGetQueryEngine = f
	return mmGetQueryEngine.mock
}

// When sets expectation for the IQueryEngineProvider.GetQueryEngine which will trigger the result defined by the following
// Then helper
func (mmGetQueryEngine *mIQueryEngineProviderMockGetQueryEngine) When(ctx context.Context) *IQueryEngineProviderMockGetQueryEngineExpectation {
	if mmGetQueryEngine.mock.funcGetQueryEngine != nil {
		mmGetQueryEngine.mock.t.Fatalf("IQueryEngineProviderMock.GetQueryEngine mock is already set by Set")
	}

	expectation := &IQueryEngineProviderMockGetQueryEngineExpectation{
		mock:   mmGetQueryEngine.mock,
		params: &IQueryEngineProviderMockGetQueryEngineParams{ctx},
	}
	mmGetQueryEngine.expectations = append(mmGetQueryEngine.expectations, expectation)
	return expectation
}

// Then sets up IQueryEngineProvider.GetQueryEngine return parameters for the expectation previously defined by the When method
func (e *IQueryEngineProviderMockGetQueryEngineExpectation) Then(i1 mm_transactor.IQueryEngine) *IQueryEngineProviderMock {
	e.results = &IQueryEngineProviderMockGetQueryEngineResults{i1}
	return e.mock
}

// GetQueryEngine implements transactor.IQueryEngineProvider
func (mmGetQueryEngine *IQueryEngineProviderMock) GetQueryEngine(ctx context.Context) (i1 mm_transactor.IQueryEngine) {
	mm_atomic.AddUint64(&mmGetQueryEngine.beforeGetQueryEngineCounter, 1)
	defer mm_atomic.AddUint64(&mmGetQueryEngine.afterGetQueryEngineCounter, 1)

	if mmGetQueryEngine.inspectFuncGetQueryEngine != nil {
		mmGetQueryEngine.inspectFuncGetQueryEngine(ctx)
	}

	mm_params := &IQueryEngineProviderMockGetQueryEngineParams{ctx}

	// Record call args
	mmGetQueryEngine.GetQueryEngineMock.mutex.Lock()
	mmGetQueryEngine.GetQueryEngineMock.callArgs = append(mmGetQueryEngine.GetQueryEngineMock.callArgs, mm_params)
	mmGetQueryEngine.GetQueryEngineMock.mutex.Unlock()

	for _, e := range mmGetQueryEngine.GetQueryEngineMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1
		}
	}

	if mmGetQueryEngine.GetQueryEngineMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetQueryEngine.GetQueryEngineMock.defaultExpectation.Counter, 1)
		mm_want := mmGetQueryEngine.GetQueryEngineMock.defaultExpectation.params
		mm_got := IQueryEngineProviderMockGetQueryEngineParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetQueryEngine.t.Errorf("IQueryEngineProviderMock.GetQueryEngine got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetQueryEngine.GetQueryEngineMock.defaultExpectation.results
		if mm_results == nil {
			mmGetQueryEngine.t.Fatal("No results are set for the IQueryEngineProviderMock.GetQueryEngine")
		}
		return (*mm_results).i1
	}
	if mmGetQueryEngine.funcGetQueryEngine != nil {
		return mmGetQueryEngine.funcGetQueryEngine(ctx)
	}
	mmGetQueryEngine.t.Fatalf("Unexpected call to IQueryEngineProviderMock.GetQueryEngine. %v", ctx)
	return
}

// GetQueryEngineAfterCounter returns a count of finished IQueryEngineProviderMock.GetQueryEngine invocations
func (mmGetQueryEngine *IQueryEngineProviderMock) GetQueryEngineAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetQueryEngine.afterGetQueryEngineCounter)
}

// GetQueryEngineBeforeCounter returns a count of IQueryEngineProviderMock.GetQueryEngine invocations
func (mmGetQueryEngine *IQueryEngineProviderMock) GetQueryEngineBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetQueryEngine.beforeGetQueryEngineCounter)
}

// Calls returns a list of arguments used in each call to IQueryEngineProviderMock.GetQueryEngine.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetQueryEngine *mIQueryEngineProviderMockGetQueryEngine) Calls() []*IQueryEngineProviderMockGetQueryEngineParams {
	mmGetQueryEngine.mutex.RLock()

	argCopy := make([]*IQueryEngineProviderMockGetQueryEngineParams, len(mmGetQueryEngine.callArgs))
	copy(argCopy, mmGetQueryEngine.callArgs)

	mmGetQueryEngine.mutex.RUnlock()

	return argCopy
}

// MinimockGetQueryEngineDone returns true if the count of the GetQueryEngine invocations corresponds
// the number of defined expectations
func (m *IQueryEngineProviderMock) MinimockGetQueryEngineDone() bool {
	for _, e := range m.GetQueryEngineMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetQueryEngineMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetQueryEngineCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetQueryEngine != nil && mm_atomic.LoadUint64(&m.afterGetQueryEngineCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetQueryEngineInspect logs each unmet expectation
func (m *IQueryEngineProviderMock) MinimockGetQueryEngineInspect() {
	for _, e := range m.GetQueryEngineMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to IQueryEngineProviderMock.GetQueryEngine with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetQueryEngineMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetQueryEngineCounter) < 1 {
		if m.GetQueryEngineMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to IQueryEngineProviderMock.GetQueryEngine")
		} else {
			m.t.Errorf("Expected call to IQueryEngineProviderMock.GetQueryEngine with params: %#v", *m.GetQueryEngineMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetQueryEngine != nil && mm_atomic.LoadUint64(&m.afterGetQueryEngineCounter) < 1 {
		m.t.Error("Expected call to IQueryEngineProviderMock.GetQueryEngine")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *IQueryEngineProviderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetQueryEngineInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *IQueryEngineProviderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *IQueryEngineProviderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetQueryEngineDone()
}

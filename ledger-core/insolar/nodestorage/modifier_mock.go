package nodestorage

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/assured-ledger/ledger-core/insolar/node"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
)

// ModifierMock implements Modifier
type ModifierMock struct {
	t minimock.Tester

	funcDeleteForPN          func(pulse pulse.Number)
	inspectFuncDeleteForPN   func(pulse pulse.Number)
	afterDeleteForPNCounter  uint64
	beforeDeleteForPNCounter uint64
	DeleteForPNMock          mModifierMockDeleteForPN

	funcSet          func(pulse pulse.Number, nodes []node.Node) (err error)
	inspectFuncSet   func(pulse pulse.Number, nodes []node.Node)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mModifierMockSet
}

// NewModifierMock returns a mock for Modifier
func NewModifierMock(t minimock.Tester) *ModifierMock {
	m := &ModifierMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteForPNMock = mModifierMockDeleteForPN{mock: m}
	m.DeleteForPNMock.callArgs = []*ModifierMockDeleteForPNParams{}

	m.SetMock = mModifierMockSet{mock: m}
	m.SetMock.callArgs = []*ModifierMockSetParams{}

	return m
}

type mModifierMockDeleteForPN struct {
	mock               *ModifierMock
	defaultExpectation *ModifierMockDeleteForPNExpectation
	expectations       []*ModifierMockDeleteForPNExpectation

	callArgs []*ModifierMockDeleteForPNParams
	mutex    sync.RWMutex
}

// ModifierMockDeleteForPNExpectation specifies expectation struct of the Modifier.DeleteForPN
type ModifierMockDeleteForPNExpectation struct {
	mock   *ModifierMock
	params *ModifierMockDeleteForPNParams

	Counter uint64
}

// ModifierMockDeleteForPNParams contains parameters of the Modifier.DeleteForPN
type ModifierMockDeleteForPNParams struct {
	pulse pulse.Number
}

// Expect sets up expected params for Modifier.DeleteForPN
func (mmDeleteForPN *mModifierMockDeleteForPN) Expect(pulse pulse.Number) *mModifierMockDeleteForPN {
	if mmDeleteForPN.mock.funcDeleteForPN != nil {
		mmDeleteForPN.mock.t.Fatalf("ModifierMock.DeleteForPN mock is already set by Set")
	}

	if mmDeleteForPN.defaultExpectation == nil {
		mmDeleteForPN.defaultExpectation = &ModifierMockDeleteForPNExpectation{}
	}

	mmDeleteForPN.defaultExpectation.params = &ModifierMockDeleteForPNParams{pulse}
	for _, e := range mmDeleteForPN.expectations {
		if minimock.Equal(e.params, mmDeleteForPN.defaultExpectation.params) {
			mmDeleteForPN.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteForPN.defaultExpectation.params)
		}
	}

	return mmDeleteForPN
}

// Inspect accepts an inspector function that has same arguments as the Modifier.DeleteForPN
func (mmDeleteForPN *mModifierMockDeleteForPN) Inspect(f func(pulse pulse.Number)) *mModifierMockDeleteForPN {
	if mmDeleteForPN.mock.inspectFuncDeleteForPN != nil {
		mmDeleteForPN.mock.t.Fatalf("Inspect function is already set for ModifierMock.DeleteForPN")
	}

	mmDeleteForPN.mock.inspectFuncDeleteForPN = f

	return mmDeleteForPN
}

// Return sets up results that will be returned by Modifier.DeleteForPN
func (mmDeleteForPN *mModifierMockDeleteForPN) Return() *ModifierMock {
	if mmDeleteForPN.mock.funcDeleteForPN != nil {
		mmDeleteForPN.mock.t.Fatalf("ModifierMock.DeleteForPN mock is already set by Set")
	}

	if mmDeleteForPN.defaultExpectation == nil {
		mmDeleteForPN.defaultExpectation = &ModifierMockDeleteForPNExpectation{mock: mmDeleteForPN.mock}
	}

	return mmDeleteForPN.mock
}

//Set uses given function f to mock the Modifier.DeleteForPN method
func (mmDeleteForPN *mModifierMockDeleteForPN) Set(f func(pulse pulse.Number)) *ModifierMock {
	if mmDeleteForPN.defaultExpectation != nil {
		mmDeleteForPN.mock.t.Fatalf("Default expectation is already set for the Modifier.DeleteForPN method")
	}

	if len(mmDeleteForPN.expectations) > 0 {
		mmDeleteForPN.mock.t.Fatalf("Some expectations are already set for the Modifier.DeleteForPN method")
	}

	mmDeleteForPN.mock.funcDeleteForPN = f
	return mmDeleteForPN.mock
}

// DeleteForPN implements Modifier
func (mmDeleteForPN *ModifierMock) DeleteForPN(pulse pulse.Number) {
	mm_atomic.AddUint64(&mmDeleteForPN.beforeDeleteForPNCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteForPN.afterDeleteForPNCounter, 1)

	if mmDeleteForPN.inspectFuncDeleteForPN != nil {
		mmDeleteForPN.inspectFuncDeleteForPN(pulse)
	}

	mm_params := &ModifierMockDeleteForPNParams{pulse}

	// Record call args
	mmDeleteForPN.DeleteForPNMock.mutex.Lock()
	mmDeleteForPN.DeleteForPNMock.callArgs = append(mmDeleteForPN.DeleteForPNMock.callArgs, mm_params)
	mmDeleteForPN.DeleteForPNMock.mutex.Unlock()

	for _, e := range mmDeleteForPN.DeleteForPNMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmDeleteForPN.DeleteForPNMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteForPN.DeleteForPNMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteForPN.DeleteForPNMock.defaultExpectation.params
		mm_got := ModifierMockDeleteForPNParams{pulse}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteForPN.t.Errorf("ModifierMock.DeleteForPN got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmDeleteForPN.funcDeleteForPN != nil {
		mmDeleteForPN.funcDeleteForPN(pulse)
		return
	}
	mmDeleteForPN.t.Fatalf("Unexpected call to ModifierMock.DeleteForPN. %v", pulse)

}

// DeleteForPNAfterCounter returns a count of finished ModifierMock.DeleteForPN invocations
func (mmDeleteForPN *ModifierMock) DeleteForPNAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteForPN.afterDeleteForPNCounter)
}

// DeleteForPNBeforeCounter returns a count of ModifierMock.DeleteForPN invocations
func (mmDeleteForPN *ModifierMock) DeleteForPNBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteForPN.beforeDeleteForPNCounter)
}

// Calls returns a list of arguments used in each call to ModifierMock.DeleteForPN.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteForPN *mModifierMockDeleteForPN) Calls() []*ModifierMockDeleteForPNParams {
	mmDeleteForPN.mutex.RLock()

	argCopy := make([]*ModifierMockDeleteForPNParams, len(mmDeleteForPN.callArgs))
	copy(argCopy, mmDeleteForPN.callArgs)

	mmDeleteForPN.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteForPNDone returns true if the count of the DeleteForPN invocations corresponds
// the number of defined expectations
func (m *ModifierMock) MinimockDeleteForPNDone() bool {
	for _, e := range m.DeleteForPNMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteForPNMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteForPNCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteForPN != nil && mm_atomic.LoadUint64(&m.afterDeleteForPNCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteForPNInspect logs each unmet expectation
func (m *ModifierMock) MinimockDeleteForPNInspect() {
	for _, e := range m.DeleteForPNMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ModifierMock.DeleteForPN with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteForPNMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteForPNCounter) < 1 {
		if m.DeleteForPNMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ModifierMock.DeleteForPN")
		} else {
			m.t.Errorf("Expected call to ModifierMock.DeleteForPN with params: %#v", *m.DeleteForPNMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteForPN != nil && mm_atomic.LoadUint64(&m.afterDeleteForPNCounter) < 1 {
		m.t.Error("Expected call to ModifierMock.DeleteForPN")
	}
}

type mModifierMockSet struct {
	mock               *ModifierMock
	defaultExpectation *ModifierMockSetExpectation
	expectations       []*ModifierMockSetExpectation

	callArgs []*ModifierMockSetParams
	mutex    sync.RWMutex
}

// ModifierMockSetExpectation specifies expectation struct of the Modifier.Set
type ModifierMockSetExpectation struct {
	mock    *ModifierMock
	params  *ModifierMockSetParams
	results *ModifierMockSetResults
	Counter uint64
}

// ModifierMockSetParams contains parameters of the Modifier.Set
type ModifierMockSetParams struct {
	pulse pulse.Number
	nodes []node.Node
}

// ModifierMockSetResults contains results of the Modifier.Set
type ModifierMockSetResults struct {
	err error
}

// Expect sets up expected params for Modifier.Set
func (mmSet *mModifierMockSet) Expect(pulse pulse.Number, nodes []node.Node) *mModifierMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("ModifierMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &ModifierMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &ModifierMockSetParams{pulse, nodes}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the Modifier.Set
func (mmSet *mModifierMockSet) Inspect(f func(pulse pulse.Number, nodes []node.Node)) *mModifierMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for ModifierMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by Modifier.Set
func (mmSet *mModifierMockSet) Return(err error) *ModifierMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("ModifierMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &ModifierMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &ModifierMockSetResults{err}
	return mmSet.mock
}

//Set uses given function f to mock the Modifier.Set method
func (mmSet *mModifierMockSet) Set(f func(pulse pulse.Number, nodes []node.Node) (err error)) *ModifierMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the Modifier.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the Modifier.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the Modifier.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mModifierMockSet) When(pulse pulse.Number, nodes []node.Node) *ModifierMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("ModifierMock.Set mock is already set by Set")
	}

	expectation := &ModifierMockSetExpectation{
		mock:   mmSet.mock,
		params: &ModifierMockSetParams{pulse, nodes},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up Modifier.Set return parameters for the expectation previously defined by the When method
func (e *ModifierMockSetExpectation) Then(err error) *ModifierMock {
	e.results = &ModifierMockSetResults{err}
	return e.mock
}

// Set implements Modifier
func (mmSet *ModifierMock) Set(pulse pulse.Number, nodes []node.Node) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(pulse, nodes)
	}

	mm_params := &ModifierMockSetParams{pulse, nodes}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_got := ModifierMockSetParams{pulse, nodes}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("ModifierMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the ModifierMock.Set")
		}
		return (*mm_results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(pulse, nodes)
	}
	mmSet.t.Fatalf("Unexpected call to ModifierMock.Set. %v %v", pulse, nodes)
	return
}

// SetAfterCounter returns a count of finished ModifierMock.Set invocations
func (mmSet *ModifierMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of ModifierMock.Set invocations
func (mmSet *ModifierMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to ModifierMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mModifierMockSet) Calls() []*ModifierMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*ModifierMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *ModifierMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *ModifierMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ModifierMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ModifierMock.Set")
		} else {
			m.t.Errorf("Expected call to ModifierMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to ModifierMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ModifierMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteForPNInspect()

		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ModifierMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ModifierMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteForPNDone() &&
		m.MinimockSetDone()
}

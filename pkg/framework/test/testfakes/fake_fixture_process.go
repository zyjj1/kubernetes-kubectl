// Code generated by counterfeiter. DO NOT EDIT.
package testfakes

import (
	"sync"

	"k8s.io/kubectl/pkg/framework/test"
)

type FakeFixtureProcess struct {
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	URLStub         func() (string, error)
	uRLMutex        sync.RWMutex
	uRLArgsForCall  []struct{}
	uRLReturns      struct {
		result1 string
		result2 error
	}
	uRLReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFixtureProcess) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startReturns.result1
}

func (fake *FakeFixtureProcess) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeFixtureProcess) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFixtureProcess) StartReturnsOnCall(i int, result1 error) {
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFixtureProcess) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub()
	}
}

func (fake *FakeFixtureProcess) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeFixtureProcess) URL() (string, error) {
	fake.uRLMutex.Lock()
	ret, specificReturn := fake.uRLReturnsOnCall[len(fake.uRLArgsForCall)]
	fake.uRLArgsForCall = append(fake.uRLArgsForCall, struct{}{})
	fake.recordInvocation("URL", []interface{}{})
	fake.uRLMutex.Unlock()
	if fake.URLStub != nil {
		return fake.URLStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uRLReturns.result1, fake.uRLReturns.result2
}

func (fake *FakeFixtureProcess) URLCallCount() int {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return len(fake.uRLArgsForCall)
}

func (fake *FakeFixtureProcess) URLReturns(result1 string, result2 error) {
	fake.URLStub = nil
	fake.uRLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFixtureProcess) URLReturnsOnCall(i int, result1 string, result2 error) {
	fake.URLStub = nil
	if fake.uRLReturnsOnCall == nil {
		fake.uRLReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.uRLReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFixtureProcess) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFixtureProcess) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ test.FixtureProcess = new(FakeFixtureProcess)

// Code generated by counterfeiter. DO NOT EDIT.
package compositefakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2action/composite"
)

type FakeServiceActor struct {
	GetServicesWithPlansStub        func(...v2action.Filter) (v2action.ServicesWithPlans, v2action.Warnings, error)
	getServicesWithPlansMutex       sync.RWMutex
	getServicesWithPlansArgsForCall []struct {
		arg1 []v2action.Filter
	}
	getServicesWithPlansReturns struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}
	getServicesWithPlansReturnsOnCall map[int]struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}
	ServiceExistsWithNameStub        func(string) (bool, v2action.Warnings, error)
	serviceExistsWithNameMutex       sync.RWMutex
	serviceExistsWithNameArgsForCall []struct {
		arg1 string
	}
	serviceExistsWithNameReturns struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}
	serviceExistsWithNameReturnsOnCall map[int]struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceActor) GetServicesWithPlans(arg1 ...v2action.Filter) (v2action.ServicesWithPlans, v2action.Warnings, error) {
	fake.getServicesWithPlansMutex.Lock()
	ret, specificReturn := fake.getServicesWithPlansReturnsOnCall[len(fake.getServicesWithPlansArgsForCall)]
	fake.getServicesWithPlansArgsForCall = append(fake.getServicesWithPlansArgsForCall, struct {
		arg1 []v2action.Filter
	}{arg1})
	fake.recordInvocation("GetServicesWithPlans", []interface{}{arg1})
	fake.getServicesWithPlansMutex.Unlock()
	if fake.GetServicesWithPlansStub != nil {
		return fake.GetServicesWithPlansStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServicesWithPlansReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeServiceActor) GetServicesWithPlansCallCount() int {
	fake.getServicesWithPlansMutex.RLock()
	defer fake.getServicesWithPlansMutex.RUnlock()
	return len(fake.getServicesWithPlansArgsForCall)
}

func (fake *FakeServiceActor) GetServicesWithPlansCalls(stub func(...v2action.Filter) (v2action.ServicesWithPlans, v2action.Warnings, error)) {
	fake.getServicesWithPlansMutex.Lock()
	defer fake.getServicesWithPlansMutex.Unlock()
	fake.GetServicesWithPlansStub = stub
}

func (fake *FakeServiceActor) GetServicesWithPlansArgsForCall(i int) []v2action.Filter {
	fake.getServicesWithPlansMutex.RLock()
	defer fake.getServicesWithPlansMutex.RUnlock()
	argsForCall := fake.getServicesWithPlansArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceActor) GetServicesWithPlansReturns(result1 v2action.ServicesWithPlans, result2 v2action.Warnings, result3 error) {
	fake.getServicesWithPlansMutex.Lock()
	defer fake.getServicesWithPlansMutex.Unlock()
	fake.GetServicesWithPlansStub = nil
	fake.getServicesWithPlansReturns = struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceActor) GetServicesWithPlansReturnsOnCall(i int, result1 v2action.ServicesWithPlans, result2 v2action.Warnings, result3 error) {
	fake.getServicesWithPlansMutex.Lock()
	defer fake.getServicesWithPlansMutex.Unlock()
	fake.GetServicesWithPlansStub = nil
	if fake.getServicesWithPlansReturnsOnCall == nil {
		fake.getServicesWithPlansReturnsOnCall = make(map[int]struct {
			result1 v2action.ServicesWithPlans
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServicesWithPlansReturnsOnCall[i] = struct {
		result1 v2action.ServicesWithPlans
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceActor) ServiceExistsWithName(arg1 string) (bool, v2action.Warnings, error) {
	fake.serviceExistsWithNameMutex.Lock()
	ret, specificReturn := fake.serviceExistsWithNameReturnsOnCall[len(fake.serviceExistsWithNameArgsForCall)]
	fake.serviceExistsWithNameArgsForCall = append(fake.serviceExistsWithNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ServiceExistsWithName", []interface{}{arg1})
	fake.serviceExistsWithNameMutex.Unlock()
	if fake.ServiceExistsWithNameStub != nil {
		return fake.ServiceExistsWithNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.serviceExistsWithNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeServiceActor) ServiceExistsWithNameCallCount() int {
	fake.serviceExistsWithNameMutex.RLock()
	defer fake.serviceExistsWithNameMutex.RUnlock()
	return len(fake.serviceExistsWithNameArgsForCall)
}

func (fake *FakeServiceActor) ServiceExistsWithNameCalls(stub func(string) (bool, v2action.Warnings, error)) {
	fake.serviceExistsWithNameMutex.Lock()
	defer fake.serviceExistsWithNameMutex.Unlock()
	fake.ServiceExistsWithNameStub = stub
}

func (fake *FakeServiceActor) ServiceExistsWithNameArgsForCall(i int) string {
	fake.serviceExistsWithNameMutex.RLock()
	defer fake.serviceExistsWithNameMutex.RUnlock()
	argsForCall := fake.serviceExistsWithNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceActor) ServiceExistsWithNameReturns(result1 bool, result2 v2action.Warnings, result3 error) {
	fake.serviceExistsWithNameMutex.Lock()
	defer fake.serviceExistsWithNameMutex.Unlock()
	fake.ServiceExistsWithNameStub = nil
	fake.serviceExistsWithNameReturns = struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceActor) ServiceExistsWithNameReturnsOnCall(i int, result1 bool, result2 v2action.Warnings, result3 error) {
	fake.serviceExistsWithNameMutex.Lock()
	defer fake.serviceExistsWithNameMutex.Unlock()
	fake.ServiceExistsWithNameStub = nil
	if fake.serviceExistsWithNameReturnsOnCall == nil {
		fake.serviceExistsWithNameReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.serviceExistsWithNameReturnsOnCall[i] = struct {
		result1 bool
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServicesWithPlansMutex.RLock()
	defer fake.getServicesWithPlansMutex.RUnlock()
	fake.serviceExistsWithNameMutex.RLock()
	defer fake.serviceExistsWithNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceActor) recordInvocation(key string, args []interface{}) {
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

var _ composite.ServiceActor = new(FakeServiceActor)

// This file was generated by counterfeiter
package applicationfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/commandregistry"
	"github.com/cloudfoundry/cli/cf/commands/application"
	"github.com/cloudfoundry/cli/cf/flags"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeStopper struct {
	MetaDataStub        func() commandregistry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct{}
	metaDataReturns     struct {
		result1 commandregistry.CommandMetadata
	}
	SetDependencyStub        func(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}
	setDependencyReturns struct {
		result1 commandregistry.Command
	}
	RequirementsStub        func(requirementsFactory requirements.Factory, context flags.FlagContext) []requirements.Requirement
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
	}
	ExecuteStub        func(context flags.FlagContext) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		context flags.FlagContext
	}
	executeReturns struct {
		result1 error
	}
	ApplicationStopStub        func(app models.Application, orgName string, spaceName string) (updatedApp models.Application, err error)
	applicationStopMutex       sync.RWMutex
	applicationStopArgsForCall []struct {
		app       models.Application
		orgName   string
		spaceName string
	}
	applicationStopReturns struct {
		result1 models.Application
		result2 error
	}
}

func (fake *FakeStopper) MetaData() commandregistry.CommandMetadata {
	fake.metaDataMutex.Lock()
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	} else {
		return fake.metaDataReturns.result1
	}
}

func (fake *FakeStopper) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeStopper) MetaDataReturns(result1 commandregistry.CommandMetadata) {
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeStopper) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	fake.setDependencyMutex.Lock()
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}{deps, pluginCall})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(deps, pluginCall)
	} else {
		return fake.setDependencyReturns.result1
	}
}

func (fake *FakeStopper) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeStopper) SetDependencyArgsForCall(i int) (commandregistry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return fake.setDependencyArgsForCall[i].deps, fake.setDependencyArgsForCall[i].pluginCall
}

func (fake *FakeStopper) SetDependencyReturns(result1 commandregistry.Command) {
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeStopper) Requirements(requirementsFactory requirements.Factory, context flags.FlagContext) []requirements.Requirement {
	fake.requirementsMutex.Lock()
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}{requirementsFactory, context})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(requirementsFactory, context)
	} else {
		return fake.requirementsReturns.result1
	}
}

func (fake *FakeStopper) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeStopper) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return fake.requirementsArgsForCall[i].requirementsFactory, fake.requirementsArgsForCall[i].context
}

func (fake *FakeStopper) RequirementsReturns(result1 []requirements.Requirement) {
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
	}{result1}
}

func (fake *FakeStopper) Execute(context flags.FlagContext) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		context flags.FlagContext
	}{context})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(context)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeStopper) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeStopper) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].context
}

func (fake *FakeStopper) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStopper) ApplicationStop(app models.Application, orgName string, spaceName string) (updatedApp models.Application, err error) {
	fake.applicationStopMutex.Lock()
	fake.applicationStopArgsForCall = append(fake.applicationStopArgsForCall, struct {
		app       models.Application
		orgName   string
		spaceName string
	}{app, orgName, spaceName})
	fake.applicationStopMutex.Unlock()
	if fake.ApplicationStopStub != nil {
		return fake.ApplicationStopStub(app, orgName, spaceName)
	} else {
		return fake.applicationStopReturns.result1, fake.applicationStopReturns.result2
	}
}

func (fake *FakeStopper) ApplicationStopCallCount() int {
	fake.applicationStopMutex.RLock()
	defer fake.applicationStopMutex.RUnlock()
	return len(fake.applicationStopArgsForCall)
}

func (fake *FakeStopper) ApplicationStopArgsForCall(i int) (models.Application, string, string) {
	fake.applicationStopMutex.RLock()
	defer fake.applicationStopMutex.RUnlock()
	return fake.applicationStopArgsForCall[i].app, fake.applicationStopArgsForCall[i].orgName, fake.applicationStopArgsForCall[i].spaceName
}

func (fake *FakeStopper) ApplicationStopReturns(result1 models.Application, result2 error) {
	fake.ApplicationStopStub = nil
	fake.applicationStopReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

var _ application.Stopper = new(FakeStopper)
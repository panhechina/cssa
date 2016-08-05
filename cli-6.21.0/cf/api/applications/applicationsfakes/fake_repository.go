// This file was generated by counterfeiter
package applicationsfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/applications"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeRepository struct {
	CreateStub        func(params models.AppParams) (createdApp models.Application, apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		params models.AppParams
	}
	createReturns struct {
		result1 models.Application
		result2 error
	}
	GetAppStub        func(appGUID string) (models.Application, error)
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		appGUID string
	}
	getAppReturns struct {
		result1 models.Application
		result2 error
	}
	ReadStub        func(name string) (app models.Application, apiErr error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		name string
	}
	readReturns struct {
		result1 models.Application
		result2 error
	}
	ReadFromSpaceStub        func(name string, spaceGUID string) (app models.Application, apiErr error)
	readFromSpaceMutex       sync.RWMutex
	readFromSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	readFromSpaceReturns struct {
		result1 models.Application
		result2 error
	}
	UpdateStub        func(appGUID string, params models.AppParams) (updatedApp models.Application, apiErr error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		appGUID string
		params  models.AppParams
	}
	updateReturns struct {
		result1 models.Application
		result2 error
	}
	DeleteStub        func(appGUID string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		appGUID string
	}
	deleteReturns struct {
		result1 error
	}
	ReadEnvStub        func(guid string) (*models.Environment, error)
	readEnvMutex       sync.RWMutex
	readEnvArgsForCall []struct {
		guid string
	}
	readEnvReturns struct {
		result1 *models.Environment
		result2 error
	}
	CreateRestageRequestStub        func(guid string) (apiErr error)
	createRestageRequestMutex       sync.RWMutex
	createRestageRequestArgsForCall []struct {
		guid string
	}
	createRestageRequestReturns struct {
		result1 error
	}
}

func (fake *FakeRepository) Create(params models.AppParams) (createdApp models.Application, apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		params models.AppParams
	}{params})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(params)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRepository) CreateArgsForCall(i int) models.AppParams {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].params
}

func (fake *FakeRepository) CreateReturns(result1 models.Application, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetApp(appGUID string) (models.Application, error) {
	fake.getAppMutex.Lock()
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub(appGUID)
	} else {
		return fake.getAppReturns.result1, fake.getAppReturns.result2
	}
}

func (fake *FakeRepository) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeRepository) GetAppArgsForCall(i int) string {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return fake.getAppArgsForCall[i].appGUID
}

func (fake *FakeRepository) GetAppReturns(result1 models.Application, result2 error) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Read(name string) (app models.Application, apiErr error) {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		name string
	}{name})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(name)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeRepository) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeRepository) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].name
}

func (fake *FakeRepository) ReadReturns(result1 models.Application, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) ReadFromSpace(name string, spaceGUID string) (app models.Application, apiErr error) {
	fake.readFromSpaceMutex.Lock()
	fake.readFromSpaceArgsForCall = append(fake.readFromSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.readFromSpaceMutex.Unlock()
	if fake.ReadFromSpaceStub != nil {
		return fake.ReadFromSpaceStub(name, spaceGUID)
	} else {
		return fake.readFromSpaceReturns.result1, fake.readFromSpaceReturns.result2
	}
}

func (fake *FakeRepository) ReadFromSpaceCallCount() int {
	fake.readFromSpaceMutex.RLock()
	defer fake.readFromSpaceMutex.RUnlock()
	return len(fake.readFromSpaceArgsForCall)
}

func (fake *FakeRepository) ReadFromSpaceArgsForCall(i int) (string, string) {
	fake.readFromSpaceMutex.RLock()
	defer fake.readFromSpaceMutex.RUnlock()
	return fake.readFromSpaceArgsForCall[i].name, fake.readFromSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeRepository) ReadFromSpaceReturns(result1 models.Application, result2 error) {
	fake.ReadFromSpaceStub = nil
	fake.readFromSpaceReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Update(appGUID string, params models.AppParams) (updatedApp models.Application, apiErr error) {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		appGUID string
		params  models.AppParams
	}{appGUID, params})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(appGUID, params)
	} else {
		return fake.updateReturns.result1, fake.updateReturns.result2
	}
}

func (fake *FakeRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeRepository) UpdateArgsForCall(i int) (string, models.AppParams) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].appGUID, fake.updateArgsForCall[i].params
}

func (fake *FakeRepository) UpdateReturns(result1 models.Application, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Delete(appGUID string) (apiErr error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(appGUID)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].appGUID
}

func (fake *FakeRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) ReadEnv(guid string) (*models.Environment, error) {
	fake.readEnvMutex.Lock()
	fake.readEnvArgsForCall = append(fake.readEnvArgsForCall, struct {
		guid string
	}{guid})
	fake.readEnvMutex.Unlock()
	if fake.ReadEnvStub != nil {
		return fake.ReadEnvStub(guid)
	} else {
		return fake.readEnvReturns.result1, fake.readEnvReturns.result2
	}
}

func (fake *FakeRepository) ReadEnvCallCount() int {
	fake.readEnvMutex.RLock()
	defer fake.readEnvMutex.RUnlock()
	return len(fake.readEnvArgsForCall)
}

func (fake *FakeRepository) ReadEnvArgsForCall(i int) string {
	fake.readEnvMutex.RLock()
	defer fake.readEnvMutex.RUnlock()
	return fake.readEnvArgsForCall[i].guid
}

func (fake *FakeRepository) ReadEnvReturns(result1 *models.Environment, result2 error) {
	fake.ReadEnvStub = nil
	fake.readEnvReturns = struct {
		result1 *models.Environment
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CreateRestageRequest(guid string) (apiErr error) {
	fake.createRestageRequestMutex.Lock()
	fake.createRestageRequestArgsForCall = append(fake.createRestageRequestArgsForCall, struct {
		guid string
	}{guid})
	fake.createRestageRequestMutex.Unlock()
	if fake.CreateRestageRequestStub != nil {
		return fake.CreateRestageRequestStub(guid)
	} else {
		return fake.createRestageRequestReturns.result1
	}
}

func (fake *FakeRepository) CreateRestageRequestCallCount() int {
	fake.createRestageRequestMutex.RLock()
	defer fake.createRestageRequestMutex.RUnlock()
	return len(fake.createRestageRequestArgsForCall)
}

func (fake *FakeRepository) CreateRestageRequestArgsForCall(i int) string {
	fake.createRestageRequestMutex.RLock()
	defer fake.createRestageRequestMutex.RUnlock()
	return fake.createRestageRequestArgsForCall[i].guid
}

func (fake *FakeRepository) CreateRestageRequestReturns(result1 error) {
	fake.CreateRestageRequestStub = nil
	fake.createRestageRequestReturns = struct {
		result1 error
	}{result1}
}

var _ applications.Repository = new(FakeRepository)

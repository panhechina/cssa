// This file was generated by counterfeiter
package passwordfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/password"
	. "github.com/cloudfoundry/cli/cf/i18n"
)

type FakeRepository struct {
	UpdatePasswordStub        func(old string, new string) error
	updatePasswordMutex       sync.RWMutex
	updatePasswordArgsForCall []struct {
		old string
		new string
	}
	updatePasswordReturns struct {
		result1 error
	}
}

func (fake *FakeRepository) UpdatePassword(old string, new string) error {
	fake.updatePasswordMutex.Lock()
	fake.updatePasswordArgsForCall = append(fake.updatePasswordArgsForCall, struct {
		old string
		new string
	}{old, new})
	fake.updatePasswordMutex.Unlock()
	if fake.UpdatePasswordStub != nil {
		return fake.UpdatePasswordStub(old, new)
	} else {
		return fake.updatePasswordReturns.result1
	}
}

func (fake *FakeRepository) UpdatePasswordCallCount() int {
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	return len(fake.updatePasswordArgsForCall)
}

func (fake *FakeRepository) UpdatePasswordArgsForCall(i int) (string, string) {
	fake.updatePasswordMutex.RLock()
	defer fake.updatePasswordMutex.RUnlock()
	return fake.updatePasswordArgsForCall[i].old, fake.updatePasswordArgsForCall[i].new
}

func (fake *FakeRepository) UpdatePasswordReturns(result1 error) {
	fake.UpdatePasswordStub = nil
	fake.updatePasswordReturns = struct {
		result1 error
	}{result1}
}

var _ password.Repository = new(FakeRepository)

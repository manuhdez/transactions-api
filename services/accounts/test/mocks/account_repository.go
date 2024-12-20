// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"

	account "github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the Repository type
type AccountRepository struct {
	mock.Mock
}

type AccountRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountRepository) EXPECT() *AccountRepository_Expecter {
	return &AccountRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *AccountRepository) Create(_a0 context.Context, _a1 account.Account) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, account.Account) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AccountRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 account.Account
func (_e *AccountRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *AccountRepository_Create_Call {
	return &AccountRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *AccountRepository_Create_Call) Run(run func(_a0 context.Context, _a1 account.Account)) *AccountRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(account.Account))
	})
	return _c
}

func (_c *AccountRepository_Create_Call) Return(_a0 error) *AccountRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_Create_Call) RunAndReturn(run func(context.Context, account.Account) error) *AccountRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *AccountRepository) Delete(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type AccountRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AccountRepository_Expecter) Delete(_a0 interface{}, _a1 interface{}) *AccountRepository_Delete_Call {
	return &AccountRepository_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *AccountRepository_Delete_Call) Run(run func(_a0 context.Context, _a1 string)) *AccountRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountRepository_Delete_Call) Return(_a0 error) *AccountRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_Delete_Call) RunAndReturn(run func(context.Context, string) error) *AccountRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *AccountRepository) Find(_a0 context.Context, _a1 string) (account.Account, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 account.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (account.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) account.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(account.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type AccountRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AccountRepository_Expecter) Find(_a0 interface{}, _a1 interface{}) *AccountRepository_Find_Call {
	return &AccountRepository_Find_Call{Call: _e.mock.On("Find", _a0, _a1)}
}

func (_c *AccountRepository_Find_Call) Run(run func(_a0 context.Context, _a1 string)) *AccountRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountRepository_Find_Call) Return(_a0 account.Account, _a1 error) *AccountRepository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepository_Find_Call) RunAndReturn(run func(context.Context, string) (account.Account, error)) *AccountRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUserId provides a mock function with given fields: _a0, _a1
func (_m *AccountRepository) GetByUserId(_a0 context.Context, _a1 string) ([]account.Account, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserId")
	}

	var r0 []account.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]account.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []account.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]account.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepository_GetByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUserId'
type AccountRepository_GetByUserId_Call struct {
	*mock.Call
}

// GetByUserId is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *AccountRepository_Expecter) GetByUserId(_a0 interface{}, _a1 interface{}) *AccountRepository_GetByUserId_Call {
	return &AccountRepository_GetByUserId_Call{Call: _e.mock.On("GetByUserId", _a0, _a1)}
}

func (_c *AccountRepository_GetByUserId_Call) Run(run func(_a0 context.Context, _a1 string)) *AccountRepository_GetByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountRepository_GetByUserId_Call) Return(_a0 []account.Account, _a1 error) *AccountRepository_GetByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepository_GetByUserId_Call) RunAndReturn(run func(context.Context, string) ([]account.Account, error)) *AccountRepository_GetByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBalance provides a mock function with given fields: _a0, _a1, _a2
func (_m *AccountRepository) UpdateBalance(_a0 context.Context, _a1 string, _a2 float32) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, float32) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_UpdateBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBalance'
type AccountRepository_UpdateBalance_Call struct {
	*mock.Call
}

// UpdateBalance is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 float32
func (_e *AccountRepository_Expecter) UpdateBalance(_a0 interface{}, _a1 interface{}, _a2 interface{}) *AccountRepository_UpdateBalance_Call {
	return &AccountRepository_UpdateBalance_Call{Call: _e.mock.On("UpdateBalance", _a0, _a1, _a2)}
}

func (_c *AccountRepository_UpdateBalance_Call) Run(run func(_a0 context.Context, _a1 string, _a2 float32)) *AccountRepository_UpdateBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(float32))
	})
	return _c
}

func (_c *AccountRepository_UpdateBalance_Call) Return(_a0 error) *AccountRepository_UpdateBalance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_UpdateBalance_Call) RunAndReturn(run func(context.Context, string, float32) error) *AccountRepository_UpdateBalance_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

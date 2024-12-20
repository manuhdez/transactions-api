// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
)

// UserRepository is an autogenerated mock type for the Repository type
type UserRepository struct {
	mock.Mock
}

type UserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepository) EXPECT() *UserRepository_Expecter {
	return &UserRepository_Expecter{mock: &_m.Mock}
}

// All provides a mock function with given fields: ctx
func (_m *UserRepository) All(ctx context.Context) ([]user.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]user.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []user.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type UserRepository_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
//   - ctx context.Context
func (_e *UserRepository_Expecter) All(ctx interface{}) *UserRepository_All_Call {
	return &UserRepository_All_Call{Call: _e.mock.On("All", ctx)}
}

func (_c *UserRepository_All_Call) Run(run func(ctx context.Context)) *UserRepository_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *UserRepository_All_Call) Return(_a0 []user.User, _a1 error) *UserRepository_All_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_All_Call) RunAndReturn(run func(context.Context) ([]user.User, error)) *UserRepository_All_Call {
	_c.Call.Return(run)
	return _c
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_FindByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByEmail'
type UserRepository_FindByEmail_Call struct {
	*mock.Call
}

// FindByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *UserRepository_Expecter) FindByEmail(ctx interface{}, email interface{}) *UserRepository_FindByEmail_Call {
	return &UserRepository_FindByEmail_Call{Call: _e.mock.On("FindByEmail", ctx, email)}
}

func (_c *UserRepository_FindByEmail_Call) Run(run func(ctx context.Context, email string)) *UserRepository_FindByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepository_FindByEmail_Call) Return(_a0 user.User, _a1 error) *UserRepository_FindByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_FindByEmail_Call) RunAndReturn(run func(context.Context, string) (user.User, error)) *UserRepository_FindByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) Save(ctx context.Context, _a1 user.User) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, user.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type UserRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 user.User
func (_e *UserRepository_Expecter) Save(ctx interface{}, _a1 interface{}) *UserRepository_Save_Call {
	return &UserRepository_Save_Call{Call: _e.mock.On("Save", ctx, _a1)}
}

func (_c *UserRepository_Save_Call) Run(run func(ctx context.Context, _a1 user.User)) *UserRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(user.User))
	})
	return _c
}

func (_c *UserRepository_Save_Call) Return(_a0 error) *UserRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_Save_Call) RunAndReturn(run func(context.Context, user.User) error) *UserRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

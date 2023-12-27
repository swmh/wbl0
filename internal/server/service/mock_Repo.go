// Code generated by mockery v2.39.1. DO NOT EDIT.

package service

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRepo is an autogenerated mock type for the Repo type
type MockRepo struct {
	mock.Mock
}

type MockRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepo) EXPECT() *MockRepo_Expecter {
	return &MockRepo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockRepo) Get(ctx context.Context, id string) ([]byte, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockRepo_Expecter) Get(ctx interface{}, id interface{}) *MockRepo_Get_Call {
	return &MockRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockRepo_Get_Call) Run(run func(ctx context.Context, id string)) *MockRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepo_Get_Call) Return(_a0 []byte, _a1 error) *MockRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_Get_Call) RunAndReturn(run func(context.Context, string) ([]byte, error)) *MockRepo_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetNOrders provides a mock function with given fields: ctx, limit
func (_m *MockRepo) GetNOrders(ctx context.Context, limit int) ([]Order, error) {
	ret := _m.Called(ctx, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetNOrders")
	}

	var r0 []Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]Order, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []Order); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetNOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNOrders'
type MockRepo_GetNOrders_Call struct {
	*mock.Call
}

// GetNOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - limit int
func (_e *MockRepo_Expecter) GetNOrders(ctx interface{}, limit interface{}) *MockRepo_GetNOrders_Call {
	return &MockRepo_GetNOrders_Call{Call: _e.mock.On("GetNOrders", ctx, limit)}
}

func (_c *MockRepo_GetNOrders_Call) Run(run func(ctx context.Context, limit int)) *MockRepo_GetNOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRepo_GetNOrders_Call) Return(_a0 []Order, _a1 error) *MockRepo_GetNOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetNOrders_Call) RunAndReturn(run func(context.Context, int) ([]Order, error)) *MockRepo_GetNOrders_Call {
	_c.Call.Return(run)
	return _c
}

// IsNoSuchRow provides a mock function with given fields: err
func (_m *MockRepo) IsNoSuchRow(err error) bool {
	ret := _m.Called(err)

	if len(ret) == 0 {
		panic("no return value specified for IsNoSuchRow")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockRepo_IsNoSuchRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNoSuchRow'
type MockRepo_IsNoSuchRow_Call struct {
	*mock.Call
}

// IsNoSuchRow is a helper method to define mock.On call
//   - err error
func (_e *MockRepo_Expecter) IsNoSuchRow(err interface{}) *MockRepo_IsNoSuchRow_Call {
	return &MockRepo_IsNoSuchRow_Call{Call: _e.mock.On("IsNoSuchRow", err)}
}

func (_c *MockRepo_IsNoSuchRow_Call) Run(run func(err error)) *MockRepo_IsNoSuchRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockRepo_IsNoSuchRow_Call) Return(_a0 bool) *MockRepo_IsNoSuchRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_IsNoSuchRow_Call) RunAndReturn(run func(error) bool) *MockRepo_IsNoSuchRow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepo creates a new instance of MockRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepo {
	mock := &MockRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

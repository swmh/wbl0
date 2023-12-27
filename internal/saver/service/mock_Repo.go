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

// Insert provides a mock function with given fields: ctx, id, value
func (_m *MockRepo) Insert(ctx context.Context, id string, value []byte) error {
	ret := _m.Called(ctx, id, value)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, id, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type MockRepo_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - value []byte
func (_e *MockRepo_Expecter) Insert(ctx interface{}, id interface{}, value interface{}) *MockRepo_Insert_Call {
	return &MockRepo_Insert_Call{Call: _e.mock.On("Insert", ctx, id, value)}
}

func (_c *MockRepo_Insert_Call) Run(run func(ctx context.Context, id string, value []byte)) *MockRepo_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *MockRepo_Insert_Call) Return(_a0 error) *MockRepo_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_Insert_Call) RunAndReturn(run func(context.Context, string, []byte) error) *MockRepo_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// IsAlreadyExists provides a mock function with given fields: err
func (_m *MockRepo) IsAlreadyExists(err error) bool {
	ret := _m.Called(err)

	if len(ret) == 0 {
		panic("no return value specified for IsAlreadyExists")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockRepo_IsAlreadyExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAlreadyExists'
type MockRepo_IsAlreadyExists_Call struct {
	*mock.Call
}

// IsAlreadyExists is a helper method to define mock.On call
//   - err error
func (_e *MockRepo_Expecter) IsAlreadyExists(err interface{}) *MockRepo_IsAlreadyExists_Call {
	return &MockRepo_IsAlreadyExists_Call{Call: _e.mock.On("IsAlreadyExists", err)}
}

func (_c *MockRepo_IsAlreadyExists_Call) Run(run func(err error)) *MockRepo_IsAlreadyExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockRepo_IsAlreadyExists_Call) Return(_a0 bool) *MockRepo_IsAlreadyExists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_IsAlreadyExists_Call) RunAndReturn(run func(error) bool) *MockRepo_IsAlreadyExists_Call {
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

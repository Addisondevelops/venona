// Copyright 2023 The Codefresh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v2.33.1. DO NOT EDIT.

package logger

import (
	log15 "github.com/inconshreveable/log15"
	mock "github.com/stretchr/testify/mock"
)

// MockLogger is an autogenerated mock type for the Logger type
type MockLogger struct {
	mock.Mock
}

type MockLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLogger) EXPECT() *MockLogger_Expecter {
	return &MockLogger_Expecter{mock: &_m.Mock}
}

// Crit provides a mock function with given fields: msg, ctx
func (_m *MockLogger) Crit(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// MockLogger_Crit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Crit'
type MockLogger_Crit_Call struct {
	*mock.Call
}

// Crit is a helper method to define mock.On call
//   - msg string
//   - ctx ...interface{}
func (_e *MockLogger_Expecter) Crit(msg interface{}, ctx ...interface{}) *MockLogger_Crit_Call {
	return &MockLogger_Crit_Call{Call: _e.mock.On("Crit",
		append([]interface{}{msg}, ctx...)...)}
}

func (_c *MockLogger_Crit_Call) Run(run func(msg string, ctx ...interface{})) *MockLogger_Crit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Crit_Call) Return() *MockLogger_Crit_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Crit_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Crit_Call {
	_c.Call.Return(run)
	return _c
}

// Debug provides a mock function with given fields: msg, ctx
func (_m *MockLogger) Debug(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// MockLogger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type MockLogger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//   - msg string
//   - ctx ...interface{}
func (_e *MockLogger_Expecter) Debug(msg interface{}, ctx ...interface{}) *MockLogger_Debug_Call {
	return &MockLogger_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{msg}, ctx...)...)}
}

func (_c *MockLogger_Debug_Call) Run(run func(msg string, ctx ...interface{})) *MockLogger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Debug_Call) Return() *MockLogger_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Debug_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: msg, ctx
func (_m *MockLogger) Error(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// MockLogger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type MockLogger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - msg string
//   - ctx ...interface{}
func (_e *MockLogger_Expecter) Error(msg interface{}, ctx ...interface{}) *MockLogger_Error_Call {
	return &MockLogger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{msg}, ctx...)...)}
}

func (_c *MockLogger_Error_Call) Run(run func(msg string, ctx ...interface{})) *MockLogger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Error_Call) Return() *MockLogger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Error_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Error_Call {
	_c.Call.Return(run)
	return _c
}

// GetHandler provides a mock function with given fields:
func (_m *MockLogger) GetHandler() log15.Handler {
	ret := _m.Called()

	var r0 log15.Handler
	if rf, ok := ret.Get(0).(func() log15.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log15.Handler)
		}
	}

	return r0
}

// MockLogger_GetHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHandler'
type MockLogger_GetHandler_Call struct {
	*mock.Call
}

// GetHandler is a helper method to define mock.On call
func (_e *MockLogger_Expecter) GetHandler() *MockLogger_GetHandler_Call {
	return &MockLogger_GetHandler_Call{Call: _e.mock.On("GetHandler")}
}

func (_c *MockLogger_GetHandler_Call) Run(run func()) *MockLogger_GetHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLogger_GetHandler_Call) Return(_a0 log15.Handler) *MockLogger_GetHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_GetHandler_Call) RunAndReturn(run func() log15.Handler) *MockLogger_GetHandler_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: msg, ctx
func (_m *MockLogger) Info(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// MockLogger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type MockLogger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - msg string
//   - ctx ...interface{}
func (_e *MockLogger_Expecter) Info(msg interface{}, ctx ...interface{}) *MockLogger_Info_Call {
	return &MockLogger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{msg}, ctx...)...)}
}

func (_c *MockLogger_Info_Call) Run(run func(msg string, ctx ...interface{})) *MockLogger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Info_Call) Return() *MockLogger_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Info_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Info_Call {
	_c.Call.Return(run)
	return _c
}

// New provides a mock function with given fields: ctx
func (_m *MockLogger) New(ctx ...interface{}) log15.Logger {
	var _ca []interface{}
	_ca = append(_ca, ctx...)
	ret := _m.Called(_ca...)

	var r0 log15.Logger
	if rf, ok := ret.Get(0).(func(...interface{}) log15.Logger); ok {
		r0 = rf(ctx...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log15.Logger)
		}
	}

	return r0
}

// MockLogger_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type MockLogger_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - ctx ...interface{}
func (_e *MockLogger_Expecter) New(ctx ...interface{}) *MockLogger_New_Call {
	return &MockLogger_New_Call{Call: _e.mock.On("New",
		append([]interface{}{}, ctx...)...)}
}

func (_c *MockLogger_New_Call) Run(run func(ctx ...interface{})) *MockLogger_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_New_Call) Return(_a0 log15.Logger) *MockLogger_New_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_New_Call) RunAndReturn(run func(...interface{}) log15.Logger) *MockLogger_New_Call {
	_c.Call.Return(run)
	return _c
}

// SetHandler provides a mock function with given fields: h
func (_m *MockLogger) SetHandler(h log15.Handler) {
	_m.Called(h)
}

// MockLogger_SetHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHandler'
type MockLogger_SetHandler_Call struct {
	*mock.Call
}

// SetHandler is a helper method to define mock.On call
//   - h log15.Handler
func (_e *MockLogger_Expecter) SetHandler(h interface{}) *MockLogger_SetHandler_Call {
	return &MockLogger_SetHandler_Call{Call: _e.mock.On("SetHandler", h)}
}

func (_c *MockLogger_SetHandler_Call) Run(run func(h log15.Handler)) *MockLogger_SetHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(log15.Handler))
	})
	return _c
}

func (_c *MockLogger_SetHandler_Call) Return() *MockLogger_SetHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_SetHandler_Call) RunAndReturn(run func(log15.Handler)) *MockLogger_SetHandler_Call {
	_c.Call.Return(run)
	return _c
}

// Warn provides a mock function with given fields: msg, ctx
func (_m *MockLogger) Warn(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// MockLogger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type MockLogger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//   - msg string
//   - ctx ...interface{}
func (_e *MockLogger_Expecter) Warn(msg interface{}, ctx ...interface{}) *MockLogger_Warn_Call {
	return &MockLogger_Warn_Call{Call: _e.mock.On("Warn",
		append([]interface{}{msg}, ctx...)...)}
}

func (_c *MockLogger_Warn_Call) Run(run func(msg string, ctx ...interface{})) *MockLogger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Warn_Call) Return() *MockLogger_Warn_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Warn_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Warn_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLogger creates a new instance of MockLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLogger {
	mock := &MockLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
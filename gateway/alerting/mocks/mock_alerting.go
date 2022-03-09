// Code generated by MockGen. DO NOT EDIT.
// Source: odfe-cli/gateway/alerting (interfaces: Gateway)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGateway is a mock of Gateway interface
type MockGateway struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayMockRecorder
}

// MockGatewayMockRecorder is the mock recorder for MockGateway
type MockGatewayMockRecorder struct {
	mock *MockGateway
}

// NewMockGateway creates a new mock instance
func NewMockGateway(ctrl *gomock.Controller) *MockGateway {
	mock := &MockGateway{ctrl: ctrl}
	mock.recorder = &MockGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGateway) EXPECT() *MockGatewayMockRecorder {
	return m.recorder
}

// CreateMonitor mocks base method
func (m *MockGateway) CreateMonitor(arg0 context.Context, arg1 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMonitor", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMonitor indicates an expected call of CreateMonitor
func (mr *MockGatewayMockRecorder) CreateMonitor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMonitor", reflect.TypeOf((*MockGateway)(nil).CreateMonitor), arg0, arg1)
}

// DeleteMonitor mocks base method
func (m *MockGateway) DeleteMonitor(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMonitor", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMonitor indicates an expected call of DeleteMonitor
func (mr *MockGatewayMockRecorder) DeleteMonitor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMonitor", reflect.TypeOf((*MockGateway)(nil).DeleteMonitor), arg0, arg1)
}

// GetMonitor mocks base method
func (m *MockGateway) GetMonitor(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonitor", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMonitor indicates an expected call of GetMonitor
func (mr *MockGatewayMockRecorder) GetMonitor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonitor", reflect.TypeOf((*MockGateway)(nil).GetMonitor), arg0, arg1)
}

// UpdateMonitor mocks base method
func (m *MockGateway) UpdateMonitor(arg0 context.Context, arg1 string, arg2 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMonitor", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMonitor indicates an expected call of UpdateMonitor
func (mr *MockGatewayMockRecorder) UpdateMonitor(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMonitor", reflect.TypeOf((*MockGateway)(nil).UpdateMonitor), arg0, arg1, arg2)
}

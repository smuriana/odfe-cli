// Code generated by MockGen. DO NOT EDIT.
// Source: odfe-cli/controller/alerting (interfaces: Controller)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	alerting "odfe-cli/entity/alerting"
	reflect "reflect"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// CreateMonitors mocks base method
func (m *MockController) CreateMonitors(arg0 context.Context, arg1 alerting.CreateMonitorRequest) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMonitors", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMonitors indicates an expected call of CreateMonitors
func (mr *MockControllerMockRecorder) CreateMonitors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMonitors", reflect.TypeOf((*MockController)(nil).CreateMonitors), arg0, arg1)
}

// DeleteMonitor mocks base method
func (m *MockController) DeleteMonitor(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMonitor", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMonitor indicates an expected call of DeleteMonitor
func (mr *MockControllerMockRecorder) DeleteMonitor(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMonitor", reflect.TypeOf((*MockController)(nil).DeleteMonitor), arg0, arg1, arg2)
}

// GetMonitor mocks base method
func (m *MockController) GetMonitor(arg0 context.Context, arg1 string) (*alerting.MonitorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonitor", arg0, arg1)
	ret0, _ := ret[0].(*alerting.MonitorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMonitor indicates an expected call of GetMonitor
func (mr *MockControllerMockRecorder) GetMonitor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonitor", reflect.TypeOf((*MockController)(nil).GetMonitor), arg0, arg1)
}

// UpdateMonitor mocks base method
func (m *MockController) UpdateMonitor(arg0 context.Context, arg1 alerting.UpdateMonitorUserInput, arg2 bool) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMonitor", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMonitor indicates an expected call of UpdateMonitor
func (mr *MockControllerMockRecorder) UpdateMonitor(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMonitor", reflect.TypeOf((*MockController)(nil).UpdateMonitor), arg0, arg1, arg2)
}

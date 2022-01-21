// Code generated by MockGen. DO NOT EDIT.
// Source: odfe-cli/gateway/ad (interfaces: Gateway)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// CreateDetector mocks base method
func (m *MockGateway) CreateDetector(arg0 context.Context, arg1 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDetector", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDetector indicates an expected call of CreateDetector
func (mr *MockGatewayMockRecorder) CreateDetector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDetector", reflect.TypeOf((*MockGateway)(nil).CreateDetector), arg0, arg1)
}

// DeleteDetector mocks base method
func (m *MockGateway) DeleteDetector(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDetector", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDetector indicates an expected call of DeleteDetector
func (mr *MockGatewayMockRecorder) DeleteDetector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDetector", reflect.TypeOf((*MockGateway)(nil).DeleteDetector), arg0, arg1)
}

// GetDetector mocks base method
func (m *MockGateway) GetDetector(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetector", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetector indicates an expected call of GetDetector
func (mr *MockGatewayMockRecorder) GetDetector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetector", reflect.TypeOf((*MockGateway)(nil).GetDetector), arg0, arg1)
}

// SearchDetector mocks base method
func (m *MockGateway) SearchDetector(arg0 context.Context, arg1 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchDetector", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchDetector indicates an expected call of SearchDetector
func (mr *MockGatewayMockRecorder) SearchDetector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchDetector", reflect.TypeOf((*MockGateway)(nil).SearchDetector), arg0, arg1)
}

// StartDetector mocks base method
func (m *MockGateway) StartDetector(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDetector", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartDetector indicates an expected call of StartDetector
func (mr *MockGatewayMockRecorder) StartDetector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDetector", reflect.TypeOf((*MockGateway)(nil).StartDetector), arg0, arg1)
}

// StopDetector mocks base method
func (m *MockGateway) StopDetector(arg0 context.Context, arg1 string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopDetector", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopDetector indicates an expected call of StopDetector
func (mr *MockGatewayMockRecorder) StopDetector(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopDetector", reflect.TypeOf((*MockGateway)(nil).StopDetector), arg0, arg1)
}

// UpdateDetector mocks base method
func (m *MockGateway) UpdateDetector(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDetector", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDetector indicates an expected call of UpdateDetector
func (mr *MockGatewayMockRecorder) UpdateDetector(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDetector", reflect.TypeOf((*MockGateway)(nil).UpdateDetector), arg0, arg1, arg2)
}

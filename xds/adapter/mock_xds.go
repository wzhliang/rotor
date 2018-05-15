// Code generated by MockGen. DO NOT EDIT.
// Source: xds.go

package adapter

import (
	gomock "github.com/golang/mock/gomock"
	poller "github.com/turbinelabs/rotor/xds/poller"
	reflect "reflect"
)

// MockXDS is a mock of XDS interface
type MockXDS struct {
	ctrl     *gomock.Controller
	recorder *MockXDSMockRecorder
}

// MockXDSMockRecorder is the mock recorder for MockXDS
type MockXDSMockRecorder struct {
	mock *MockXDS
}

// NewMockXDS creates a new mock instance
func NewMockXDS(ctrl *gomock.Controller) *MockXDS {
	mock := &MockXDS{ctrl: ctrl}
	mock.recorder = &MockXDSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockXDS) EXPECT() *MockXDSMockRecorder {
	return m.recorder
}

// Consume mocks base method
func (m *MockXDS) Consume(arg0 *poller.Objects) error {
	ret := m.ctrl.Call(m, "Consume", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Consume indicates an expected call of Consume
func (mr *MockXDSMockRecorder) Consume(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockXDS)(nil).Consume), arg0)
}

// Run mocks base method
func (m *MockXDS) Run() error {
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockXDSMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockXDS)(nil).Run))
}

// Stop mocks base method
func (m *MockXDS) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockXDSMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockXDS)(nil).Stop))
}

// Addr mocks base method
func (m *MockXDS) Addr() string {
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(string)
	return ret0
}

// Addr indicates an expected call of Addr
func (mr *MockXDSMockRecorder) Addr() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockXDS)(nil).Addr))
}
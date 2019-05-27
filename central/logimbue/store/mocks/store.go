// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/logimbue/store (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddLog mocks base method
func (m *MockStore) AddLog(arg0 string) error {
	ret := m.ctrl.Call(m, "AddLog", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLog indicates an expected call of AddLog
func (mr *MockStoreMockRecorder) AddLog(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLog", reflect.TypeOf((*MockStore)(nil).AddLog), arg0)
}

// GetLogs mocks base method
func (m *MockStore) GetLogs() ([]string, error) {
	ret := m.ctrl.Call(m, "GetLogs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockStoreMockRecorder) GetLogs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockStore)(nil).GetLogs))
}

// GetLogsRange mocks base method
func (m *MockStore) GetLogsRange() (int64, int64, error) {
	ret := m.ctrl.Call(m, "GetLogsRange")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLogsRange indicates an expected call of GetLogsRange
func (mr *MockStoreMockRecorder) GetLogsRange() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogsRange", reflect.TypeOf((*MockStore)(nil).GetLogsRange))
}

// RemoveLogs mocks base method
func (m *MockStore) RemoveLogs(arg0, arg1 int64) error {
	ret := m.ctrl.Call(m, "RemoveLogs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLogs indicates an expected call of RemoveLogs
func (mr *MockStoreMockRecorder) RemoveLogs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLogs", reflect.TypeOf((*MockStore)(nil).RemoveLogs), arg0, arg1)
}

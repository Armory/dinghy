// Code generated by MockGen. DO NOT EDIT.
// Source: source.go

// Package source is a generated GoMock package.
package source

import (
	global "github.com/armory/dinghy/pkg/settings/global"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSourceConfiguration is a mock of SourceConfiguration interface
type MockSourceConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockSourceConfigurationMockRecorder
}

// MockSourceConfigurationMockRecorder is the mock recorder for MockSourceConfiguration
type MockSourceConfigurationMockRecorder struct {
	mock *MockSourceConfiguration
}

// NewMockSourceConfiguration creates a new mock instance
func NewMockSourceConfiguration(ctrl *gomock.Controller) *MockSourceConfiguration {
	mock := &MockSourceConfiguration{ctrl: ctrl}
	mock.recorder = &MockSourceConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceConfiguration) EXPECT() *MockSourceConfigurationMockRecorder {
	return m.recorder
}

// GetSourceName mocks base method
func (m *MockSourceConfiguration) GetSourceName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceName indicates an expected call of GetSourceName
func (mr *MockSourceConfigurationMockRecorder) GetSourceName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceName", reflect.TypeOf((*MockSourceConfiguration)(nil).GetSourceName))
}

// LoadSetupSettings mocks base method
func (m *MockSourceConfiguration) LoadSetupSettings() (*global.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSetupSettings")
	ret0, _ := ret[0].(*global.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadSetupSettings indicates an expected call of LoadSetupSettings
func (mr *MockSourceConfigurationMockRecorder) LoadSetupSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSetupSettings", reflect.TypeOf((*MockSourceConfiguration)(nil).LoadSetupSettings))
}

// GetSettings mocks base method
func (m *MockSourceConfiguration) GetSettings(key string) (*global.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", key)
	ret0, _ := ret[0].(*global.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings
func (mr *MockSourceConfigurationMockRecorder) GetSettings(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockSourceConfiguration)(nil).GetSettings), key)
}
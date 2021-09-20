// Code generated by MockGen. DO NOT EDIT.
// Source: execution.go

// Package execution is a generated GoMock package.
package execution

import (
	reflect "reflect"

	setting "github.com/gari8/sub-server/setting"
	gomock "github.com/golang/mock/gomock"
)

// MockFileManager is a mock of FileManager interface.
type MockFileManager struct {
	ctrl     *gomock.Controller
	recorder *MockFileManagerMockRecorder
}

// MockFileManagerMockRecorder is the mock recorder for MockFileManager.
type MockFileManagerMockRecorder struct {
	mock *MockFileManager
}

// NewMockFileManager creates a new mock instance.
func NewMockFileManager(ctrl *gomock.Controller) *MockFileManager {
	mock := &MockFileManager{ctrl: ctrl}
	mock.recorder = &MockFileManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileManager) EXPECT() *MockFileManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFileManager) Create(content []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", content)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockFileManagerMockRecorder) Create(content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFileManager)(nil).Create), content)
}

// Read mocks base method.
func (m *MockFileManager) Read() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockFileManagerMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileManager)(nil).Read))
}

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Serve mocks base method.
func (m *MockServer) Serve(setting setting.Setting) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve", setting)
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockServerMockRecorder) Serve(setting interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockServer)(nil).Serve), setting)
}

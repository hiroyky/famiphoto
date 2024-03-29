// Code generated by MockGen. DO NOT EDIT.
// Source: drivers/es/index.go

// Package mock_es is a generated GoMock package.
package mock_es

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIndexBody is a mock of IndexBody interface.
type MockIndexBody struct {
	ctrl     *gomock.Controller
	recorder *MockIndexBodyMockRecorder
}

// MockIndexBodyMockRecorder is the mock recorder for MockIndexBody.
type MockIndexBodyMockRecorder struct {
	mock *MockIndexBody
}

// NewMockIndexBody creates a new mock instance.
func NewMockIndexBody(ctrl *gomock.Controller) *MockIndexBody {
	mock := &MockIndexBody{ctrl: ctrl}
	mock.recorder = &MockIndexBodyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexBody) EXPECT() *MockIndexBodyMockRecorder {
	return m.recorder
}

// BodyReader mocks base method.
func (m *MockIndexBody) BodyReader() (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BodyReader")
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BodyReader indicates an expected call of BodyReader.
func (mr *MockIndexBodyMockRecorder) BodyReader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyReader", reflect.TypeOf((*MockIndexBody)(nil).BodyReader))
}

// ID mocks base method.
func (m *MockIndexBody) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockIndexBodyMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockIndexBody)(nil).ID))
}

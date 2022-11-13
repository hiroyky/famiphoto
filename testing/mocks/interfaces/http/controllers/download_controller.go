// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/http/controllers/download_controller.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockDownloadController is a mock of DownloadController interface.
type MockDownloadController struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadControllerMockRecorder
}

// MockDownloadControllerMockRecorder is the mock recorder for MockDownloadController.
type MockDownloadControllerMockRecorder struct {
	mock *MockDownloadController
}

// NewMockDownloadController creates a new mock instance.
func NewMockDownloadController(ctrl *gomock.Controller) *MockDownloadController {
	mock := &MockDownloadController{ctrl: ctrl}
	mock.recorder = &MockDownloadControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloadController) EXPECT() *MockDownloadControllerMockRecorder {
	return m.recorder
}

// GetFileDownload mocks base method.
func (m *MockDownloadController) GetFileDownload(ctx echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileDownload", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetFileDownload indicates an expected call of GetFileDownload.
func (mr *MockDownloadControllerMockRecorder) GetFileDownload(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileDownload", reflect.TypeOf((*MockDownloadController)(nil).GetFileDownload), ctx)
}

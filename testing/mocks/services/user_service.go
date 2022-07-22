// Code generated by MockGen. DO NOT EDIT.
// Source: services/user_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AuthUserPassword mocks base method.
func (m *MockUserService) AuthUserPassword(ctx context.Context, userID, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUserPassword", ctx, userID, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// AuthUserPassword indicates an expected call of AuthUserPassword.
func (mr *MockUserServiceMockRecorder) AuthUserPassword(ctx, userID, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUserPassword", reflect.TypeOf((*MockUserService)(nil).AuthUserPassword), ctx, userID, password)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/http/middlewares/auth_middleware.go

// Package mock_middlewares is a generated GoMock package.
package mock_middlewares

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockAuthMiddleware is a mock of AuthMiddleware interface.
type MockAuthMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMiddlewareMockRecorder
}

// MockAuthMiddlewareMockRecorder is the mock recorder for MockAuthMiddleware.
type MockAuthMiddlewareMockRecorder struct {
	mock *MockAuthMiddleware
}

// NewMockAuthMiddleware creates a new mock instance.
func NewMockAuthMiddleware(ctrl *gomock.Controller) *MockAuthMiddleware {
	mock := &MockAuthMiddleware{ctrl: ctrl}
	mock.recorder = &MockAuthMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthMiddleware) EXPECT() *MockAuthMiddlewareMockRecorder {
	return m.recorder
}

// AuthAccessToken mocks base method.
func (m *MockAuthMiddleware) AuthAccessToken() func(http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthAccessToken")
	ret0, _ := ret[0].(func(http.Handler) http.Handler)
	return ret0
}

// AuthAccessToken indicates an expected call of AuthAccessToken.
func (mr *MockAuthMiddlewareMockRecorder) AuthAccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthAccessToken", reflect.TypeOf((*MockAuthMiddleware)(nil).AuthAccessToken))
}

// AuthClientSecret mocks base method.
func (m *MockAuthMiddleware) AuthClientSecret(next echo.HandlerFunc) echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthClientSecret", next)
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// AuthClientSecret indicates an expected call of AuthClientSecret.
func (mr *MockAuthMiddlewareMockRecorder) AuthClientSecret(next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthClientSecret", reflect.TypeOf((*MockAuthMiddleware)(nil).AuthClientSecret), next)
}
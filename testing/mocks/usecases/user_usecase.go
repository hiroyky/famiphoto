// Code generated by MockGen. DO NOT EDIT.
// Source: usecases/user_usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserUseCase) CreateUser(ctx context.Context, userID, name, password string, now time.Time) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, userID, name, password, now)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUseCaseMockRecorder) CreateUser(ctx, userID, name, password, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUseCase)(nil).CreateUser), ctx, userID, name, password, now)
}

// GetUser mocks base method.
func (m *MockUserUseCase) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, userID)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserUseCaseMockRecorder) GetUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserUseCase)(nil).GetUser), ctx, userID)
}

// GetUserPassword mocks base method.
func (m *MockUserUseCase) GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPassword", ctx, userID)
	ret0, _ := ret[0].(*entities.UserPassword)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPassword indicates an expected call of GetUserPassword.
func (mr *MockUserUseCaseMockRecorder) GetUserPassword(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPassword", reflect.TypeOf((*MockUserUseCase)(nil).GetUserPassword), ctx, userID)
}

// GetUsers mocks base method.
func (m *MockUserUseCase) GetUsers(ctx context.Context, userID *string, limit, offset int) (entities.UserList, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, userID, limit, offset)
	ret0, _ := ret[0].(entities.UserList)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserUseCaseMockRecorder) GetUsers(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserUseCase)(nil).GetUsers), ctx, userID, limit, offset)
}

// ValidateToCreateUser mocks base method.
func (m *MockUserUseCase) ValidateToCreateUser(ctx context.Context, userID, name, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToCreateUser", ctx, userID, name, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateToCreateUser indicates an expected call of ValidateToCreateUser.
func (mr *MockUserUseCaseMockRecorder) ValidateToCreateUser(ctx, userID, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToCreateUser", reflect.TypeOf((*MockUserUseCase)(nil).ValidateToCreateUser), ctx, userID, name, password)
}

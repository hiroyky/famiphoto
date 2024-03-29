// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/user_adapter.go

// Package mock_infrastructures is a generated GoMock package.
package mock_infrastructures

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
	filters "github.com/hiroyky/famiphoto/infrastructures/filters"
)

// MockUserAdapter is a mock of UserAdapter interface.
type MockUserAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockUserAdapterMockRecorder
}

// MockUserAdapterMockRecorder is the mock recorder for MockUserAdapter.
type MockUserAdapterMockRecorder struct {
	mock *MockUserAdapter
}

// NewMockUserAdapter creates a new mock instance.
func NewMockUserAdapter(ctrl *gomock.Controller) *MockUserAdapter {
	mock := &MockUserAdapter{ctrl: ctrl}
	mock.recorder = &MockUserAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAdapter) EXPECT() *MockUserAdapterMockRecorder {
	return m.recorder
}

// CountUsers mocks base method.
func (m *MockUserAdapter) CountUsers(ctx context.Context, filter *filters.UserFilter) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUsers", ctx, filter)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUsers indicates an expected call of CountUsers.
func (mr *MockUserAdapterMockRecorder) CountUsers(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUsers", reflect.TypeOf((*MockUserAdapter)(nil).CountUsers), ctx, filter)
}

// CreateUser mocks base method.
func (m *MockUserAdapter) CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user, password, isInitializedPassword, now)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserAdapterMockRecorder) CreateUser(ctx, user, password, isInitializedPassword, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserAdapter)(nil).CreateUser), ctx, user, password, isInitializedPassword, now)
}

// ExistUser mocks base method.
func (m *MockUserAdapter) ExistUser(ctx context.Context, userID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistUser", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistUser indicates an expected call of ExistUser.
func (mr *MockUserAdapterMockRecorder) ExistUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistUser", reflect.TypeOf((*MockUserAdapter)(nil).ExistUser), ctx, userID)
}

// GetUser mocks base method.
func (m *MockUserAdapter) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, userID)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserAdapterMockRecorder) GetUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserAdapter)(nil).GetUser), ctx, userID)
}

// GetUserPassword mocks base method.
func (m *MockUserAdapter) GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPassword", ctx, userID)
	ret0, _ := ret[0].(*entities.UserPassword)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPassword indicates an expected call of GetUserPassword.
func (mr *MockUserAdapterMockRecorder) GetUserPassword(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPassword", reflect.TypeOf((*MockUserAdapter)(nil).GetUserPassword), ctx, userID)
}

// GetUsers mocks base method.
func (m *MockUserAdapter) GetUsers(ctx context.Context, filter *filters.UserFilter, limit, offset int) (entities.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, filter, limit, offset)
	ret0, _ := ret[0].(entities.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserAdapterMockRecorder) GetUsers(ctx, filter, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserAdapter)(nil).GetUsers), ctx, filter, limit, offset)
}

// UpdateUserProfile mocks base method.
func (m *MockUserAdapter) UpdateUserProfile(ctx context.Context, user *entities.User) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProfile", ctx, user)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserProfile indicates an expected call of UpdateUserProfile.
func (mr *MockUserAdapterMockRecorder) UpdateUserProfile(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProfile", reflect.TypeOf((*MockUserAdapter)(nil).UpdateUserProfile), ctx, user)
}

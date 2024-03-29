// Code generated by MockGen. DO NOT EDIT.
// Source: drivers/redis/redis_driver.go

// Package mock_redis is a generated GoMock package.
package mock_redis

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockDriver) Del(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockDriverMockRecorder) Del(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockDriver)(nil).Del), ctx, key)
}

// Get mocks base method.
func (m *MockDriver) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDriverMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDriver)(nil).Get), ctx, key)
}

// GetDel mocks base method.
func (m *MockDriver) GetDel(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDel", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDel indicates an expected call of GetDel.
func (mr *MockDriverMockRecorder) GetDel(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDel", reflect.TypeOf((*MockDriver)(nil).GetDel), ctx, key)
}

// SAdd mocks base method.
func (m *MockDriver) SAdd(ctx context.Context, key string, members ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SAdd", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SAdd indicates an expected call of SAdd.
func (mr *MockDriverMockRecorder) SAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockDriver)(nil).SAdd), varargs...)
}

// SMembers mocks base method.
func (m *MockDriver) SMembers(ctx context.Context, key string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembers", ctx, key)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SMembers indicates an expected call of SMembers.
func (mr *MockDriverMockRecorder) SMembers(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembers", reflect.TypeOf((*MockDriver)(nil).SMembers), ctx, key)
}

// SRem mocks base method.
func (m *MockDriver) SRem(ctx context.Context, key string, members ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SRem", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SRem indicates an expected call of SRem.
func (mr *MockDriverMockRecorder) SRem(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRem", reflect.TypeOf((*MockDriver)(nil).SRem), varargs...)
}

// SetEx mocks base method.
func (m *MockDriver) SetEx(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEx", ctx, key, val, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEx indicates an expected call of SetEx.
func (mr *MockDriverMockRecorder) SetEx(ctx, key, val, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEx", reflect.TypeOf((*MockDriver)(nil).SetEx), ctx, key, val, expiration)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: usecases/group_usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
)

// MockGroupUseCase is a mock of GroupUseCase interface.
type MockGroupUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGroupUseCaseMockRecorder
}

// MockGroupUseCaseMockRecorder is the mock recorder for MockGroupUseCase.
type MockGroupUseCaseMockRecorder struct {
	mock *MockGroupUseCase
}

// NewMockGroupUseCase creates a new mock instance.
func NewMockGroupUseCase(ctrl *gomock.Controller) *MockGroupUseCase {
	mock := &MockGroupUseCase{ctrl: ctrl}
	mock.recorder = &MockGroupUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupUseCase) EXPECT() *MockGroupUseCaseMockRecorder {
	return m.recorder
}

// GetGroup mocks base method.
func (m *MockGroupUseCase) GetGroup(ctx context.Context, groupID string) (*entities.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, groupID)
	ret0, _ := ret[0].(*entities.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockGroupUseCaseMockRecorder) GetGroup(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockGroupUseCase)(nil).GetGroup), ctx, groupID)
}

// GetUserBelongingGroups mocks base method.
func (m *MockGroupUseCase) GetUserBelongingGroups(ctx context.Context, userID string) ([]*entities.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBelongingGroups", ctx, userID)
	ret0, _ := ret[0].([]*entities.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBelongingGroups indicates an expected call of GetUserBelongingGroups.
func (mr *MockGroupUseCaseMockRecorder) GetUserBelongingGroups(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBelongingGroups", reflect.TypeOf((*MockGroupUseCase)(nil).GetUserBelongingGroups), ctx, userID)
}

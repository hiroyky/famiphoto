// Code generated by MockGen. DO NOT EDIT.
// Source: services/photo_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
)

// MockPhotoService is a mock of PhotoService interface.
type MockPhotoService struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoServiceMockRecorder
}

// MockPhotoServiceMockRecorder is the mock recorder for MockPhotoService.
type MockPhotoServiceMockRecorder struct {
	mock *MockPhotoService
}

// NewMockPhotoService creates a new mock instance.
func NewMockPhotoService(ctrl *gomock.Controller) *MockPhotoService {
	mock := &MockPhotoService{ctrl: ctrl}
	mock.recorder = &MockPhotoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoService) EXPECT() *MockPhotoServiceMockRecorder {
	return m.recorder
}

// RegisterPhoto mocks base method.
func (m *MockPhotoService) RegisterPhoto(ctx context.Context, filePath, fileHash, ownerID, groupID string) (*entities.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterPhoto", ctx, filePath, fileHash, ownerID, groupID)
	ret0, _ := ret[0].(*entities.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterPhoto indicates an expected call of RegisterPhoto.
func (mr *MockPhotoServiceMockRecorder) RegisterPhoto(ctx, filePath, fileHash, ownerID, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPhoto", reflect.TypeOf((*MockPhotoService)(nil).RegisterPhoto), ctx, filePath, fileHash, ownerID, groupID)
}
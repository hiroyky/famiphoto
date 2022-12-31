// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/photo_upload_sign_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hiroyky/famiphoto/infrastructures/models"
)

// MockPhotoUploadSignRepository is a mock of PhotoUploadSignRepository interface.
type MockPhotoUploadSignRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoUploadSignRepositoryMockRecorder
}

// MockPhotoUploadSignRepositoryMockRecorder is the mock recorder for MockPhotoUploadSignRepository.
type MockPhotoUploadSignRepositoryMockRecorder struct {
	mock *MockPhotoUploadSignRepository
}

// NewMockPhotoUploadSignRepository creates a new mock instance.
func NewMockPhotoUploadSignRepository(ctrl *gomock.Controller) *MockPhotoUploadSignRepository {
	mock := &MockPhotoUploadSignRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoUploadSignRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoUploadSignRepository) EXPECT() *MockPhotoUploadSignRepositoryMockRecorder {
	return m.recorder
}

// GetSign mocks base method.
func (m *MockPhotoUploadSignRepository) GetSign(ctx context.Context, token string) (*models.PhotoUploadSign, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSign", ctx, token)
	ret0, _ := ret[0].(*models.PhotoUploadSign)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSign indicates an expected call of GetSign.
func (mr *MockPhotoUploadSignRepositoryMockRecorder) GetSign(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSign", reflect.TypeOf((*MockPhotoUploadSignRepository)(nil).GetSign), ctx, token)
}

// SetSignToken mocks base method.
func (m *MockPhotoUploadSignRepository) SetSignToken(ctx context.Context, token, userID, groupID string, expireIn int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSignToken", ctx, token, userID, groupID, expireIn)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSignToken indicates an expected call of SetSignToken.
func (mr *MockPhotoUploadSignRepositoryMockRecorder) SetSignToken(ctx, token, userID, groupID, expireIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSignToken", reflect.TypeOf((*MockPhotoUploadSignRepository)(nil).SetSignToken), ctx, token, userID, groupID, expireIn)
}

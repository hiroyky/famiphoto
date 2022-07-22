// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/photo_thumbnail_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPhotoThumbnailRepository is a mock of PhotoThumbnailRepository interface.
type MockPhotoThumbnailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoThumbnailRepositoryMockRecorder
}

// MockPhotoThumbnailRepositoryMockRecorder is the mock recorder for MockPhotoThumbnailRepository.
type MockPhotoThumbnailRepositoryMockRecorder struct {
	mock *MockPhotoThumbnailRepository
}

// NewMockPhotoThumbnailRepository creates a new mock instance.
func NewMockPhotoThumbnailRepository(ctrl *gomock.Controller) *MockPhotoThumbnailRepository {
	mock := &MockPhotoThumbnailRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoThumbnailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoThumbnailRepository) EXPECT() *MockPhotoThumbnailRepositoryMockRecorder {
	return m.recorder
}

// SavePreview mocks base method.
func (m *MockPhotoThumbnailRepository) SavePreview(ctx context.Context, photoID int, data []byte, groupID, ownerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePreview", ctx, photoID, data, groupID, ownerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePreview indicates an expected call of SavePreview.
func (mr *MockPhotoThumbnailRepositoryMockRecorder) SavePreview(ctx, photoID, data, groupID, ownerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePreview", reflect.TypeOf((*MockPhotoThumbnailRepository)(nil).SavePreview), ctx, photoID, data, groupID, ownerID)
}

// SaveThumbnail mocks base method.
func (m *MockPhotoThumbnailRepository) SaveThumbnail(ctx context.Context, photoID int, data []byte, groupID, ownerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveThumbnail", ctx, photoID, data, groupID, ownerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThumbnail indicates an expected call of SaveThumbnail.
func (mr *MockPhotoThumbnailRepositoryMockRecorder) SaveThumbnail(ctx, photoID, data, groupID, ownerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThumbnail", reflect.TypeOf((*MockPhotoThumbnailRepository)(nil).SaveThumbnail), ctx, photoID, data, groupID, ownerID)
}
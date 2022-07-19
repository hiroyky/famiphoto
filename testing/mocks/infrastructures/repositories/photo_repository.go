// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/photo_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dbmodels "github.com/hiroyky/famiphoto/infrastructures/dbmodels"
)

// MockPhotoRepository is a mock of PhotoRepository interface.
type MockPhotoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoRepositoryMockRecorder
}

// MockPhotoRepositoryMockRecorder is the mock recorder for MockPhotoRepository.
type MockPhotoRepositoryMockRecorder struct {
	mock *MockPhotoRepository
}

// NewMockPhotoRepository creates a new mock instance.
func NewMockPhotoRepository(ctrl *gomock.Controller) *MockPhotoRepository {
	mock := &MockPhotoRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoRepository) EXPECT() *MockPhotoRepositoryMockRecorder {
	return m.recorder
}

// CountPhotos mocks base method.
func (m *MockPhotoRepository) CountPhotos(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountPhotos", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountPhotos indicates an expected call of CountPhotos.
func (mr *MockPhotoRepositoryMockRecorder) CountPhotos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountPhotos", reflect.TypeOf((*MockPhotoRepository)(nil).CountPhotos), ctx)
}

// GetPhotoByFilePath mocks base method.
func (m *MockPhotoRepository) GetPhotoByFilePath(ctx context.Context, fileHash string) (*dbmodels.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoByFilePath", ctx, fileHash)
	ret0, _ := ret[0].(*dbmodels.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoByFilePath indicates an expected call of GetPhotoByFilePath.
func (mr *MockPhotoRepositoryMockRecorder) GetPhotoByFilePath(ctx, fileHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoByFilePath", reflect.TypeOf((*MockPhotoRepository)(nil).GetPhotoByFilePath), ctx, fileHash)
}

// GetPhotos mocks base method.
func (m *MockPhotoRepository) GetPhotos(ctx context.Context, limit, offset int64) ([]*dbmodels.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotos", ctx, limit, offset)
	ret0, _ := ret[0].([]*dbmodels.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotos indicates an expected call of GetPhotos.
func (mr *MockPhotoRepositoryMockRecorder) GetPhotos(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotos", reflect.TypeOf((*MockPhotoRepository)(nil).GetPhotos), ctx, limit, offset)
}

// InsertPhoto mocks base method.
func (m *MockPhotoRepository) InsertPhoto(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPhoto", ctx, photo)
	ret0, _ := ret[0].(*dbmodels.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPhoto indicates an expected call of InsertPhoto.
func (mr *MockPhotoRepositoryMockRecorder) InsertPhoto(ctx, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPhoto", reflect.TypeOf((*MockPhotoRepository)(nil).InsertPhoto), ctx, photo)
}

// UpdatePhoto mocks base method.
func (m *MockPhotoRepository) UpdatePhoto(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", ctx, photo)
	ret0, _ := ret[0].(*dbmodels.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockPhotoRepositoryMockRecorder) UpdatePhoto(ctx, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).UpdatePhoto), ctx, photo)
}

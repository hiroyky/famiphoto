// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/exif_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dbmodels "github.com/hiroyky/famiphoto/infrastructures/dbmodels"
)

// MockExifRepository is a mock of ExifRepository interface.
type MockExifRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExifRepositoryMockRecorder
}

// MockExifRepositoryMockRecorder is the mock recorder for MockExifRepository.
type MockExifRepositoryMockRecorder struct {
	mock *MockExifRepository
}

// NewMockExifRepository creates a new mock instance.
func NewMockExifRepository(ctrl *gomock.Controller) *MockExifRepository {
	mock := &MockExifRepository{ctrl: ctrl}
	mock.recorder = &MockExifRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExifRepository) EXPECT() *MockExifRepositoryMockRecorder {
	return m.recorder
}

// GetPhotoMetaDataByPhotoID mocks base method.
func (m *MockExifRepository) GetPhotoMetaDataByPhotoID(ctx context.Context, photoID int) (dbmodels.ExifSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoMetaDataByPhotoID", ctx, photoID)
	ret0, _ := ret[0].(dbmodels.ExifSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoMetaDataByPhotoID indicates an expected call of GetPhotoMetaDataByPhotoID.
func (mr *MockExifRepositoryMockRecorder) GetPhotoMetaDataByPhotoID(ctx, photoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoMetaDataByPhotoID", reflect.TypeOf((*MockExifRepository)(nil).GetPhotoMetaDataByPhotoID), ctx, photoID)
}

// GetPhotoMetaItemByTagID mocks base method.
func (m *MockExifRepository) GetPhotoMetaItemByTagID(ctx context.Context, photoID, tagID int) (*dbmodels.Exif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoMetaItemByTagID", ctx, photoID, tagID)
	ret0, _ := ret[0].(*dbmodels.Exif)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoMetaItemByTagID indicates an expected call of GetPhotoMetaItemByTagID.
func (mr *MockExifRepositoryMockRecorder) GetPhotoMetaItemByTagID(ctx, photoID, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoMetaItemByTagID", reflect.TypeOf((*MockExifRepository)(nil).GetPhotoMetaItemByTagID), ctx, photoID, tagID)
}

// GetPhotoMetaItemsByPhotoIDsTagID mocks base method.
func (m *MockExifRepository) GetPhotoMetaItemsByPhotoIDsTagID(ctx context.Context, photoIDs []int, tagID int) (dbmodels.ExifSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoMetaItemsByPhotoIDsTagID", ctx, photoIDs, tagID)
	ret0, _ := ret[0].(dbmodels.ExifSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoMetaItemsByPhotoIDsTagID indicates an expected call of GetPhotoMetaItemsByPhotoIDsTagID.
func (mr *MockExifRepositoryMockRecorder) GetPhotoMetaItemsByPhotoIDsTagID(ctx, photoIDs, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoMetaItemsByPhotoIDsTagID", reflect.TypeOf((*MockExifRepository)(nil).GetPhotoMetaItemsByPhotoIDsTagID), ctx, photoIDs, tagID)
}

// InsertPhotoMetaItem mocks base method.
func (m *MockExifRepository) InsertPhotoMetaItem(ctx context.Context, exif *dbmodels.Exif) (*dbmodels.Exif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPhotoMetaItem", ctx, exif)
	ret0, _ := ret[0].(*dbmodels.Exif)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPhotoMetaItem indicates an expected call of InsertPhotoMetaItem.
func (mr *MockExifRepositoryMockRecorder) InsertPhotoMetaItem(ctx, exif interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPhotoMetaItem", reflect.TypeOf((*MockExifRepository)(nil).InsertPhotoMetaItem), ctx, exif)
}

// UpdatePhotoMetaItem mocks base method.
func (m *MockExifRepository) UpdatePhotoMetaItem(ctx context.Context, exif *dbmodels.Exif) (*dbmodels.Exif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhotoMetaItem", ctx, exif)
	ret0, _ := ret[0].(*dbmodels.Exif)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePhotoMetaItem indicates an expected call of UpdatePhotoMetaItem.
func (mr *MockExifRepositoryMockRecorder) UpdatePhotoMetaItem(ctx, exif interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhotoMetaItem", reflect.TypeOf((*MockExifRepository)(nil).UpdatePhotoMetaItem), ctx, exif)
}

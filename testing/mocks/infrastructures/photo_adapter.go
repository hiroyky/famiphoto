// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/photo_adapter.go

// Package mock_infrastructures is a generated GoMock package.
package mock_infrastructures

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
)

// MockPhotoAdapter is a mock of PhotoAdapter interface.
type MockPhotoAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoAdapterMockRecorder
}

// MockPhotoAdapterMockRecorder is the mock recorder for MockPhotoAdapter.
type MockPhotoAdapterMockRecorder struct {
	mock *MockPhotoAdapter
}

// NewMockPhotoAdapter creates a new mock instance.
func NewMockPhotoAdapter(ctrl *gomock.Controller) *MockPhotoAdapter {
	mock := &MockPhotoAdapter{ctrl: ctrl}
	mock.recorder = &MockPhotoAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoAdapter) EXPECT() *MockPhotoAdapterMockRecorder {
	return m.recorder
}

// CountPhotos mocks base method.
func (m *MockPhotoAdapter) CountPhotos(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountPhotos", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountPhotos indicates an expected call of CountPhotos.
func (mr *MockPhotoAdapterMockRecorder) CountPhotos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountPhotos", reflect.TypeOf((*MockPhotoAdapter)(nil).CountPhotos), ctx)
}

// ExistPhotoFileByFilePath mocks base method.
func (m *MockPhotoAdapter) ExistPhotoFileByFilePath(ctx context.Context, filePath string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistPhotoFileByFilePath", ctx, filePath)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistPhotoFileByFilePath indicates an expected call of ExistPhotoFileByFilePath.
func (mr *MockPhotoAdapterMockRecorder) ExistPhotoFileByFilePath(ctx, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistPhotoFileByFilePath", reflect.TypeOf((*MockPhotoAdapter)(nil).ExistPhotoFileByFilePath), ctx, filePath)
}

// GetPhotoByPhotoID mocks base method.
func (m *MockPhotoAdapter) GetPhotoByPhotoID(ctx context.Context, photoID int) (*entities.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoByPhotoID", ctx, photoID)
	ret0, _ := ret[0].(*entities.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoByPhotoID indicates an expected call of GetPhotoByPhotoID.
func (mr *MockPhotoAdapterMockRecorder) GetPhotoByPhotoID(ctx, photoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoByPhotoID", reflect.TypeOf((*MockPhotoAdapter)(nil).GetPhotoByPhotoID), ctx, photoID)
}

// GetPhotoFileByFilePath mocks base method.
func (m *MockPhotoAdapter) GetPhotoFileByFilePath(ctx context.Context, filePath string) (*entities.PhotoFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoFileByFilePath", ctx, filePath)
	ret0, _ := ret[0].(*entities.PhotoFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoFileByFilePath indicates an expected call of GetPhotoFileByFilePath.
func (mr *MockPhotoAdapterMockRecorder) GetPhotoFileByFilePath(ctx, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoFileByFilePath", reflect.TypeOf((*MockPhotoAdapter)(nil).GetPhotoFileByFilePath), ctx, filePath)
}

// GetPhotoFileByPhotoFileID mocks base method.
func (m *MockPhotoAdapter) GetPhotoFileByPhotoFileID(ctx context.Context, photoFileID int) (*entities.PhotoFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoFileByPhotoFileID", ctx, photoFileID)
	ret0, _ := ret[0].(*entities.PhotoFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoFileByPhotoFileID indicates an expected call of GetPhotoFileByPhotoFileID.
func (mr *MockPhotoAdapterMockRecorder) GetPhotoFileByPhotoFileID(ctx, photoFileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoFileByPhotoFileID", reflect.TypeOf((*MockPhotoAdapter)(nil).GetPhotoFileByPhotoFileID), ctx, photoFileID)
}

// GetPhotoFilesByPhotoID mocks base method.
func (m *MockPhotoAdapter) GetPhotoFilesByPhotoID(ctx context.Context, photoID int) ([]*entities.PhotoFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoFilesByPhotoID", ctx, photoID)
	ret0, _ := ret[0].([]*entities.PhotoFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoFilesByPhotoID indicates an expected call of GetPhotoFilesByPhotoID.
func (mr *MockPhotoAdapterMockRecorder) GetPhotoFilesByPhotoID(ctx, photoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoFilesByPhotoID", reflect.TypeOf((*MockPhotoAdapter)(nil).GetPhotoFilesByPhotoID), ctx, photoID)
}

// GetPhotoMetaByPhotoID mocks base method.
func (m *MockPhotoAdapter) GetPhotoMetaByPhotoID(ctx context.Context, photoID int) (entities.PhotoMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoMetaByPhotoID", ctx, photoID)
	ret0, _ := ret[0].(entities.PhotoMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoMetaByPhotoID indicates an expected call of GetPhotoMetaByPhotoID.
func (mr *MockPhotoAdapterMockRecorder) GetPhotoMetaByPhotoID(ctx, photoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoMetaByPhotoID", reflect.TypeOf((*MockPhotoAdapter)(nil).GetPhotoMetaByPhotoID), ctx, photoID)
}

// GetPhotoMetaItemByPhotoIDTagID mocks base method.
func (m *MockPhotoAdapter) GetPhotoMetaItemByPhotoIDTagID(ctx context.Context, photoID, tagID int) (*entities.PhotoMetaItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotoMetaItemByPhotoIDTagID", ctx, photoID, tagID)
	ret0, _ := ret[0].(*entities.PhotoMetaItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotoMetaItemByPhotoIDTagID indicates an expected call of GetPhotoMetaItemByPhotoIDTagID.
func (mr *MockPhotoAdapterMockRecorder) GetPhotoMetaItemByPhotoIDTagID(ctx, photoID, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotoMetaItemByPhotoIDTagID", reflect.TypeOf((*MockPhotoAdapter)(nil).GetPhotoMetaItemByPhotoIDTagID), ctx, photoID, tagID)
}

// GetPhotos mocks base method.
func (m *MockPhotoAdapter) GetPhotos(ctx context.Context, limit, offset int) (entities.PhotoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotos", ctx, limit, offset)
	ret0, _ := ret[0].(entities.PhotoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotos indicates an expected call of GetPhotos.
func (mr *MockPhotoAdapterMockRecorder) GetPhotos(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotos", reflect.TypeOf((*MockPhotoAdapter)(nil).GetPhotos), ctx, limit, offset)
}

// UpsertPhotoByFilePath mocks base method.
func (m *MockPhotoAdapter) UpsertPhotoByFilePath(ctx context.Context, photo *entities.Photo) (*entities.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPhotoByFilePath", ctx, photo)
	ret0, _ := ret[0].(*entities.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertPhotoByFilePath indicates an expected call of UpsertPhotoByFilePath.
func (mr *MockPhotoAdapterMockRecorder) UpsertPhotoByFilePath(ctx, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPhotoByFilePath", reflect.TypeOf((*MockPhotoAdapter)(nil).UpsertPhotoByFilePath), ctx, photo)
}

// UpsertPhotoMetaItemByPhotoTagID mocks base method.
func (m *MockPhotoAdapter) UpsertPhotoMetaItemByPhotoTagID(ctx context.Context, photoID int, metaItem *entities.PhotoMetaItem) (*entities.PhotoMetaItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPhotoMetaItemByPhotoTagID", ctx, photoID, metaItem)
	ret0, _ := ret[0].(*entities.PhotoMetaItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertPhotoMetaItemByPhotoTagID indicates an expected call of UpsertPhotoMetaItemByPhotoTagID.
func (mr *MockPhotoAdapterMockRecorder) UpsertPhotoMetaItemByPhotoTagID(ctx, photoID, metaItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPhotoMetaItemByPhotoTagID", reflect.TypeOf((*MockPhotoAdapter)(nil).UpsertPhotoMetaItemByPhotoTagID), ctx, photoID, metaItem)
}

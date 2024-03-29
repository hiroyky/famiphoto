// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/photo_storage_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	os "os"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hiroyky/famiphoto/infrastructures/models"
)

// MockPhotoStorageRepository is a mock of PhotoStorageRepository interface.
type MockPhotoStorageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoStorageRepositoryMockRecorder
}

// MockPhotoStorageRepositoryMockRecorder is the mock recorder for MockPhotoStorageRepository.
type MockPhotoStorageRepositoryMockRecorder struct {
	mock *MockPhotoStorageRepository
}

// NewMockPhotoStorageRepository creates a new mock instance.
func NewMockPhotoStorageRepository(ctrl *gomock.Controller) *MockPhotoStorageRepository {
	mock := &MockPhotoStorageRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoStorageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoStorageRepository) EXPECT() *MockPhotoStorageRepositoryMockRecorder {
	return m.recorder
}

// CreateUserDir mocks base method.
func (m *MockPhotoStorageRepository) CreateUserDir(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserDir", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserDir indicates an expected call of CreateUserDir.
func (mr *MockPhotoStorageRepositoryMockRecorder) CreateUserDir(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserDir", reflect.TypeOf((*MockPhotoStorageRepository)(nil).CreateUserDir), userID)
}

// LoadContent mocks base method.
func (m *MockPhotoStorageRepository) LoadContent(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadContent", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadContent indicates an expected call of LoadContent.
func (mr *MockPhotoStorageRepositoryMockRecorder) LoadContent(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadContent", reflect.TypeOf((*MockPhotoStorageRepository)(nil).LoadContent), path)
}

// ParsePhotoMetaFromFile mocks base method.
func (m *MockPhotoStorageRepository) ParsePhotoMetaFromFile(path string) ([]models.IfdEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParsePhotoMetaFromFile", path)
	ret0, _ := ret[0].([]models.IfdEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParsePhotoMetaFromFile indicates an expected call of ParsePhotoMetaFromFile.
func (mr *MockPhotoStorageRepositoryMockRecorder) ParsePhotoMetaFromFile(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParsePhotoMetaFromFile", reflect.TypeOf((*MockPhotoStorageRepository)(nil).ParsePhotoMetaFromFile), path)
}

// ReadDir mocks base method.
func (m *MockPhotoStorageRepository) ReadDir(dirPath string) ([]os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", dirPath)
	ret0, _ := ret[0].([]os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockPhotoStorageRepositoryMockRecorder) ReadDir(dirPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockPhotoStorageRepository)(nil).ReadDir), dirPath)
}

// SaveContent mocks base method.
func (m *MockPhotoStorageRepository) SaveContent(userID, fileName string, dateTimeOriginal time.Time, data []byte) (os.FileInfo, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveContent", userID, fileName, dateTimeOriginal, data)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SaveContent indicates an expected call of SaveContent.
func (mr *MockPhotoStorageRepositoryMockRecorder) SaveContent(userID, fileName, dateTimeOriginal, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveContent", reflect.TypeOf((*MockPhotoStorageRepository)(nil).SaveContent), userID, fileName, dateTimeOriginal, data)
}

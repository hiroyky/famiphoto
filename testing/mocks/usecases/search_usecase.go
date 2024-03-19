// Code generated by MockGen. DO NOT EDIT.
// Source: usecases/search_usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
)

// MockSearchUseCase is a mock of SearchUseCase interface.
type MockSearchUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchUseCaseMockRecorder
}

// MockSearchUseCaseMockRecorder is the mock recorder for MockSearchUseCase.
type MockSearchUseCaseMockRecorder struct {
	mock *MockSearchUseCase
}

// NewMockSearchUseCase creates a new mock instance.
func NewMockSearchUseCase(ctrl *gomock.Controller) *MockSearchUseCase {
	mock := &MockSearchUseCase{ctrl: ctrl}
	mock.recorder = &MockSearchUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchUseCase) EXPECT() *MockSearchUseCaseMockRecorder {
	return m.recorder
}

// AggregateDateTimeOriginal mocks base method.
func (m *MockSearchUseCase) AggregateDateTimeOriginal(ctx context.Context, year, month int) (entities.PhotoDateTimeAggregation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateDateTimeOriginal", ctx, year, month)
	ret0, _ := ret[0].(entities.PhotoDateTimeAggregation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AggregateDateTimeOriginal indicates an expected call of AggregateDateTimeOriginal.
func (mr *MockSearchUseCaseMockRecorder) AggregateDateTimeOriginal(ctx, year, month interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateDateTimeOriginal", reflect.TypeOf((*MockSearchUseCase)(nil).AggregateDateTimeOriginal), ctx, year, month)
}

// AppendAllPhotoDocuments mocks base method.
func (m *MockSearchUseCase) AppendAllPhotoDocuments(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendAllPhotoDocuments", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendAllPhotoDocuments indicates an expected call of AppendAllPhotoDocuments.
func (mr *MockSearchUseCaseMockRecorder) AppendAllPhotoDocuments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendAllPhotoDocuments", reflect.TypeOf((*MockSearchUseCase)(nil).AppendAllPhotoDocuments), ctx)
}

// SearchPhotoByPhotoID mocks base method.
func (m *MockSearchUseCase) SearchPhotoByPhotoID(ctx context.Context, id int) (*entities.PhotoSearchResultItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPhotoByPhotoID", ctx, id)
	ret0, _ := ret[0].(*entities.PhotoSearchResultItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPhotoByPhotoID indicates an expected call of SearchPhotoByPhotoID.
func (mr *MockSearchUseCaseMockRecorder) SearchPhotoByPhotoID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPhotoByPhotoID", reflect.TypeOf((*MockSearchUseCase)(nil).SearchPhotoByPhotoID), ctx, id)
}

// SearchPhotos mocks base method.
func (m *MockSearchUseCase) SearchPhotos(ctx context.Context, id *int, limit, offset int) (*entities.PhotoSearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPhotos", ctx, id, limit, offset)
	ret0, _ := ret[0].(*entities.PhotoSearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPhotos indicates an expected call of SearchPhotos.
func (mr *MockSearchUseCaseMockRecorder) SearchPhotos(ctx, id, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPhotos", reflect.TypeOf((*MockSearchUseCase)(nil).SearchPhotos), ctx, id, limit, offset)
}

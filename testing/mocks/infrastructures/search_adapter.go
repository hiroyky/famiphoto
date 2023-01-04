// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/search_adapter.go

// Package mock_infrastructures is a generated GoMock package.
package mock_infrastructures

import (
	context "context"
	reflect "reflect"

	esutil "github.com/elastic/go-elasticsearch/v8/esutil"
	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
	filters "github.com/hiroyky/famiphoto/infrastructures/filters"
)

// MockSearchAdapter is a mock of SearchAdapter interface.
type MockSearchAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockSearchAdapterMockRecorder
}

// MockSearchAdapterMockRecorder is the mock recorder for MockSearchAdapter.
type MockSearchAdapterMockRecorder struct {
	mock *MockSearchAdapter
}

// NewMockSearchAdapter creates a new mock instance.
func NewMockSearchAdapter(ctrl *gomock.Controller) *MockSearchAdapter {
	mock := &MockSearchAdapter{ctrl: ctrl}
	mock.recorder = &MockSearchAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchAdapter) EXPECT() *MockSearchAdapterMockRecorder {
	return m.recorder
}

// BulkInsertPhotos mocks base method.
func (m *MockSearchAdapter) BulkInsertPhotos(ctx context.Context, photos entities.PhotoList) (*esutil.BulkIndexerStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkInsertPhotos", ctx, photos)
	ret0, _ := ret[0].(*esutil.BulkIndexerStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkInsertPhotos indicates an expected call of BulkInsertPhotos.
func (mr *MockSearchAdapterMockRecorder) BulkInsertPhotos(ctx, photos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkInsertPhotos", reflect.TypeOf((*MockSearchAdapter)(nil).BulkInsertPhotos), ctx, photos)
}

// InsertPhoto mocks base method.
func (m *MockSearchAdapter) InsertPhoto(ctx context.Context, photo *entities.Photo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPhoto", ctx, photo)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPhoto indicates an expected call of InsertPhoto.
func (mr *MockSearchAdapterMockRecorder) InsertPhoto(ctx, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPhoto", reflect.TypeOf((*MockSearchAdapter)(nil).InsertPhoto), ctx, photo)
}

// SearchPhotos mocks base method.
func (m *MockSearchAdapter) SearchPhotos(ctx context.Context, q *filters.PhotoSearchQuery) (*entities.PhotoSearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPhotos", ctx, q)
	ret0, _ := ret[0].(*entities.PhotoSearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPhotos indicates an expected call of SearchPhotos.
func (mr *MockSearchAdapterMockRecorder) SearchPhotos(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPhotos", reflect.TypeOf((*MockSearchAdapter)(nil).SearchPhotos), ctx, q)
}

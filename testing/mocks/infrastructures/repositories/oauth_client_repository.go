// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/oauth_client_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dbmodels "github.com/hiroyky/famiphoto/infrastructures/dbmodels"
)

// MockOAuthClientRepository is a mock of OAuthClientRepository interface.
type MockOAuthClientRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthClientRepositoryMockRecorder
}

// MockOAuthClientRepositoryMockRecorder is the mock recorder for MockOAuthClientRepository.
type MockOAuthClientRepositoryMockRecorder struct {
	mock *MockOAuthClientRepository
}

// NewMockOAuthClientRepository creates a new mock instance.
func NewMockOAuthClientRepository(ctrl *gomock.Controller) *MockOAuthClientRepository {
	mock := &MockOAuthClientRepository{ctrl: ctrl}
	mock.recorder = &MockOAuthClientRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthClientRepository) EXPECT() *MockOAuthClientRepositoryMockRecorder {
	return m.recorder
}

// CreateOAuthClient mocks base method.
func (m *MockOAuthClientRepository) CreateOAuthClient(ctx context.Context, client *dbmodels.OauthClient, redirectURLs []*dbmodels.OauthClientRedirectURL) (*dbmodels.OauthClient, []*dbmodels.OauthClientRedirectURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOAuthClient", ctx, client, redirectURLs)
	ret0, _ := ret[0].(*dbmodels.OauthClient)
	ret1, _ := ret[1].([]*dbmodels.OauthClientRedirectURL)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateOAuthClient indicates an expected call of CreateOAuthClient.
func (mr *MockOAuthClientRepositoryMockRecorder) CreateOAuthClient(ctx, client, redirectURLs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOAuthClient", reflect.TypeOf((*MockOAuthClientRepository)(nil).CreateOAuthClient), ctx, client, redirectURLs)
}

// ExistOauthClient mocks base method.
func (m *MockOAuthClientRepository) ExistOauthClient(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistOauthClient", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistOauthClient indicates an expected call of ExistOauthClient.
func (mr *MockOAuthClientRepositoryMockRecorder) ExistOauthClient(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistOauthClient", reflect.TypeOf((*MockOAuthClientRepository)(nil).ExistOauthClient), ctx, id)
}

// GetByOauthClientID mocks base method.
func (m *MockOAuthClientRepository) GetByOauthClientID(ctx context.Context, id string) (*dbmodels.OauthClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByOauthClientID", ctx, id)
	ret0, _ := ret[0].(*dbmodels.OauthClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByOauthClientID indicates an expected call of GetByOauthClientID.
func (mr *MockOAuthClientRepositoryMockRecorder) GetByOauthClientID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByOauthClientID", reflect.TypeOf((*MockOAuthClientRepository)(nil).GetByOauthClientID), ctx, id)
}
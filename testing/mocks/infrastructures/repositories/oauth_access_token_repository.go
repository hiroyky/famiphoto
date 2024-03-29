// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructures/repositories/oauth_access_token_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
	models "github.com/hiroyky/famiphoto/infrastructures/models"
)

// MockOauthAccessTokenRepository is a mock of OauthAccessTokenRepository interface.
type MockOauthAccessTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOauthAccessTokenRepositoryMockRecorder
}

// MockOauthAccessTokenRepositoryMockRecorder is the mock recorder for MockOauthAccessTokenRepository.
type MockOauthAccessTokenRepositoryMockRecorder struct {
	mock *MockOauthAccessTokenRepository
}

// NewMockOauthAccessTokenRepository creates a new mock instance.
func NewMockOauthAccessTokenRepository(ctrl *gomock.Controller) *MockOauthAccessTokenRepository {
	mock := &MockOauthAccessTokenRepository{ctrl: ctrl}
	mock.recorder = &MockOauthAccessTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthAccessTokenRepository) EXPECT() *MockOauthAccessTokenRepositoryMockRecorder {
	return m.recorder
}

// GetSession mocks base method.
func (m *MockOauthAccessTokenRepository) GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", ctx, accessToken)
	ret0, _ := ret[0].(*entities.OauthSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockOauthAccessTokenRepositoryMockRecorder) GetSession(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockOauthAccessTokenRepository)(nil).GetSession), ctx, accessToken)
}

// SetAccessToken mocks base method.
func (m *MockOauthAccessTokenRepository) SetAccessToken(ctx context.Context, oauthAccessToken *models.OauthAccessToken, accessToken string, expireIn int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessToken", ctx, oauthAccessToken, accessToken, expireIn)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAccessToken indicates an expected call of SetAccessToken.
func (mr *MockOauthAccessTokenRepositoryMockRecorder) SetAccessToken(ctx, oauthAccessToken, accessToken, expireIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessToken", reflect.TypeOf((*MockOauthAccessTokenRepository)(nil).SetAccessToken), ctx, oauthAccessToken, accessToken, expireIn)
}

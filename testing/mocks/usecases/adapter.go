// Code generated by MockGen. DO NOT EDIT.
// Source: usecases/adapter.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/hiroyky/famiphoto/entities"
	usecases "github.com/hiroyky/famiphoto/usecases"
)

// MockUserAdapter is a mock of UserAdapter interface.
type MockUserAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockUserAdapterMockRecorder
}

// MockUserAdapterMockRecorder is the mock recorder for MockUserAdapter.
type MockUserAdapterMockRecorder struct {
	mock *MockUserAdapter
}

// NewMockUserAdapter creates a new mock instance.
func NewMockUserAdapter(ctrl *gomock.Controller) *MockUserAdapter {
	mock := &MockUserAdapter{ctrl: ctrl}
	mock.recorder = &MockUserAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAdapter) EXPECT() *MockUserAdapterMockRecorder {
	return m.recorder
}

// CountUsers mocks base method.
func (m *MockUserAdapter) CountUsers(ctx context.Context, filter *usecases.UserFilter) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUsers", ctx, filter)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUsers indicates an expected call of CountUsers.
func (mr *MockUserAdapterMockRecorder) CountUsers(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUsers", reflect.TypeOf((*MockUserAdapter)(nil).CountUsers), ctx, filter)
}

// CreateUser mocks base method.
func (m *MockUserAdapter) CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user, password, isInitializedPassword, now)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserAdapterMockRecorder) CreateUser(ctx, user, password, isInitializedPassword, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserAdapter)(nil).CreateUser), ctx, user, password, isInitializedPassword, now)
}

// ExistUser mocks base method.
func (m *MockUserAdapter) ExistUser(ctx context.Context, userID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistUser", ctx, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistUser indicates an expected call of ExistUser.
func (mr *MockUserAdapterMockRecorder) ExistUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistUser", reflect.TypeOf((*MockUserAdapter)(nil).ExistUser), ctx, userID)
}

// GetUser mocks base method.
func (m *MockUserAdapter) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, userID)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserAdapterMockRecorder) GetUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserAdapter)(nil).GetUser), ctx, userID)
}

// GetUsers mocks base method.
func (m *MockUserAdapter) GetUsers(ctx context.Context, filter *usecases.UserFilter, limit, offset int) (entities.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, filter, limit, offset)
	ret0, _ := ret[0].(entities.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserAdapterMockRecorder) GetUsers(ctx, filter, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserAdapter)(nil).GetUsers), ctx, filter, limit, offset)
}

// MockPasswordService is a mock of PasswordService interface.
type MockPasswordService struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordServiceMockRecorder
}

// MockPasswordServiceMockRecorder is the mock recorder for MockPasswordService.
type MockPasswordServiceMockRecorder struct {
	mock *MockPasswordService
}

// NewMockPasswordService creates a new mock instance.
func NewMockPasswordService(ctrl *gomock.Controller) *MockPasswordService {
	mock := &MockPasswordService{ctrl: ctrl}
	mock.recorder = &MockPasswordServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordService) EXPECT() *MockPasswordServiceMockRecorder {
	return m.recorder
}

// GeneratePassword mocks base method.
func (m *MockPasswordService) GeneratePassword(length int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePassword", length)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePassword indicates an expected call of GeneratePassword.
func (mr *MockPasswordServiceMockRecorder) GeneratePassword(length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePassword", reflect.TypeOf((*MockPasswordService)(nil).GeneratePassword), length)
}

// HashPassword mocks base method.
func (m *MockPasswordService) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockPasswordServiceMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockPasswordService)(nil).HashPassword), password)
}

// MatchPassword mocks base method.
func (m *MockPasswordService) MatchPassword(password, hash string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchPassword", password, hash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchPassword indicates an expected call of MatchPassword.
func (mr *MockPasswordServiceMockRecorder) MatchPassword(password, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchPassword", reflect.TypeOf((*MockPasswordService)(nil).MatchPassword), password, hash)
}

// MockOauthClientAdapter is a mock of OauthClientAdapter interface.
type MockOauthClientAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockOauthClientAdapterMockRecorder
}

// MockOauthClientAdapterMockRecorder is the mock recorder for MockOauthClientAdapter.
type MockOauthClientAdapterMockRecorder struct {
	mock *MockOauthClientAdapter
}

// NewMockOauthClientAdapter creates a new mock instance.
func NewMockOauthClientAdapter(ctrl *gomock.Controller) *MockOauthClientAdapter {
	mock := &MockOauthClientAdapter{ctrl: ctrl}
	mock.recorder = &MockOauthClientAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthClientAdapter) EXPECT() *MockOauthClientAdapterMockRecorder {
	return m.recorder
}

// CreateOAuthClient mocks base method.
func (m *MockOauthClientAdapter) CreateOAuthClient(ctx context.Context, client *entities.OauthClient, clientSecret string) (*entities.OauthClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOAuthClient", ctx, client, clientSecret)
	ret0, _ := ret[0].(*entities.OauthClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOAuthClient indicates an expected call of CreateOAuthClient.
func (mr *MockOauthClientAdapterMockRecorder) CreateOAuthClient(ctx, client, clientSecret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOAuthClient", reflect.TypeOf((*MockOauthClientAdapter)(nil).CreateOAuthClient), ctx, client, clientSecret)
}

// ExistOauthClient mocks base method.
func (m *MockOauthClientAdapter) ExistOauthClient(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistOauthClient", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistOauthClient indicates an expected call of ExistOauthClient.
func (mr *MockOauthClientAdapterMockRecorder) ExistOauthClient(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistOauthClient", reflect.TypeOf((*MockOauthClientAdapter)(nil).ExistOauthClient), ctx, id)
}

// GetByOauthClientID mocks base method.
func (m *MockOauthClientAdapter) GetByOauthClientID(ctx context.Context, id string) (*entities.OauthClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByOauthClientID", ctx, id)
	ret0, _ := ret[0].(*entities.OauthClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByOauthClientID indicates an expected call of GetByOauthClientID.
func (mr *MockOauthClientAdapterMockRecorder) GetByOauthClientID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByOauthClientID", reflect.TypeOf((*MockOauthClientAdapter)(nil).GetByOauthClientID), ctx, id)
}

// MockOauthClientRedirectURLAdapter is a mock of OauthClientRedirectURLAdapter interface.
type MockOauthClientRedirectURLAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockOauthClientRedirectURLAdapterMockRecorder
}

// MockOauthClientRedirectURLAdapterMockRecorder is the mock recorder for MockOauthClientRedirectURLAdapter.
type MockOauthClientRedirectURLAdapterMockRecorder struct {
	mock *MockOauthClientRedirectURLAdapter
}

// NewMockOauthClientRedirectURLAdapter creates a new mock instance.
func NewMockOauthClientRedirectURLAdapter(ctrl *gomock.Controller) *MockOauthClientRedirectURLAdapter {
	mock := &MockOauthClientRedirectURLAdapter{ctrl: ctrl}
	mock.recorder = &MockOauthClientRedirectURLAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthClientRedirectURLAdapter) EXPECT() *MockOauthClientRedirectURLAdapterMockRecorder {
	return m.recorder
}

// CreateOAuthClientRedirectURL mocks base method.
func (m *MockOauthClientRedirectURLAdapter) CreateOAuthClientRedirectURL(ctx context.Context, url *entities.OAuthClientRedirectURL) (*entities.OAuthClientRedirectURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOAuthClientRedirectURL", ctx, url)
	ret0, _ := ret[0].(*entities.OAuthClientRedirectURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOAuthClientRedirectURL indicates an expected call of CreateOAuthClientRedirectURL.
func (mr *MockOauthClientRedirectURLAdapterMockRecorder) CreateOAuthClientRedirectURL(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOAuthClientRedirectURL", reflect.TypeOf((*MockOauthClientRedirectURLAdapter)(nil).CreateOAuthClientRedirectURL), ctx, url)
}

// GetOAuthClientRedirectURLsByOAuthClientID mocks base method.
func (m *MockOauthClientRedirectURLAdapter) GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) ([]*entities.OAuthClientRedirectURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOAuthClientRedirectURLsByOAuthClientID", ctx, oauthClientID)
	ret0, _ := ret[0].([]*entities.OAuthClientRedirectURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOAuthClientRedirectURLsByOAuthClientID indicates an expected call of GetOAuthClientRedirectURLsByOAuthClientID.
func (mr *MockOauthClientRedirectURLAdapterMockRecorder) GetOAuthClientRedirectURLsByOAuthClientID(ctx, oauthClientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOAuthClientRedirectURLsByOAuthClientID", reflect.TypeOf((*MockOauthClientRedirectURLAdapter)(nil).GetOAuthClientRedirectURLsByOAuthClientID), ctx, oauthClientID)
}

// MockOauthAccessTokenAdapter is a mock of OauthAccessTokenAdapter interface.
type MockOauthAccessTokenAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockOauthAccessTokenAdapterMockRecorder
}

// MockOauthAccessTokenAdapterMockRecorder is the mock recorder for MockOauthAccessTokenAdapter.
type MockOauthAccessTokenAdapterMockRecorder struct {
	mock *MockOauthAccessTokenAdapter
}

// NewMockOauthAccessTokenAdapter creates a new mock instance.
func NewMockOauthAccessTokenAdapter(ctrl *gomock.Controller) *MockOauthAccessTokenAdapter {
	mock := &MockOauthAccessTokenAdapter{ctrl: ctrl}
	mock.recorder = &MockOauthAccessTokenAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthAccessTokenAdapter) EXPECT() *MockOauthAccessTokenAdapterMockRecorder {
	return m.recorder
}

// SetClientCredentialAccessToken mocks base method.
func (m *MockOauthAccessTokenAdapter) SetClientCredentialAccessToken(ctx context.Context, clientID, accessToken string, expireAt int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetClientCredentialAccessToken", ctx, clientID, accessToken, expireAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetClientCredentialAccessToken indicates an expected call of SetClientCredentialAccessToken.
func (mr *MockOauthAccessTokenAdapterMockRecorder) SetClientCredentialAccessToken(ctx, clientID, accessToken, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientCredentialAccessToken", reflect.TypeOf((*MockOauthAccessTokenAdapter)(nil).SetClientCredentialAccessToken), ctx, clientID, accessToken, expireAt)
}

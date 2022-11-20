// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/http/graph/generated/generated.go

// Package mock_generated is a generated GoMock package.
package mock_generated

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	generated "github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	model "github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

// MockResolverRoot is a mock of ResolverRoot interface.
type MockResolverRoot struct {
	ctrl     *gomock.Controller
	recorder *MockResolverRootMockRecorder
}

// MockResolverRootMockRecorder is the mock recorder for MockResolverRoot.
type MockResolverRootMockRecorder struct {
	mock *MockResolverRoot
}

// NewMockResolverRoot creates a new mock instance.
func NewMockResolverRoot(ctrl *gomock.Controller) *MockResolverRoot {
	mock := &MockResolverRoot{ctrl: ctrl}
	mock.recorder = &MockResolverRootMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolverRoot) EXPECT() *MockResolverRootMockRecorder {
	return m.recorder
}

// Group mocks base method.
func (m *MockResolverRoot) Group() generated.GroupResolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group")
	ret0, _ := ret[0].(generated.GroupResolver)
	return ret0
}

// Group indicates an expected call of Group.
func (mr *MockResolverRootMockRecorder) Group() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockResolverRoot)(nil).Group))
}

// Mutation mocks base method.
func (m *MockResolverRoot) Mutation() generated.MutationResolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mutation")
	ret0, _ := ret[0].(generated.MutationResolver)
	return ret0
}

// Mutation indicates an expected call of Mutation.
func (mr *MockResolverRootMockRecorder) Mutation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mutation", reflect.TypeOf((*MockResolverRoot)(nil).Mutation))
}

// OauthClient mocks base method.
func (m *MockResolverRoot) OauthClient() generated.OauthClientResolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OauthClient")
	ret0, _ := ret[0].(generated.OauthClientResolver)
	return ret0
}

// OauthClient indicates an expected call of OauthClient.
func (mr *MockResolverRootMockRecorder) OauthClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OauthClient", reflect.TypeOf((*MockResolverRoot)(nil).OauthClient))
}

// Photo mocks base method.
func (m *MockResolverRoot) Photo() generated.PhotoResolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Photo")
	ret0, _ := ret[0].(generated.PhotoResolver)
	return ret0
}

// Photo indicates an expected call of Photo.
func (mr *MockResolverRootMockRecorder) Photo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Photo", reflect.TypeOf((*MockResolverRoot)(nil).Photo))
}

// PhotoFile mocks base method.
func (m *MockResolverRoot) PhotoFile() generated.PhotoFileResolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhotoFile")
	ret0, _ := ret[0].(generated.PhotoFileResolver)
	return ret0
}

// PhotoFile indicates an expected call of PhotoFile.
func (mr *MockResolverRootMockRecorder) PhotoFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhotoFile", reflect.TypeOf((*MockResolverRoot)(nil).PhotoFile))
}

// Query mocks base method.
func (m *MockResolverRoot) Query() generated.QueryResolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query")
	ret0, _ := ret[0].(generated.QueryResolver)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockResolverRootMockRecorder) Query() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockResolverRoot)(nil).Query))
}

// User mocks base method.
func (m *MockResolverRoot) User() generated.UserResolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(generated.UserResolver)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockResolverRootMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockResolverRoot)(nil).User))
}

// MockGroupResolver is a mock of GroupResolver interface.
type MockGroupResolver struct {
	ctrl     *gomock.Controller
	recorder *MockGroupResolverMockRecorder
}

// MockGroupResolverMockRecorder is the mock recorder for MockGroupResolver.
type MockGroupResolverMockRecorder struct {
	mock *MockGroupResolver
}

// NewMockGroupResolver creates a new mock instance.
func NewMockGroupResolver(ctrl *gomock.Controller) *MockGroupResolver {
	mock := &MockGroupResolver{ctrl: ctrl}
	mock.recorder = &MockGroupResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupResolver) EXPECT() *MockGroupResolverMockRecorder {
	return m.recorder
}

// UserPagination mocks base method.
func (m *MockGroupResolver) UserPagination(ctx context.Context, obj *model.Group, limit, offset *int) (*model.UserPagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserPagination", ctx, obj, limit, offset)
	ret0, _ := ret[0].(*model.UserPagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserPagination indicates an expected call of UserPagination.
func (mr *MockGroupResolverMockRecorder) UserPagination(ctx, obj, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserPagination", reflect.TypeOf((*MockGroupResolver)(nil).UserPagination), ctx, obj, limit, offset)
}

// MockMutationResolver is a mock of MutationResolver interface.
type MockMutationResolver struct {
	ctrl     *gomock.Controller
	recorder *MockMutationResolverMockRecorder
}

// MockMutationResolverMockRecorder is the mock recorder for MockMutationResolver.
type MockMutationResolverMockRecorder struct {
	mock *MockMutationResolver
}

// NewMockMutationResolver creates a new mock instance.
func NewMockMutationResolver(ctrl *gomock.Controller) *MockMutationResolver {
	mock := &MockMutationResolver{ctrl: ctrl}
	mock.recorder = &MockMutationResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutationResolver) EXPECT() *MockMutationResolverMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method.
func (m *MockMutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, input)
	ret0, _ := ret[0].(*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockMutationResolverMockRecorder) CreateGroup(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockMutationResolver)(nil).CreateGroup), ctx, input)
}

// CreateOauthClient mocks base method.
func (m *MockMutationResolver) CreateOauthClient(ctx context.Context, input model.CreateOauthClientInput) (*model.OauthClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOauthClient", ctx, input)
	ret0, _ := ret[0].(*model.OauthClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOauthClient indicates an expected call of CreateOauthClient.
func (mr *MockMutationResolverMockRecorder) CreateOauthClient(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOauthClient", reflect.TypeOf((*MockMutationResolver)(nil).CreateOauthClient), ctx, input)
}

// CreateUser mocks base method.
func (m *MockMutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, input)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockMutationResolverMockRecorder) CreateUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockMutationResolver)(nil).CreateUser), ctx, input)
}

// MockOauthClientResolver is a mock of OauthClientResolver interface.
type MockOauthClientResolver struct {
	ctrl     *gomock.Controller
	recorder *MockOauthClientResolverMockRecorder
}

// MockOauthClientResolverMockRecorder is the mock recorder for MockOauthClientResolver.
type MockOauthClientResolverMockRecorder struct {
	mock *MockOauthClientResolver
}

// NewMockOauthClientResolver creates a new mock instance.
func NewMockOauthClientResolver(ctrl *gomock.Controller) *MockOauthClientResolver {
	mock := &MockOauthClientResolver{ctrl: ctrl}
	mock.recorder = &MockOauthClientResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthClientResolver) EXPECT() *MockOauthClientResolverMockRecorder {
	return m.recorder
}

// RedirectUrls mocks base method.
func (m *MockOauthClientResolver) RedirectUrls(ctx context.Context, obj *model.OauthClient) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedirectUrls", ctx, obj)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RedirectUrls indicates an expected call of RedirectUrls.
func (mr *MockOauthClientResolverMockRecorder) RedirectUrls(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedirectUrls", reflect.TypeOf((*MockOauthClientResolver)(nil).RedirectUrls), ctx, obj)
}

// MockPhotoResolver is a mock of PhotoResolver interface.
type MockPhotoResolver struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoResolverMockRecorder
}

// MockPhotoResolverMockRecorder is the mock recorder for MockPhotoResolver.
type MockPhotoResolverMockRecorder struct {
	mock *MockPhotoResolver
}

// NewMockPhotoResolver creates a new mock instance.
func NewMockPhotoResolver(ctrl *gomock.Controller) *MockPhotoResolver {
	mock := &MockPhotoResolver{ctrl: ctrl}
	mock.recorder = &MockPhotoResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoResolver) EXPECT() *MockPhotoResolverMockRecorder {
	return m.recorder
}

// ExifData mocks base method.
func (m *MockPhotoResolver) ExifData(ctx context.Context, obj *model.Photo) ([]*model.PhotoExif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExifData", ctx, obj)
	ret0, _ := ret[0].([]*model.PhotoExif)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExifData indicates an expected call of ExifData.
func (mr *MockPhotoResolverMockRecorder) ExifData(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExifData", reflect.TypeOf((*MockPhotoResolver)(nil).ExifData), ctx, obj)
}

// Files mocks base method.
func (m *MockPhotoResolver) Files(ctx context.Context, obj *model.Photo) ([]*model.PhotoFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Files", ctx, obj)
	ret0, _ := ret[0].([]*model.PhotoFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Files indicates an expected call of Files.
func (mr *MockPhotoResolverMockRecorder) Files(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Files", reflect.TypeOf((*MockPhotoResolver)(nil).Files), ctx, obj)
}

// Group mocks base method.
func (m *MockPhotoResolver) Group(ctx context.Context, obj *model.Photo) (*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", ctx, obj)
	ret0, _ := ret[0].(*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Group indicates an expected call of Group.
func (mr *MockPhotoResolverMockRecorder) Group(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockPhotoResolver)(nil).Group), ctx, obj)
}

// Owner mocks base method.
func (m *MockPhotoResolver) Owner(ctx context.Context, obj *model.Photo) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner", ctx, obj)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Owner indicates an expected call of Owner.
func (mr *MockPhotoResolverMockRecorder) Owner(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockPhotoResolver)(nil).Owner), ctx, obj)
}

// MockPhotoFileResolver is a mock of PhotoFileResolver interface.
type MockPhotoFileResolver struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoFileResolverMockRecorder
}

// MockPhotoFileResolverMockRecorder is the mock recorder for MockPhotoFileResolver.
type MockPhotoFileResolverMockRecorder struct {
	mock *MockPhotoFileResolver
}

// NewMockPhotoFileResolver creates a new mock instance.
func NewMockPhotoFileResolver(ctrl *gomock.Controller) *MockPhotoFileResolver {
	mock := &MockPhotoFileResolver{ctrl: ctrl}
	mock.recorder = &MockPhotoFileResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoFileResolver) EXPECT() *MockPhotoFileResolverMockRecorder {
	return m.recorder
}

// Group mocks base method.
func (m *MockPhotoFileResolver) Group(ctx context.Context, obj *model.PhotoFile) (*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", ctx, obj)
	ret0, _ := ret[0].(*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Group indicates an expected call of Group.
func (mr *MockPhotoFileResolverMockRecorder) Group(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockPhotoFileResolver)(nil).Group), ctx, obj)
}

// Owner mocks base method.
func (m *MockPhotoFileResolver) Owner(ctx context.Context, obj *model.PhotoFile) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner", ctx, obj)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Owner indicates an expected call of Owner.
func (mr *MockPhotoFileResolverMockRecorder) Owner(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockPhotoFileResolver)(nil).Owner), ctx, obj)
}

// Photo mocks base method.
func (m *MockPhotoFileResolver) Photo(ctx context.Context, obj *model.PhotoFile) (*model.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Photo", ctx, obj)
	ret0, _ := ret[0].(*model.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Photo indicates an expected call of Photo.
func (mr *MockPhotoFileResolverMockRecorder) Photo(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Photo", reflect.TypeOf((*MockPhotoFileResolver)(nil).Photo), ctx, obj)
}

// MockQueryResolver is a mock of QueryResolver interface.
type MockQueryResolver struct {
	ctrl     *gomock.Controller
	recorder *MockQueryResolverMockRecorder
}

// MockQueryResolverMockRecorder is the mock recorder for MockQueryResolver.
type MockQueryResolverMockRecorder struct {
	mock *MockQueryResolver
}

// NewMockQueryResolver creates a new mock instance.
func NewMockQueryResolver(ctrl *gomock.Controller) *MockQueryResolver {
	mock := &MockQueryResolver{ctrl: ctrl}
	mock.recorder = &MockQueryResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryResolver) EXPECT() *MockQueryResolverMockRecorder {
	return m.recorder
}

// ExistUserID mocks base method.
func (m *MockQueryResolver) ExistUserID(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistUserID", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistUserID indicates an expected call of ExistUserID.
func (mr *MockQueryResolverMockRecorder) ExistUserID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistUserID", reflect.TypeOf((*MockQueryResolver)(nil).ExistUserID), ctx, id)
}

// Me mocks base method.
func (m *MockQueryResolver) Me(ctx context.Context) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Me", ctx)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Me indicates an expected call of Me.
func (mr *MockQueryResolverMockRecorder) Me(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Me", reflect.TypeOf((*MockQueryResolver)(nil).Me), ctx)
}

// Photo mocks base method.
func (m *MockQueryResolver) Photo(ctx context.Context, id string) (*model.Photo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Photo", ctx, id)
	ret0, _ := ret[0].(*model.Photo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Photo indicates an expected call of Photo.
func (mr *MockQueryResolverMockRecorder) Photo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Photo", reflect.TypeOf((*MockQueryResolver)(nil).Photo), ctx, id)
}

// PhotoFile mocks base method.
func (m *MockQueryResolver) PhotoFile(ctx context.Context, id string) (*model.PhotoFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhotoFile", ctx, id)
	ret0, _ := ret[0].(*model.PhotoFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PhotoFile indicates an expected call of PhotoFile.
func (mr *MockQueryResolverMockRecorder) PhotoFile(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhotoFile", reflect.TypeOf((*MockQueryResolver)(nil).PhotoFile), ctx, id)
}

// PhotoFiles mocks base method.
func (m *MockQueryResolver) PhotoFiles(ctx context.Context, photoID string) ([]*model.PhotoFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhotoFiles", ctx, photoID)
	ret0, _ := ret[0].([]*model.PhotoFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PhotoFiles indicates an expected call of PhotoFiles.
func (mr *MockQueryResolverMockRecorder) PhotoFiles(ctx, photoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhotoFiles", reflect.TypeOf((*MockQueryResolver)(nil).PhotoFiles), ctx, photoID)
}

// Photos mocks base method.
func (m *MockQueryResolver) Photos(ctx context.Context, id, ownerID, groupID *string, limit, offset *int) (*model.PhotoPagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Photos", ctx, id, ownerID, groupID, limit, offset)
	ret0, _ := ret[0].(*model.PhotoPagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Photos indicates an expected call of Photos.
func (mr *MockQueryResolverMockRecorder) Photos(ctx, id, ownerID, groupID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Photos", reflect.TypeOf((*MockQueryResolver)(nil).Photos), ctx, id, ownerID, groupID, limit, offset)
}

// User mocks base method.
func (m *MockQueryResolver) User(ctx context.Context, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User.
func (mr *MockQueryResolverMockRecorder) User(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockQueryResolver)(nil).User), ctx, id)
}

// Users mocks base method.
func (m *MockQueryResolver) Users(ctx context.Context, id *string, limit, offset *int) (*model.UserPagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users", ctx, id, limit, offset)
	ret0, _ := ret[0].(*model.UserPagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Users indicates an expected call of Users.
func (mr *MockQueryResolverMockRecorder) Users(ctx, id, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockQueryResolver)(nil).Users), ctx, id, limit, offset)
}

// MockUserResolver is a mock of UserResolver interface.
type MockUserResolver struct {
	ctrl     *gomock.Controller
	recorder *MockUserResolverMockRecorder
}

// MockUserResolverMockRecorder is the mock recorder for MockUserResolver.
type MockUserResolverMockRecorder struct {
	mock *MockUserResolver
}

// NewMockUserResolver creates a new mock instance.
func NewMockUserResolver(ctrl *gomock.Controller) *MockUserResolver {
	mock := &MockUserResolver{ctrl: ctrl}
	mock.recorder = &MockUserResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserResolver) EXPECT() *MockUserResolverMockRecorder {
	return m.recorder
}

// BelongGroups mocks base method.
func (m *MockUserResolver) BelongGroups(ctx context.Context, obj *model.User) ([]*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BelongGroups", ctx, obj)
	ret0, _ := ret[0].([]*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BelongGroups indicates an expected call of BelongGroups.
func (mr *MockUserResolverMockRecorder) BelongGroups(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BelongGroups", reflect.TypeOf((*MockUserResolver)(nil).BelongGroups), ctx, obj)
}

// Password mocks base method.
func (m *MockUserResolver) Password(ctx context.Context, obj *model.User) (*model.UserPassword, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Password", ctx, obj)
	ret0, _ := ret[0].(*model.UserPassword)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Password indicates an expected call of Password.
func (mr *MockUserResolverMockRecorder) Password(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Password", reflect.TypeOf((*MockUserResolver)(nil).Password), ctx, obj)
}

package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/hiroyky/famiphoto/utils/cast"
	"time"
)

type UserAdapter interface {
	GetUser(ctx context.Context, userID string) (*entities.User, error)
	GetUsers(ctx context.Context, filter *filters.UserFilter, limit, offset int) (entities.UserList, error)
	CountUsers(ctx context.Context, filter *filters.UserFilter) (int, error)
	ExistUser(ctx context.Context, userID string) (bool, error)
	GetUsersBelongingGroup(ctx context.Context, groupID string, limit, offset int) (entities.UserList, error)
	CountUsersBelongingGroup(ctx context.Context, groupID string) (int, error)
	CreateUser(ctx context.Context, user *entities.User, groupName, password string, isInitializedPassword bool, now time.Time) (*entities.User, error)
	GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error)
}

func NewUserAdapter(userRepo repositories.UserRepository, groupRepo repositories.GroupRepository, userPasswordRepo repositories.UserPasswordRepository) UserAdapter {
	return &userAdapter{
		userRepo:         userRepo,
		groupRepo:        groupRepo,
		userPasswordRepo: userPasswordRepo,
	}
}

type userAdapter struct {
	userRepo         repositories.UserRepository
	groupRepo        repositories.GroupRepository
	userPasswordRepo repositories.UserPasswordRepository
}

func (a *userAdapter) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	dbUser, err := a.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return a.toUserEntity(dbUser), nil
}

func (a *userAdapter) GetUsers(ctx context.Context, filter *filters.UserFilter, limit, offset int) (entities.UserList, error) {
	dbUsers, err := a.userRepo.GetUsers(ctx, filter, limit, offset)
	if err != nil {
		return nil, err
	}
	return array.Map(dbUsers, a.toUserEntity), nil
}

func (a *userAdapter) CountUsers(ctx context.Context, filter *filters.UserFilter) (int, error) {
	return a.userRepo.CountUsers(ctx, filter)
}

func (a *userAdapter) ExistUser(ctx context.Context, userID string) (bool, error) {
	return a.userRepo.ExistUser(ctx, userID)
}

func (a *userAdapter) GetUsersBelongingGroup(ctx context.Context, groupID string, limit, offset int) (entities.UserList, error) {
	dbUsers, err := a.groupRepo.GetUsersByGroupID(ctx, groupID, limit, offset)
	if err != nil {
		return nil, err
	}
	return array.Map(dbUsers, a.toUserEntity), nil
}

func (a *userAdapter) CountUsersBelongingGroup(ctx context.Context, groupID string) (int, error) {
	return a.groupRepo.CountUsersByGroupID(ctx, groupID)
}

func (a *userAdapter) CreateUser(ctx context.Context, user *entities.User, groupName, password string, isInitializedPassword bool, now time.Time) (*entities.User, error) {
	dbUser := &dbmodels.User{
		UserID: user.UserID,
		Name:   user.Name,
		Status: a.toDBUserStatus(user.Status),
	}
	dbPassword := &dbmodels.UserPassword{
		UserID:         user.UserID,
		Password:       password,
		LastModifiedAt: now,
		IsInitialized:  cast.BoolToInt8(isInitializedPassword),
	}
	dbGroup := &dbmodels.Group{
		GroupID: user.UserID,
		Name:    groupName,
	}

	dstDBUser, err := a.userRepo.CreateUser(ctx, dbUser, dbGroup, dbPassword)
	if err != nil {
		return nil, err
	}

	return a.toUserEntity(dstDBUser), nil
}

func (a *userAdapter) GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error) {
	up, err := a.userPasswordRepo.GetUserPassword(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &entities.UserPassword{
		UserId:         up.UserID,
		Password:       up.Password,
		LastModifiedAt: up.LastModifiedAt,
		IsInitialized:  cast.IntToBool(up.IsInitialized),
	}, nil
}

func (a *userAdapter) toDBUserStatus(s entities.UserStatus) int {
	return int(s)
}

func (a *userAdapter) toEntityUserStatus(s int) entities.UserStatus {
	return entities.UserStatus(s)
}

func (a *userAdapter) toUserEntity(user *dbmodels.User) *entities.User {
	return &entities.User{
		UserID: user.UserID,
		Name:   user.Name,
		Status: a.toEntityUserStatus(user.Status),
	}
}

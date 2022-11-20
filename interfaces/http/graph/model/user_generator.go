package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/hiroyky/famiphoto/utils/gql"
)

const UserName = "User"

func NewUser(u *entities.User) *User {
	var st UserStatus
	switch u.Status {
	case entities.UserStatusActive:
		st = UserStatusActive
	case entities.UserStatusWithdrawal:
		st = UserStatusWithdrawal
	}

	return &User{
		ID:     gql.EncodeStrID(UserName, u.UserID),
		UserID: u.UserID,
		Name:   u.Name,
		Status: st,
	}
}

func NewUserPagination(users entities.UserList, total, limit, offset int) *UserPagination {
	return &UserPagination{
		PageInfo: newPaginationInfo(total, len(users), limit, offset),
		Nodes:    cast.Array(users, NewUser),
	}
}

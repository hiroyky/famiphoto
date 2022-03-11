package model

import (
	"github.com/hiroyky/famiphoto/entities"
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
		Name:   u.Name,
		Status: st,
	}
}

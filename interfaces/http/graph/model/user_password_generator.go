package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/gql"
	"time"
)

const UserPasswordName = "UserPassword"

func NewUserPassword(p *entities.UserPassword) *UserPassword {
	return &UserPassword{
		ID:            gql.EncodeStrID(UserPasswordName, p.UserId),
		LastModified:  p.LastModifiedAt.Format(time.RFC3339),
		IsInitialized: p.IsInitialized,
	}
}
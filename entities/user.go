package entities

import "time"

type User struct {
	UserID string
	Name   string
	Status UserStatus
}

type UserList []*User

type UserStatus int

const (
	UserStatusActive     UserStatus = 1
	UserStatusWithdrawal UserStatus = 2
)

type UserPassword struct {
	UserId         string
	Password       string
	LastModifiedAt time.Time
	IsInitialized  bool
}

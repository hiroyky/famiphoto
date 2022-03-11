package usecases

import (
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type UserUseCase interface {
}

type userUseCase struct {
	userAdapter UserAdapter
}

func (u *userUseCase) CreateUser(userID, name string, password string, now time.Time) (*entities.User, error) {
	user := &entities.User{
		UserID: userID,
		Name:   name,
		Status: entities.UserStatusActive,
	}

	createdUser, err := u.userAdapter.CreateUser(user, password, true, now)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

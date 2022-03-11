package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type UserAdapter interface {
	GetUser(ctx context.Context, userID string) (*entities.User, error)
	ExistUser(ctx context.Context, userID string) (bool, error)
	CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	MatchPassword(password string, hash string) (bool, error)
	GeneratePassword() (string, error)
}

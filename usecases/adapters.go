package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type UserAdapter interface {
	GetUser(ctx context.Context, userID string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error)
}

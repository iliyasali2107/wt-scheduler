package user_service

import (
	"context"

	"github.com/iliyasali2107/wt-scheduler/internal/models"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user models.User) error
}

type UserService struct {
	UserStorage UserStorage
}

func NewUserService(ur UserStorage) UserService {
	return UserService{
		UserStorage: ur,
	}
}

package auth_service

import (
	"context"
	"fmt"

	"github.com/iliyasali2107/wt-scheduler/internal/dto"
	"github.com/iliyasali2107/wt-scheduler/internal/models"
)

type UserStorage interface {
	CreateUser(ctx context.Context, req models.User) (models.User, error)
}

type AuthService struct {
	userStorage UserStorage
}

func NewAuthService(ur UserStorage) AuthService {
	return AuthService{
		userStorage: ur,
	}
}

func (as *AuthService) Register(ctx context.Context, req dto.RegisterRequestBody) (*dto.RegisterResponseBody, error) {
	if req.Name == "" || req.Password == "" {
		return nil, fmt.Errorf("invalid credentials")
	}

	if req.Number == "" && req.TelegramId == "" {
		return nil, fmt.Errorf("number or telegramId is empty")
	}

	user := models.User

	res, err := as.userStorage.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponseBody{User: res}, nil

}

func (as *AuthService) Login(ctx context.Context, req dto.LoginRequestBody) (*dto.LoginResponsBody, error) {
	if req.Name == "" || req.Password == "" {
		return err
	}

}

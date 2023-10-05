package auth_service

import (
	"context"

	"github.com/iliyasali2107/wt-scheduler/internal/models"
)

type AuthService interface {
	Register(ctx context.Context, user models.User) error
	Login(ctx context.Context, user models.User) error
	Validate(ctx context.Context)
}

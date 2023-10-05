package postgres

import (
	"context"

	"github.com/iliyasali2107/wt-scheduler/internal/models"
	// "github.com/jackc/pgx/v5"
)

type PostgresStorage struct {
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (ps *PostgresStorage) CreateUser(ctx context.Context, user models.User) error {
	return nil
}

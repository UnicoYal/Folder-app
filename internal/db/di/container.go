package di

import (
	"context"
	"fmt"

	"github.com/folder-app/config"
	"github.com/folder-app/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Container for postgres init
type PostgresContainer struct {
	pg *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.PostgresConfig) (*PostgresContainer, error) {
	pg, err := db.NewPostgresPool(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("db.NewPostgres: %w", err)
	}

	return &PostgresContainer{pg: pg}, nil
}

func (c *PostgresContainer) ProvidePostgres() *pgxpool.Pool {
	return c.pg
}

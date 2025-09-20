package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FoldersRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *FoldersRepository {
	return &FoldersRepository{
		db: db,
	}
}

func (r *FoldersRepository) CreateUser(ctx context.Context) {
	r.db.Close()
}

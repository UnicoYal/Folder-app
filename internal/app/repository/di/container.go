package di

import (
	"github.com/folder-app/internal/app/repository"
	"github.com/folder-app/internal/app/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Container for folders repository init
type FoldersRepositoryContainer struct {
	repo usecase.Repository
}

type DI interface {
	ProvidePostgres() *pgxpool.Pool
}

func New(di DI) *FoldersRepositoryContainer {
	return &FoldersRepositoryContainer{
		repo: repository.NewRepository(di.ProvidePostgres()),
	}
}

func (c *FoldersRepositoryContainer) ProvideFoldersRepository() usecase.Repository {
	return c.repo
}

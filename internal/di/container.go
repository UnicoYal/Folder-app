package di

import (
	"context"
	"fmt"

	"github.com/folder-app/config"
	repositorydi "github.com/folder-app/internal/app/repository/di"
	usecasedi "github.com/folder-app/internal/app/usecase/di"
	postgresdi "github.com/folder-app/internal/db/di"
	"github.com/folder-app/pkg/di"
	"github.com/rs/zerolog"
)

// Container for app init
type Container struct {
	*di.BaseContainer

	*FolderContainer
}

func New(cfg *config.Config) (*Container, error) {
	var (
		c   Container
		err error
	)

	c.BaseContainer, err = di.NewBaseContainer(cfg)
	if err != nil {
		return nil, fmt.Errorf("di.NewBaseContainer: %w", err)
	}

	c.FolderContainer, err = NewFolderContainer(c, cfg)
	if err != nil {
		return nil, fmt.Errorf("NewFolderContainer: %w", err)
	}

	return &c, err
}

type FolderContainer struct {
	*postgresdi.PostgresContainer
	*repositorydi.FoldersRepositoryContainer
	*usecasedi.FoldersUsecaseContainer
}

type DI interface {
	ProvideLogger() *zerolog.Logger
}

func NewFolderContainer(di DI, cfg *config.Config) (*FolderContainer, error) {
	var (
		err error
		c   = &struct {
			DI
			*FolderContainer
		}{
			DI:              di,
			FolderContainer: &FolderContainer{},
		}
	)

	ctx := context.Background()

	c.PostgresContainer, err = postgresdi.New(ctx, &cfg.PostgresConfig)
	if err != nil {
		return nil, fmt.Errorf("postgresdi.New: %w", err)
	}

	c.FoldersRepositoryContainer = repositorydi.New(c)
	c.FoldersUsecaseContainer = usecasedi.New(c)

	return c.FolderContainer, err
}

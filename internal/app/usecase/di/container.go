package di

import (
	v1 "github.com/folder-app/internal/app/api/http/v1"
	"github.com/folder-app/internal/app/usecase"
)

// Container for folders usecase init
type FoldersUsecaseContainer struct {
	usecase v1.Usecase
}

type DI interface {
	ProvideFoldersRepository() usecase.Repository
}

func New(di DI) *FoldersUsecaseContainer {
	return &FoldersUsecaseContainer{
		usecase: usecase.New(di.ProvideFoldersRepository()),
	}
}

func (c *FoldersUsecaseContainer) ProvideFoldersUsecase() v1.Usecase {
	return c.usecase
}

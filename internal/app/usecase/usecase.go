package usecase

import (
	"context"
	"fmt"
)

type Repository interface {
	CreateUser(ctx context.Context)
}

type FoldersUsecase struct {
	repository Repository
}

func New(repo Repository) *FoldersUsecase {
	return &FoldersUsecase{
		repository: repo,
	}
}

func (u *FoldersUsecase) CreateUser(ctx context.Context) {
	fmt.Println("Create usecase")
	u.repository.CreateUser(ctx)
}

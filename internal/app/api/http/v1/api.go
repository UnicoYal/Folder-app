package v1

import (
	"context"
	"errors"
	"net/http"
)

var ErrBadUsecase = errors.New("bad usecase")

type Usecase interface {
	CreateUser(ctx context.Context)
}

type API struct {
	usecase Usecase
}

func New(usecase Usecase) (*API, error) {
	if usecase == nil {
		return nil, ErrBadUsecase
	}

	return &API{
		usecase: usecase,
	}, nil
}

func (a *API) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(123)
	w.Header().Set("Asa", "ASda")
	a.usecase.CreateUser(r.Context())
}

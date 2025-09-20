package http

import (
	"fmt"
	"net/http"

	"github.com/folder-app/internal/app/api/http/router"
	v1 "github.com/folder-app/internal/app/api/http/v1"
)

type DI interface {
	ProvideHTTPMux() *http.ServeMux
	ProvideFoldersUsecase() v1.Usecase
}

func Setup(di DI) error {
	api, err := v1.New(di.ProvideFoldersUsecase())
	if err != nil {
		return fmt.Errorf("v1.New: %w", err)
	}

	router.SetupV1(api, di.ProvideHTTPMux())
	return nil
}

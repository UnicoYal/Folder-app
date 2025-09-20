package app

import (
	"fmt"
	"net/http"

	httpserver "github.com/folder-app/internal/app/api/http"
	v1 "github.com/folder-app/internal/app/api/http/v1"
	"github.com/rs/zerolog"
)

type DI interface {
	ProvideLogger() *zerolog.Logger
	ProvideHTTPMux() *http.ServeMux
	ProvideFoldersUsecase() v1.Usecase
	ProvideHTTPServer() *http.Server
}

func Setup(di DI) (*http.Server, error) {
	lg := di.ProvideLogger()

	lg.Info().Msg("app.Setup starting")

	err := httpserver.Setup(di)
	if err != nil {
		return nil, fmt.Errorf("http.Setup: %w", err)
	}

	return di.ProvideHTTPServer(), nil
}

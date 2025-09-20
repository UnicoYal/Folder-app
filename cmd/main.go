package main

import (
	"errors"
	"net/http"

	"github.com/folder-app/config"
	"github.com/folder-app/internal/app"
	"github.com/folder-app/internal/di"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot init config")
	}

	container, err := di.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot init container")
	}

	lg := container.ProvideLogger()
	lg.Info().Msg("container inited")

	srv, err := app.Setup(container)
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot start service")
	}

	lg.Info().Msgf("Listening: %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			lg.Info().Msg("server closed under request")
		} else {
			lg.Info().Msgf("server stopped: %v", err)
		}
	}
}

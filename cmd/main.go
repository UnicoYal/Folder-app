package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/folder-app/config"
	httpsrv "github.com/folder-app/internal/app/api/http"
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

	err = httpsrv.Setup(container)
	if err != nil {
		lg.Fatal().Err(err).Msg("cannot start service")
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Listener.Port),
		Handler:      container.ProvideHTTPServer(),
		ReadTimeout:  cfg.Listener.Timeout,
		WriteTimeout: cfg.Listener.Timeout,
		IdleTimeout:  cfg.Listener.IdleTimeout,
	}

	lg.Info().Msgf("Listening: %d", cfg.Listener.Port)

	if err := srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			lg.Info().Msg("server closed under request")
		} else {
			lg.Info().Msgf("server stopped: %v", err)
		}
	}
}

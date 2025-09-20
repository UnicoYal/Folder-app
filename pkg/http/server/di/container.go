package di

import (
	"fmt"
	"net/http"

	"github.com/folder-app/config"
)

// Container for http server init
type HTTPServerContainer struct {
	srv *http.Server
}

type DI interface {
	ProvideHTTPMux() *http.ServeMux
}

func New(di DI, cfg *config.Listener) *HTTPServerContainer {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      di.ProvideHTTPMux(),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &HTTPServerContainer{srv: srv}
}

func (c *HTTPServerContainer) ProvideHTTPServer() *http.Server {
	return c.srv
}

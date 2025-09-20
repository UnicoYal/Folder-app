package di

import (
	"net/http"

	"github.com/rs/zerolog"
)

// Container for http mux init
type HTTPContainer struct {
	mux *http.ServeMux
}

type DI interface {
	ProvideLogger() *zerolog.Logger
}

func New(di DI) *HTTPContainer {
	mux := http.NewServeMux()

	return &HTTPContainer{mux: mux}
}

func (c *HTTPContainer) ProvideHTTPServer() *http.ServeMux {
	return c.mux
}

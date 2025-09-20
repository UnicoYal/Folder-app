package di

import (
	"net/http"
)

// Container for http mux init
type HTTPMuxContainer struct {
	mux *http.ServeMux
}

func New() *HTTPMuxContainer {
	mux := http.NewServeMux()

	return &HTTPMuxContainer{mux: mux}
}

func (c *HTTPMuxContainer) ProvideHTTPMux() *http.ServeMux {
	return c.mux
}

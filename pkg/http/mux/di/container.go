package di

import (
	"net/http"
)

// Container for http mux init
type HTTPMuxContainer struct {
	mux *http.ServeMux
}

func New() *HTTPMuxContainer {
	return &HTTPMuxContainer{mux: http.NewServeMux()}
}

func (c *HTTPMuxContainer) ProvideHTTPMux() *http.ServeMux {
	return c.mux
}

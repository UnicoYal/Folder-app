package di

import (
	"github.com/folder-app/config"
	"github.com/folder-app/internal/logger"
	muxdi "github.com/folder-app/pkg/http/mux/di"
	serverdi "github.com/folder-app/pkg/http/server/di"
)

// Container for initing base components
type BaseContainer struct {
	*logger.LoggerContainer
	*muxdi.HTTPMuxContainer
	*serverdi.HTTPServerContainer
}

func NewBaseContainer(cfg *config.Config) (*BaseContainer, error) {
	var (
		c   BaseContainer
		err error
	)

	c.LoggerContainer = logger.New()
	c.HTTPMuxContainer = muxdi.New()

	c.HTTPServerContainer = serverdi.New(c, &cfg.Listener)

	return &c, err
}

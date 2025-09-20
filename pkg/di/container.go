package di

import (
	"github.com/folder-app/internal/logger"
	"github.com/folder-app/pkg/http/di"
)

// Container for initing base components
type BaseContainer struct {
	*logger.LoggerContainer
	*di.HTTPContainer
}

func NewBaseContainer() (*BaseContainer, error) {
	var (
		c   BaseContainer
		err error
	)

	c.LoggerContainer = logger.New()
	c.HTTPContainer = di.New(c)

	return &c, err
}

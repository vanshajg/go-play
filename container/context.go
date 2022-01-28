package container

import (
	"github.com/vanshajg/go-play/config"
	"github.com/vanshajg/go-play/logger"
	"github.com/vanshajg/go-play/repository"
)

type Container interface {
	GetRepository() repository.Repository
	GetConfig() *config.Config
	GetLogger() *logger.Logger
	GetEnv() string
}

type container struct {
	rep    repository.Repository
	config *config.Config
	logger *logger.Logger
	env    string
}

func NewContainer(rep repository.Repository, config *config.Config, logger *logger.Logger, env string) Container {
	return &container{rep: rep, config: config, logger: logger, env: env}
}

// GetRepository returns the object of repository.
func (c *container) GetRepository() repository.Repository {
	return c.rep
}

// GetConfig returns the object of configuration.
func (c *container) GetConfig() *config.Config {
	return c.config
}

// GetLogger returns the object of logger.
func (c *container) GetLogger() *logger.Logger {
	return c.logger
}

// GetEnv returns the running environment.
func (c *container) GetEnv() string {
	return c.env
}

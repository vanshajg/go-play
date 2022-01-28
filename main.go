package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vanshajg/go-play/config"
	"github.com/vanshajg/go-play/container"
	"github.com/vanshajg/go-play/logger"
	"github.com/vanshajg/go-play/migration"
	"github.com/vanshajg/go-play/offline"
	"github.com/vanshajg/go-play/repository"
	"github.com/vanshajg/go-play/router"
	"github.com/vanshajg/go-play/service"
)

func main() {
	e := echo.New()
	config, env := config.Load()
	logger := logger.NewLogger(env)
	logger.GetZapLogger().Infof("Loaded from config: application.%s.yaml", env)

	rep := repository.NewCommentRepository(logger, config)
	container := container.NewContainer(rep, config, logger, env)
	commentService := service.NewCommentService(container)

	migration.CreateDatabase(container)

	router.Init(e, container)

	// starting scheduled crons
	offline.Init(container, commentService)

	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Error(err)
	}
	defer rep.Close()
}

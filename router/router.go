package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vanshajg/go-play/config"
	"github.com/vanshajg/go-play/container"
	"github.com/vanshajg/go-play/controller"

	_ "github.com/vanshajg/go-play/docs" // for echo swagger

	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo, container container.Container) {
	e.Use(middleware.Recover())

	health := controller.NewHealthController(container)
	comment := controller.NewCommentController(container)

	e.POST(controller.ApiComment, func(c echo.Context) error {
		return comment.CreateComment(c)
	})

	e.GET(controller.ApiComment, func(c echo.Context) error {
		return comment.GetComment(c)
	})

	if container.GetEnv() == config.DEV {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.GET(controller.ApiHealth, func(c echo.Context) error {
		return health.GetHealthCheck(c)
	})
}

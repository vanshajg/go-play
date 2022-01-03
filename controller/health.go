package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vanshajg/go-play/container"
)

type HealthController struct {
	container container.Container
}

func NewHealthController(container container.Container) *HealthController {
	return &HealthController{container: container}
}

func (controller HealthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}

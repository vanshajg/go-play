package core

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/vanshajg/go-play/config"
)

type Server struct {
	Echo   *echo.Echo
	config *config.Configuration
	db     *gorm.DB
	cache  redis.Client
}

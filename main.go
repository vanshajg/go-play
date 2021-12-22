package main

import (
	"github.com/vanshajg/go-play/config"
	"github.com/vanshajg/go-play/logger"
)

func main() {
	_, env := config.Load()
	logger := logger.NewLogger(env)
	logger.GetZapLogger().Infof("Loaded from config: application.%s.yaml", env)
}

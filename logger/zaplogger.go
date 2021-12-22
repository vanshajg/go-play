package logger

import "go.uber.org/zap"

func build(config *Config) (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	return logger, err
}

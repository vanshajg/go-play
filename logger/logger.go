package logger

import (
	"fmt"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ZapConfig zap.Config        `yaml:"zap_config"`
	LogRotate lumberjack.Logger `yaml:"log_rotate"`
}

type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(env string) *Logger {
	configYaml, err := ioutil.ReadFile("./zaplogger." + env + ".yaml")
	if err != nil {
		fmt.Printf("failed to read logger info %s", err)
		os.Exit(2)
	}

	var config *Config
	if err = yaml.Unmarshal(configYaml, &config); err != nil {
		fmt.Printf("failed to read logger info %s", err)
		os.Exit(2)
	}

	var zap *zap.Logger
	zap, err = build(config)

	if err != nil {
		fmt.Printf("failed to compose logger %s", err)
		os.Exit(2)
	}
	sugar := zap.Sugar()
	logger := &Logger{Zap: sugar}
	logger.Zap.Infof("Successfully configured zaplogger with config: zaplogger.%s.yaml", env)
	_ = zap.Sync()
	return logger
}

func (log *Logger) GetZapLogger() *zap.SugaredLogger {
	return log.Zap
}

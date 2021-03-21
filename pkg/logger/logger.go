package logger

import (
	"go.uber.org/zap"
)

func config(file string) zap.Config {
	configZap := zap.NewDevelopmentConfig()
	configZap.OutputPaths = []string{"stdout", file}
	configZap.DisableStacktrace = true
	configZap.ErrorOutputPaths = []string{"stderr"}
	return configZap
}

func NewZapLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "./log"}
	config.DisableStacktrace = true
	config.ErrorOutputPaths = []string{"stderr"}

	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	return config.Build()
}

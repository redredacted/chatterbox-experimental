package logger

import (
	"go.uber.org/zap"
)

func LogConfig() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"./log/log.txt"}
	
	return cfg.Build()
}

func Fatal(err error) {
	logger, _ := LogConfig()
	
	defer logger.Sync()
	logger.Fatal(err.Error())
}

func Info(msg string) {
	logger, _ := LogConfig()
	defer logger.Sync()
	
	logger.Info(msg)
}

func Debug(msg string) {
	logger, _ := LogConfig()
	defer logger.Sync()
	
	logger.Debug(msg)
}

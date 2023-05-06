package logger

import (
	"go.uber.org/zap"
)

func logConfig() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"./redis_log.txt"}
	
	return cfg.Build()
}

func FLogFatal(err error) {
	logger, _ := logConfig()
	
	defer logger.Sync()
	logger.Fatal(err.Error())
}

func FLogInfo(msg string) {
	logger, _ := logConfig()
	defer logger.Sync()
	
	logger.Info(msg)
}

func FLogDebug(msg string) {
	logger, _ := logConfig()
	defer logger.Sync()
	
	logger.Debug(msg)
}

func LogStdout() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()	
	
	return logger
}
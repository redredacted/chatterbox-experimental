package main

import "go.uber.org/zap"

func Logger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	
	return &zap.Logger{}
}
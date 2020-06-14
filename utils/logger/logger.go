package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func StartLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Unable to start the application!, Error: %s", err.Error())
	}
}

func GetLogger() *zap.Logger {
	if logger == nil {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Unable to start the application!, Error: %s", err.Error())
		}
	}
	return logger
}

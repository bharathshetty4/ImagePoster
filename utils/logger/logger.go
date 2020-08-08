package logger

import (
	"context"
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

//TODO: get req-id for context and set it for zap
func GetLogger(ctx context.Context) *zap.SugaredLogger {
	if logger == nil {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Unable to start the application!, Error: %s", err.Error())
		}
	}
	sLogger := logger.Sugar()
	sLogger = sLogger.With("LogID", "12345")
	return sLogger
}

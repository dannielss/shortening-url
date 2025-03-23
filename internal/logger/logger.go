package logger

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func NewLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error initializing zap logger: %v", err)
	}
	Logger = logger
}

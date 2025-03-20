package logger

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("Failed to create logger", err)
		return nil
	}
	return logger
}

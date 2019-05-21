package core

import (
	"github.com/hirsch88/go-micro-framework/config"
	"go.uber.org/zap"
)

func NewLogger(config *config.AppConfig) *zap.SugaredLogger {
	var log *zap.Logger
	if config.Env != "production" {
		log, _ = zap.NewDevelopment()
	} else {
		log, _ = zap.NewProduction()
	}
	return log.Sugar()
}

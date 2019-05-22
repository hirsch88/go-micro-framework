package core

import (
	"context"
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

func NewGinEngine(lifecycle fx.Lifecycle, appConfig *config.AppConfig, log *zap.SugaredLogger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(ginzap.Ginzap(log.Desugar(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(log.Desugar(), true))

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {

			if appConfig.ShowBanner {
				fmt.Println("")
				fmt.Println("Aloha, your app is ready on http://localhost:" + appConfig.Port)
				fmt.Println("To shut it down, press <CTRL> + C at any time.")
				fmt.Println("")
			}

			return engine.Run(":" + appConfig.Port)
		},
	})

	return engine
}

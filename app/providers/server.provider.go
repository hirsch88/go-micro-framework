 package providers

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

func NewServerProvider(lifecycle fx.Lifecycle, appConfig *config.AppConfig, log *zap.SugaredLogger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.New()
	server.Use(ginzap.Ginzap(log.Desugar(), time.RFC3339, true))
	server.Use(ginzap.RecoveryWithZap(log.Desugar(), true))

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {

			if appConfig.ShowBanner {
				fmt.Println("")
				fmt.Println("Aloha, your app is ready on http://localhost:" + appConfig.Port)
				fmt.Println("To shut it down, press <CTRL> + C at any time.")
				fmt.Println("")
			}

			return server.Run(":" + appConfig.Port)
		},
	})

	return server
}

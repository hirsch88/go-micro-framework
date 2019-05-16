package testUtils

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/bootstrap"
)

func BootstrapApp() *gin.Engine {
	gin.SetMode(gin.TestMode)
	app := bootstrap.App()
	return app
}

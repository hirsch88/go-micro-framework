package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/config"
)

func App(middleware ...gin.HandlerFunc) *gin.Engine {
	app := constructApp()
	app.Use(middleware...)
	registerRoutes(app)
	return app
}

func constructApp() *gin.Engine {
	return gin.New()
}

func registerRoutes(app *gin.Engine) {
	container := config.NewContainer()
	config.Routes(app.Group("/api"), container)
}

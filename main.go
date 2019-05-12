package main

import (
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/hirsch88/go-micro-framework/lib"
	"github.com/thinkerou/favicon"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()

	app.Use(gin.Recovery())
	app.Use(lib.Logger())
	app.Use(cors.Default())
	app.Use(helmet.Default())
	app.Use(helmet.NoCache())
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(favicon.New("./public/favicon.ico"))
	app.Use(static.Serve("/", static.LocalFile("./public", false)))

	container := config.NewContainer()
	config.Routes(app.Group("/api"), container)

	app.Run(":8080")

}

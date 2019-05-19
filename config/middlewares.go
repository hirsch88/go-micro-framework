package config

import (
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/lib"
	"github.com/thinkerou/favicon"
)

func Middlewares(app *gin.Engine) {
	/*
	|--------------------------------------------------------------------------
	| Global Middlewares
	|--------------------------------------------------------------------------
	|
	| Here we define the global middlewars of our application. Those will be added
	| to the gin framework.
	|
	*/
	app.Use(
		lib.Logger(),
		cors.Default(),
		gzip.Gzip(gzip.DefaultCompression),
		helmet.NoSniff(),
		helmet.DNSPrefetchControl(),
		helmet.FrameGuard(),
		helmet.SetHSTS(true),
		helmet.IENoOpen(),
		helmet.XSSFilter(),
		helmet.NoCache(),
		favicon.New("./public/favicon.ico"),
		static.Serve("/", static.LocalFile("./public", false)),
	)
}

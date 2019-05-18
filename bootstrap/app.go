package bootstrap

import (
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/hirsch88/go-micro-framework/lib"
	"github.com/hirsch88/go-micro-framework/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/thinkerou/favicon"
)

func App() *gin.Engine {

	/*
	|--------------------------------------------------------------------------
	| Create The Application
	|--------------------------------------------------------------------------
	|
	| The first thing we will do is create a new GIN Application. All the
	| middlewares and routes will be added to this instance.
	|
	*/

	app := gin.New()

	/*
	|--------------------------------------------------------------------------
	| Register Middlewares
	|--------------------------------------------------------------------------
	|
	| Next, we need to register all the global middlewares into the application.
	| All the incoming request will be passing those middlewares.
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

	/*
	|--------------------------------------------------------------------------
	| Register Routes
	|--------------------------------------------------------------------------
	|
	| Here we add our api endpoints to the application. These routes are prefixed
	| with the default value 'api'.
	|
	*/

	routes.API(app.Group("/api"), config.NewContainer())

	/*
	|--------------------------------------------------------------------------
	| Return The Application
	|--------------------------------------------------------------------------
	|
	| This script returns the application instance. The instance is given to
	| the calling script so we can separate the building of the instances
	| from the actual running of the application and sending responses.
	|
	*/

	return app
}

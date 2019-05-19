package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/hirsch88/go-micro-framework/lib"
	"github.com/hirsch88/go-micro-framework/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	if gin.Mode() != "test" {
		config.Middlewares(app)
	}

	/*
	|--------------------------------------------------------------------------
	| Register Routes & Build up the IoC Container
	|--------------------------------------------------------------------------
	|
	| Here we add our api endpoints to the application. These routes are prefixed
	| with the default value 'api'. Moreover we pass a function to the container build
	| up, which can create a new database connection.
	|
	*/

	routes.API(app.Group(config.App().Prefix), config.NewContainer(lib.DB(
		config.Database().Dialect,
		config.Database().Connection,
		config.Database().LogMode,
		config.Database().IdleConnections,
		config.Database().OpenConnections,
	)))

	/*
	|--------------------------------------------------------------------------
	| Print Application Banner to the Console
	|--------------------------------------------------------------------------
	|
	| This script returns the application instance. The instance is given to
	| the calling script so we can separate the building of the instances
	| from the actual running of the application and sending responses.
	|
	*/

	lib.PrintBanner(
		config.App().ShowBanner,
		config.App().Port,
	)

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

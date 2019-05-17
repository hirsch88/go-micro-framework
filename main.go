/**
 * Go Micro Framework
 *
 * @author   Gery Hirschfeld (hirsch88)
 */

package main

import (
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/bootstrap"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/hirsch88/go-micro-framework/lib"
	"github.com/joho/godotenv"
	"github.com/thinkerou/favicon"
	"log"
)

func main() {

	/*
	|--------------------------------------------------------------------------
	| Load the environment variables
	|--------------------------------------------------------------------------
	|
	| First we have to load the environment variables out form the .env file.
	| These variables are set and defined in the config folder.
	|
	*/

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/*
	|--------------------------------------------------------------------------
	| Turn On The Lights
	|--------------------------------------------------------------------------
	|
	| Lets bootstrap our application and turn the lights on.
	| If you wish to add any global middleware here is the place. These middlewares
	| are not used during the api testing.
	|
	*/

	gin.SetMode(gin.ReleaseMode)

	app := bootstrap.App(
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
	| Run The GIN Application
	|--------------------------------------------------------------------------
	|
	| Once we have our application, we can listen for incoming request and send
	| the associated response.
	|
	*/

	app.Run(":" + config.App().Port)

}

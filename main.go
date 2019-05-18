/**
 * Go Micro Framework
 *
 * @author   Gery Hirschfeld (hirsch88)
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/bootstrap"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/hirsch88/go-micro-framework/lib"
)

func main() {

	/*
	|--------------------------------------------------------------------------
	| LoadEnv the environment variables
	|--------------------------------------------------------------------------
	|
	| First we have to load the environment variables out form the .env file.
	| These variables are set and defined in the config folder.
	|
	*/

	lib.LoadEnv()

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
	app := bootstrap.App()

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

/**
 * Go Micro Framework
 *
 * @author   Gery Hirschfeld (hirsch88)
 */

package main

import (
	"context"
	"github.com/hirsch88/go-micro-framework/bootstrap"
	"github.com/joho/godotenv"
	"log"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
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

	app := bootstrap.App()

	/*
	|--------------------------------------------------------------------------
	| Run The Application
	|--------------------------------------------------------------------------
	|
	| Once we have our application, we can listen for incoming request and send
	| the associated response.
	|
	*/

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

}

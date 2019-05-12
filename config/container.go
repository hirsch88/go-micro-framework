package config

import "github.com/hirsch88/go-micro-framework/app/controllers"

/*
|--------------------------------------------------------------------------
| Application Dependency Injections
|--------------------------------------------------------------------------
|
| Here is where you can inject your controllers, services and repositories,
| which are use in the routes.go file to bind the controllers to the routes.
|
*/

type Container struct {
	APIController  controllers.APIController
	UserController controllers.UserController
}

func NewContainer() Container {

	/*
	   |--------------------------------------------------------------------------
	   | Register Repositories
	   |--------------------------------------------------------------------------
	*/

	/*
	   |--------------------------------------------------------------------------
	   | Register Services
	   |--------------------------------------------------------------------------
	*/

	/*
	   |--------------------------------------------------------------------------
	   | Register Controllers
	   |--------------------------------------------------------------------------
	*/
	APIController := controllers.NewAPIController()
	UserController := controllers.NewUserController()

	/*
	   |--------------------------------------------------------------------------
	   | Return the Dependency Injection Container
	   |--------------------------------------------------------------------------
	*/
	return Container{
		*APIController,
		*UserController,
	}
}

package config

import (
	"github.com/hirsch88/go-micro-framework/app/http/controllers"
	"github.com/hirsch88/go-micro-framework/app/repositories"
	"github.com/hirsch88/go-micro-framework/app/services"
)

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

func NewContainer(providers *ProvidersConfig) Container {

	/*
	|--------------------------------------------------------------------------
	| Register Repositories
	|--------------------------------------------------------------------------
	*/
	userRepository := repositories.NewUserRepository(providers.Database)

	/*
	|--------------------------------------------------------------------------
	| Register Services
	|--------------------------------------------------------------------------
	*/
	userService := services.NewUserService(userRepository, providers.Mail)

	/*
	|--------------------------------------------------------------------------
	| Register Controllers
	|--------------------------------------------------------------------------
	*/
	apiController := controllers.NewAPIController()
	userController := controllers.NewUserController(userService)

	/*
	|--------------------------------------------------------------------------
	| Return the Dependency Injection Container
	|--------------------------------------------------------------------------
	*/
	return Container{
		*apiController,
		*userController,
	}
}

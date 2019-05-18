package config

import (
	"github.com/hirsch88/go-micro-framework/app/http/controllers"
	"github.com/hirsch88/go-micro-framework/app/repositories"
	"github.com/hirsch88/go-micro-framework/app/services"
	"github.com/jinzhu/gorm"
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

func NewContainer(db func() *gorm.DB) Container {

	/*
	|--------------------------------------------------------------------------
	| Register Repositories
	|--------------------------------------------------------------------------
	*/
	userRepository := repositories.NewUserRepository(db)

	/*
	|--------------------------------------------------------------------------
	| Register Services
	|--------------------------------------------------------------------------
	*/
	userService := services.NewUserService(userRepository)

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

package bootstrap

import (
	"github.com/hirsch88/go-micro-framework/app/http/controllers"
	"github.com/hirsch88/go-micro-framework/app/http/middlewares"
	"github.com/hirsch88/go-micro-framework/app/providers"
	"github.com/hirsch88/go-micro-framework/app/repositories"
	"github.com/hirsch88/go-micro-framework/app/services"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/hirsch88/go-micro-framework/routes"
	"go.uber.org/fx"
)

func App() *fx.App {
	/*
	|--------------------------------------------------------------------------
	| Create The Application
	|--------------------------------------------------------------------------
	|
	| The first thing we will do is create a new Application. Then we have to
	| provide all our application items like controllers, services and stuff.
	| The Goal of this file is to separate the building of the instances
	| from the actual running of the application and sending responses.
	|
	*/
	return fx.New(
		/*
		|--------------------------------------------------------------------------
		| Application Dependency Injections
		|--------------------------------------------------------------------------
		|
		| Here is where you can inject your controllers, services and repositories,
		| which are use in the routes.go file to bind the controllers to the routes.
		|
		*/

		// Configurations (./config)
		fx.Provide(config.NewAppConfig),
		fx.Provide(config.NewDatabaseConfig),
		fx.Provide(config.NewMailConfig),

		// Providers (./app/providers)
		fx.Provide(providers.NewLoggerProvider),
		fx.Provide(providers.NewTemplateProvider),
		fx.Provide(providers.NewDatabaseProvider),
		fx.Provide(providers.NewSMTPMailProvider),
		fx.Provide(providers.NewMailProvider),
		fx.Provide(providers.NewServerProvider),

		// Repositores (./app/repositores)
		fx.Provide(repositories.NewUserRepository),

		// Services (./app/services)
		fx.Provide(services.NewUserService),

		// Middlewares (./app/middlewares)
		fx.Provide(middlewares.NewLogMiddleware),

		// Controllers (./app/controllers)
		fx.Provide(controllers.NewAPIController),
		fx.Provide(controllers.NewUserController),

		/*
		|--------------------------------------------------------------------------
		| Register Middlewares
		|--------------------------------------------------------------------------
		|
		| Next, we need to register all the global middlewares into the application.
		| All the incoming request will be passing those middlewares.
		|
		*/

		fx.Invoke(middlewares.GlobalMiddlewares),

		/*
		|--------------------------------------------------------------------------
		| Invoke Register Routes
		|--------------------------------------------------------------------------
		|
		| Here we add our api endpoints to the application. These routes are prefixed
		| with the default value 'api'. Moreover we pass a function to the container build
		| up, which can create a new database connection.
		|
		*/

		fx.Invoke(routes.APIRoutes),

	)
}

package routes

import (
	"github.com/cnjack/throttle"
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/http/controllers"
	"github.com/hirsch88/go-micro-framework/config"
	"go.uber.org/dig"
	"time"
)

func APIRoutes(c RouterContext) {
	api := c.Router.Group(c.AppConfig.Prefix)
	{

		/*
		|--------------------------------------------------------------------------
		| Application Routes
		|--------------------------------------------------------------------------
		|
		| Here is where you can register all of the routes for an application.
		| It is a breeze. Simply tell gin the URIs it should respond to
		| and give it the Closure to call when that URI is requested.
		|
		*/

		api.GET("/ping", c.APIController.Ping)

		users := api.Group("/users")
		{
			users.POST(
				"",
				throttle.Policy(&throttle.Quota{
					Limit:  1,
					Within: time.Hour,
				}),
				c.UserController.Create)
		}

	}
}

type RouterContext struct {
	dig.In

	Router    *gin.Engine
	AppConfig *config.AppConfig

	APIController *controllers.APIController
	UserController *controllers.UserController
}

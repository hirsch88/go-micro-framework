package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hirsch88/go-micro-framework/app/http/middlewares"
	"github.com/hirsch88/go-micro-framework/config"
)

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

func API(router *gin.RouterGroup, container config.Container){

	router.Use(middlewares.Logger())

	router.GET("/ping", container.APIController.Ping)

	users := router.Group("/users")
	{
		users.POST("", container.UserController.Create)
	}

}

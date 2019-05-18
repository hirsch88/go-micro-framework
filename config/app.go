package config

import (
	"github.com/hirsch88/go-micro-framework/lib"
)

func App() *AppConfig {
	return &AppConfig{
		/*
		|--------------------------------------------------------------------------
		| Application Name
		|--------------------------------------------------------------------------
		|
		| This value is the name of your application. This value is used when the
		| framework needs to place the application's name in a notification or
		| any other location as required by the application or its packages.
		|
		*/

		Name: lib.Env("APP_NAME", "go-micro-framework"),

		/*
		|--------------------------------------------------------------------------
		| Application Port
		|--------------------------------------------------------------------------
		|
		| This value define on witch port the application is available. Default is
		| the standard port 8080
		|
		*/

		Port: lib.Env("APP_PORT", "go-micro-framework"),

		/*
		|--------------------------------------------------------------------------
		| Application Name
		|--------------------------------------------------------------------------
		|
		| This value is the name of your application. This value is used when the
		| framework needs to place the application's name in a notification or
		| any other location as required by the application or its packages.
		|
		*/

		Prefix: lib.Env("APP_PREFIX", "/api"),

		/*
		|--------------------------------------------------------------------------
		| Application Banner
		|--------------------------------------------------------------------------
		|
		| If you wish to see a nice log statement when the application has been
		| bootstrap then activate this config.
		|
		*/

		ShowBanner: lib.EnvBool("APP_SHOW_BANNER", true),
	}
}

type AppConfig struct {
	Name       string
	Port       string
	Prefix     string
	ShowBanner bool
}

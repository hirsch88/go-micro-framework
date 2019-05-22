package config

func NewAppConfig() *AppConfig {
	return &AppConfig{
		/*
		|--------------------------------------------------------------------------
		| Environment Name
		|--------------------------------------------------------------------------
		|
		| With the help of the environment name we can implement different behaviour for
		| our environments. For example for local development we like to have another
		| logger encoding with colors, but in production we have the JSON encoding.
		|
		*/

		Env: Env("APP_ENV", "production"),

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

		Name: Env("APP_NAME", "go-micro-framework"),

		/*
		|--------------------------------------------------------------------------
		| Application Port
		|--------------------------------------------------------------------------
		|
		| This value define on witch port the application is available. Default is
		| the standard port 8080
		|
		*/

		Port: Env("APP_PORT", "go-micro-framework"),

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

		Prefix: Env("APP_PREFIX", "/api"),

		/*
		|--------------------------------------------------------------------------
		| Application Banner
		|--------------------------------------------------------------------------
		|
		| If you wish to see a nice log statement when the application has been
		| bootstrap then activate this config.
		|
		*/

		ShowBanner: EnvBool("APP_SHOW_BANNER", true),
	}
}

type AppConfig struct {
	Env        string
	Name       string
	Port       string
	Prefix     string
	ShowBanner bool
}

package config

import (
	"github.com/hirsch88/go-micro-framework/app/providers"
)

func Providers() *ProvidersConfig {
	return &ProvidersConfig{
		/*
		|--------------------------------------------------------------------------
		| Database Provider
		|--------------------------------------------------------------------------
		*/
		Database: providers.NewDatabaseProvider(
			Database().Dialect,
			Database().Connection,
			Database().LogMode,
			Database().IdleConnections,
			Database().OpenConnections,
		),

		/*
		|--------------------------------------------------------------------------
		| Mail Provider
		|--------------------------------------------------------------------------
		*/
		Mail: providers.NewMailProvider(
			Mail().Host,
			Mail().Port,
			Mail().From,
			Mail().Password,
		),
	}
}

type ProvidersConfig struct {
	Database providers.DatabaseProvider
	Mail     providers.MailProvider
}

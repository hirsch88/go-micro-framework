package config

import (
	"github.com/hirsch88/go-micro-framework/app/providers"
	"github.com/jinzhu/gorm"
)

func Providers() *ProvidersConfig {
	return &ProvidersConfig{
		/*
		|--------------------------------------------------------------------------
		| Database Provider
		|--------------------------------------------------------------------------
		*/
		Database: providers.DatabaseProvider(
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
	Database func() *gorm.DB
	Mail     providers.MailProvider
}

package config

import "github.com/hirsch88/go-micro-framework/lib"

func Database() *DatabaseConfig {
	return &DatabaseConfig{
		/*
		|--------------------------------------------------------------------------
		| Database Connection
		|--------------------------------------------------------------------------
		|
		| Here you may specify which of the database connections you wish
		| to use as your default connection for all database work.
		|
		*/

		Dialect:         lib.Env("DB_DIALECT", "mysql"),
		Connection:      lib.Env("DB_CONNECTION", "root:root@/go-micro-framework?charset=utf8mb4&parseTime=True&loc=Local"),
		LogMode:         lib.EnvBool("DB_LOG_MODE", false),
		IdleConnections: lib.EnvNumber("DB_IDLE_CONNECTIONS", 10),
		OpenConnections: lib.EnvNumber("DB_OPEN_CONNECTIONS", 100),
	}
}

type DatabaseConfig struct {
	Dialect         string
	Connection      string
	LogMode         bool
	IdleConnections int
	OpenConnections int
}

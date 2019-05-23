package config

func NewDatabaseConfig() *DatabaseConfig {
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

		Dialect:         Env("DB_DIALECT", "mysql"),
		Connection:      Env("DB_CONNECTION", "root:root@/go-micro-framework?charset=utf8mb4&parseTime=True&loc=Local"),
		LogMode:         EnvBool("DB_LOG_MODE", false),
		IdleConnections: EnvInt("DB_IDLE_CONNECTIONS", 10),
		OpenConnections: EnvInt("DB_OPEN_CONNECTIONS", 100),
	}
}

type DatabaseConfig struct {
	Dialect         string
	Connection      string
	LogMode         bool
	IdleConnections int
	OpenConnections int
}

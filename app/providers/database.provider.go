package providers

import (
	"fmt"
	"github.com/hirsch88/go-micro-framework/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wantedly/gorm-zap"
	"go.uber.org/zap"
)

type databaseProvider struct {
	config *config.DatabaseConfig
	log *zap.SugaredLogger
}

func (p *databaseProvider) Connect() *gorm.DB {
	db, err := gorm.Open(p.config.Dialect, p.config.Connection)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(p.config.LogMode)
	db.SetLogger(gormzap.New(p.log.Desugar()))
	db.DB().SetMaxIdleConns(p.config.IdleConnections)
	db.DB().SetMaxOpenConns(p.config.OpenConnections)

	return db
}

func (p *databaseProvider) Migrate() {
	db := p.Connect()
	defer db.Close()

	// Here you might migrate your database with the models
	//db.AutoMigrate(&models.User{})
}

type DatabaseProvider interface {
	Connect() *gorm.DB
	Migrate()
}

func NewDatabaseProvider(config *config.DatabaseConfig, log *zap.SugaredLogger) DatabaseProvider {
	return &databaseProvider{config, log}
}

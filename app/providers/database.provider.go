package providers

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type databaseProvider struct {
	dialect         string
	connection      string
	logMode         bool
	idleConnections int
	openConnections int
}

func (p *databaseProvider) Connect() *gorm.DB {
	db, err := gorm.Open(p.dialect, p.connection)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(p.logMode)
	db.DB().SetMaxIdleConns(p.idleConnections)
	db.DB().SetMaxOpenConns(p.openConnections)

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

func NewDatabaseProvider(dialect string, connection string, logMode bool, idleConnections int, openConnections int) DatabaseProvider {
	return &databaseProvider{dialect, connection, logMode, idleConnections, openConnections}
}

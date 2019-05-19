package providers

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func DatabaseProvider(dialect string, connection string, logMode bool, idleConnections int, openConnections int) func() *gorm.DB {
	return func() *gorm.DB {
		db, err := gorm.Open(dialect, connection)

		if err != nil {
			fmt.Println(err.Error())
			panic("failed to connect database")
		}

		if err = db.DB().Ping(); err != nil {
			panic(err)
		}

		db.LogMode(logMode)
		db.DB().SetMaxIdleConns(idleConnections)
		db.DB().SetMaxOpenConns(openConnections)

		// Here you might migrate your database with the models
		//db.AutoMigrate(&models.User{})

		return db
	}
}

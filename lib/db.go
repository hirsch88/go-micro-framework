package lib

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/go-micro-framework?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	//db.AutoMigrate(&models.User{})

	return db
}

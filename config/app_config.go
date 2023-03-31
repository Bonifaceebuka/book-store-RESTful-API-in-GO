package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	databaseConnection *gorm.DB
)

func connect() {
	dbConnection, err := gorm.Open("mysql", "root:/book_store_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Unable to open database")
	}
	databaseConnection = dbConnection
}

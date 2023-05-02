package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {

	dsn := "root:@tcp(localhost:3307)/go_db?charset=utf8&parseTime=True&loc=Local"
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}

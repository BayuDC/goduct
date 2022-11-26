package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("bayudc:bay@tcp(127.0.0.1:3306)/goduct"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	Db = database
}

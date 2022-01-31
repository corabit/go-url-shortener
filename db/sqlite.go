package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var db *gorm.DB

func GetSqliteDatabase() *gorm.DB {
	if db != nil {
		return db
	}
	db, err := gorm.Open(sqlite.Open("url.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

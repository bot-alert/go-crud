package database

import (
	"GO/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func StartDatabase() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if nil != err {
		panic(err)
	}
	Database = db
	_ = Database.AutoMigrate(&model.User{})
	_ = Database.AutoMigrate(&model.Post{})
}

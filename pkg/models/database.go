package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

//The main db reference
var DB *gorm.DB

//Connects to the database
func Connect() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&UserTable{})
	db.AutoMigrate(&UserTableField{})

	DB = db
}

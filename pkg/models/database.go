package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the main database reference
var DB *gorm.DB

// Connect establishes a connection to the database
func Connect() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&UserTable{})
	db.AutoMigrate(&UserTableField{})

	DB = db
}

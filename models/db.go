package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&User{}, &Item{}, &Cart{}, &Order{})
	DB = database
	return database
}

func CloseDatabase() {
	dbSQL, _ := DB.DB()
	dbSQL.Close()
}

package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	connStr := "host=localhost user=postgres dbname=go_blog port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	err = db.AutoMigrate(&Post{})
	if err != nil {
		return
	}
	DB = db
}

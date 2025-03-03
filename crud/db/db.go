package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

func InitDB() {
	dsn := "host=localhost user=aritraganai password=ARga12@@ dbname=postgres port=5433 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
     log.Println("Database connected")
	// Auto migrate User table
	DB.AutoMigrate(&User{})
}
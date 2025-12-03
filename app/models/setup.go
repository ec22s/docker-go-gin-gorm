package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDataBase() {
	if godotenv.Load() != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASS"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_NAME"),
  )
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	// Execute DB migration
	DB.AutoMigrate(&User{})
}

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/webstraservices/gateway/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load environment files - Err: %s", err)
	}

	connectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to open database instance")
	}

	db.AutoMigrate(models.User{})
}
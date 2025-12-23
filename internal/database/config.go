package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func NewConnection() {
	var err error
	loadEnv()
	dbuser := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC ", dbuser, dbpassword, dbname, dbport)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("DB CONNECTED")
}

func loadEnv() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}

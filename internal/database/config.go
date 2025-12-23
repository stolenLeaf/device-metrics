package database

import (
	"fmt"
	"github.com/stolenleaf/device-metrics/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

type DBconfig struct {
	Dbuser     string
	Dbpassword string
	Dbname     string
	Dbport     string
	Dsn        string
}

func (d *DBconfig) loadDBConfig() {
	utils.LoadEnv()
	d.Dbuser = os.Getenv("DB_USERNAME")
	d.Dbpassword = os.Getenv("DB_PASSWORD")
	d.Dbname = os.Getenv("DB_NAME")
	d.Dbport = os.Getenv("DB_PORT")
	d.Dsn = fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC ", d.Dbuser, d.Dbpassword, d.Dbname, d.Dbport)
}

func NewConnection() {
	var err error
	dbconfig := DBconfig{}
	dbconfig.loadDBConfig()
	DB, err = gorm.Open(postgres.Open(dbconfig.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("DB CONNECTED")
}

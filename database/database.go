package database

import (
	"sample/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var DNS = "host=localhost user=postgres password=sparkman13 dbname=testingpanics port=5432 sslmode=disable"

func Migration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{
		
	})
	if err != nil {
		panic("cant connect to the db")
	}

	DB.AutoMigrate(model.User{})
}

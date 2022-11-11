package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Migration() {

	DSN := ("host=34.87.184.130	user=janus_user password=j@Nu5.6@t3_p@55.2O22@ dbname=mfs port=5432 sslmode=disable")
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the databse")

}

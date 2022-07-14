package database

import (
	"fmt"
	"log"

	"github.com/pankaj-nupay/jwt-in-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
		panic("Can not connect to DB")
	}
	fmt.Println("Connected to Database.")
}

func Migrate() {
	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Migration Completed.")
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pankaj-nupay/jwt-in-go/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s  password=%s  dbname=%s  sslmode=%s ",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"),
	)
	database.Connect(dsn)
	database.Migrate()
}

package main

import (
	"log"

	"github.com/9thDuck/ecommerce-api.git/pkg/db"
	"github.com/9thDuck/ecommerce-api.git/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error: err loading .env. Check if it exists")
	}

	POSTGRES_DB_DSN := utils.GetEnvOrThrow("POSTGRES_DB_DSN")

	db.Connect(POSTGRES_DB_DSN)
}

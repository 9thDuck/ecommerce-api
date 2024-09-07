package main

import (
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/users"
	"github.com/9thDuck/ecommerce-api.git/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.LogFatalCustomError("err loading .env", err)

	db.Setup()

	app := fiber.New()
	app.Listen(":3000")
}

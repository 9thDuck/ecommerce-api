package main

import (
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/entities"
	"github.com/9thDuck/ecommerce-api.git/users"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	common.LogFatalCustomError("err loading .env", err)

	// Setup app config. APP_CONFIG is a variable, after this, if all goes well it will be populated with env vars validated to some extent
	common.APP_CONFIG.ValidateAndSetup()

	// Supply all the entities
	entitySlice := []any{&entities.User{}}
	db.SetupDbInstance(entitySlice)

	app := fiber.New()
	users.RegisterRoutes(app)
	app.Listen(":3000")
}

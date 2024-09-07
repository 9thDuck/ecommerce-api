package users

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(server *fiber.App) {
	server.Post("/users/signup", signup)
	server.Post("/users/login", login)
}

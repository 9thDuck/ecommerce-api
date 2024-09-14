package users

import (
	"github.com/9thDuck/ecommerce-api.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(server *fiber.App) {
	server.Post("/users/signup", signup)
	server.Post("/users/login", login)
	server.Get("/users/me", middleware.VerifyToken, getMe)
	server.Get("/users/:id", middleware.VerifyToken, middleware.RoleGuard(middleware.ADMIN), getUserById)
	server.Delete("/users/logout", middleware.VerifyToken, logout)
	server.Delete("/users/logout-all", middleware.VerifyToken, logoutOfAllDevices)
	server.Patch("/users/:id", middleware.VerifyToken, middleware.RoleGuard(middleware.ADMIN), banNonAdmin)
}

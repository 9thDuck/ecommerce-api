package users

import (
	"github.com/9thDuck/ecommerce-api.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/users/signup", signup)
	app.Post("/users/login", login)
	app.Get("/users/me", middleware.VerifyToken, getMe)
	app.Get("/users/:id", middleware.VerifyToken, middleware.RoleGuard(middleware.ADMIN), getUserById)
	app.Delete("/users/logout", middleware.VerifyToken, logout)
	app.Delete("/users/logout-all", middleware.VerifyToken, logoutOfAllDevices)
	app.Patch("/users/:id", middleware.VerifyToken, middleware.RoleGuard(middleware.ADMIN), banNonAdmin)
}

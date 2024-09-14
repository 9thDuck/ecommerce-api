package categories

import (
	"github.com/9thDuck/ecommerce-api.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/categories", middleware.VerifyToken, getCategories)
	app.Get("/categories/:id", middleware.VerifyToken, getCategory)
	app.Post("/categories", middleware.VerifyToken, createCategory)
	app.Put("/categories/:id", middleware.VerifyToken, updateCategory)
	app.Delete("/categories/:id", middleware.VerifyToken, deleteCategory)
}

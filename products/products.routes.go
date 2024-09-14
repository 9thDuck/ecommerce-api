package products

import (
	"github.com/9thDuck/ecommerce-api.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/products", middleware.VerifyToken, getProducts)
	app.Get("/products/:id", middleware.VerifyToken, getProduct)
	app.Get("/products/category/:id", middleware.VerifyToken, getProductsByCategoryID)
	app.Get("/products/search", middleware.VerifyToken, getProductsBySearch)
	app.Post("/products", middleware.VerifyToken, createProduct)
	app.Put("/products/:id", middleware.VerifyToken, updateProduct)
	app.Delete("/products/:id", middleware.VerifyToken, deleteProduct)
}
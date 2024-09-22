package products

import (
	"github.com/9thDuck/ecommerce-api.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/products", getProducts)
	app.Get("/products/:id", getProduct)
	app.Get("/products/category/:id", getProductsByCategoryID)
	app.Get("/products/search", getProductsBySearch)
	app.Post("/products", middleware.VerifyToken, createProduct)
	app.Put("/products/:id", middleware.VerifyToken, updateProduct)
	app.Delete("/products/:id", middleware.VerifyToken, deleteProduct)
}

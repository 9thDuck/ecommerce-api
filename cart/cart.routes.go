package cart

import (
	"github.com/9thDuck/ecommerce-api.git/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterCartRoutes(app fiber.Router) {
	app.Post("/cart", middleware.VerifyToken, addToCart)
	app.Get("/cart", middleware.VerifyToken, getCart)
	app.Put("/cart/:id", middleware.VerifyToken, removeFromCart)
	app.Delete("/cart/:id", middleware.VerifyToken, deleteProductFromCart)
	app.Delete("/cart", middleware.VerifyToken, clearCart)
}

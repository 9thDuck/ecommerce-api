package checkout

import "github.com/gofiber/fiber/v2"

func RegisterCheckoutRoutes(app fiber.Router) {
	app.Post("/checkout", checkout)
	app.Delete("/checkout/:id", cancelCheckout)
}

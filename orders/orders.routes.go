package orders

import "github.com/gofiber/fiber/v2"

func RegisterOrdersRoutes(app *fiber.App) {
	app.Get("/orders/:id", getOrderById)
	app.Put("/orders/:id", cancelOrderById)
}

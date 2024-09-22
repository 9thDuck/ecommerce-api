package orders

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func cancelOrderById(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uuid.UUID)
	orderID := ctx.Params("id")
	parsedOrderID, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorCancelOrderResponse())
	}
	err = CancelOrder(nil, &Order{ID: uint(parsedOrderID), UserID: userID})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCancelOrderResponse())
	}
	return ctx.Status(fiber.StatusOK).JSON(successCancelOrderResponse())
}

func getOrderById(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uuid.UUID)
	orderID := ctx.Params("id")
	parsedOrderID, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorCancelOrderResponse())
	}
	order := Order{ID: uint(parsedOrderID), UserID: userID}
	if err := GetOrder(nil, &order); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCancelOrderResponse())
	}
	return ctx.Status(fiber.StatusOK).JSON(successGetOrderResponse(&order))
}

// func updateOrder(ctx *fiber.Ctx) error {
// 	orderID := ctx.Params("id")
// 	parsedOrderID, err := strconv.ParseUint(orderID, 10, 64)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(errorCancelOrderResponse())
// 	}
// 	order := Order{ID: uint(parsedOrderID)}
// 	if err := Get(nil, &order); err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCancelOrderResponse())
// 	}

// }

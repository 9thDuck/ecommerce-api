package checkout

import (
	"strconv"

	"github.com/9thDuck/ecommerce-api.git/orders"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func checkout(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user").(map[string]interface{})["id"].(string)
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCheckoutResponse())
	}

	orderID, err := processCheckout(ctx.Context(), parsedUserID)
	if err != nil {
		if err.Error() == "cart is empty" {
			return ctx.Status(fiber.StatusBadRequest).JSON(errorEmptyCartResponse())
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCheckoutResponse())
	}

	return ctx.JSON(successCheckoutResponse(orderID))
}

func cancelCheckout(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user").(map[string]interface{})["id"].(string)
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCheckoutResponse())
	}

	orderID := ctx.Params("order_id")
	if orderID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorCheckoutBadRequestResponse())
	}

	parsedOrderID, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorCheckoutBadRequestResponse())
	}

	order := &orders.Order{ID: uint(parsedOrderID)}

	err = orders.GetOrder(nil, order)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCheckoutResponse())
	}

	if order.UserID != parsedUserID {
		return ctx.Status(fiber.StatusForbidden).JSON(errorCheckoutForbiddenResponse())
	}

	if order.Status != orders.OrderStatusPending {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorCheckoutBadRequestResponse())
	}

	err = orders.CancelOrder(nil, order)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorCancelCheckoutResponse())
	}

	order.Status = orders.OrderStatusCancelled

	return ctx.JSON(successCancelCheckoutResponse(order))
}

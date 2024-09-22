package cart

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var validate = validator.New()

func addToCart(c *fiber.Ctx) error {
	var input struct {
		ProductID uint `json:"product_id" validate:"required,number"`
		Quantity  int  `json:"quantity" validate:"required,number,min=1"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorAddToCartResponse())
	}

	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorAddToCartResponse())
	}

	userID := c.Locals("user").(map[string]interface{})["id"].(string)
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorAddToCartResponse())
	}

	err = addToCartOrIncrement(parsedUserID, input.ProductID, input.Quantity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorAddToCartResponse())
	}

	return c.JSON(successAddToCartResponse())
}

func getCart(c *fiber.Ctx) error {
	userID := c.Locals("user").(map[string]interface{})["id"].(string)
	userId, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorGetCartResponse())
	}

	cart, err := getCartByUserID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorGetCartResponse())
	}

	return c.JSON(successGetCartResponse(cart))
}

func removeFromCart(ctx *fiber.Ctx) error {
	ID := ctx.Locals("user").(map[string]interface{})["id"].(string)
	userID, err := uuid.Parse(ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorRemoveFromCartResponse())
	}

	productID := ctx.Params("id")
	productId, err := strconv.Atoi(productID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorRemoveFromCartResponse())
	}

	err = remove(userID, uint(productId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorRemoveFromCartResponse())
	}

	return ctx.JSON(successRemoveFromCartResponse())
}

func deleteProductFromCart(ctx *fiber.Ctx) error {
	ID := ctx.Locals("user").(map[string]interface{})["id"].(string)
	userID, err := uuid.Parse(ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorDeleteProductFromCartResponse())
	}

	productID := ctx.Params("id")
	productId, err := strconv.Atoi(productID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorDeleteProductFromCartResponse())
	}

	err = deleteCartItem(userID, uint(productId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorDeleteProductFromCartResponse())
	}

	return ctx.JSON(successRemoveFromCartResponse())
}

func clearCart(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user").(map[string]interface{})["id"].(string)
	userId, err := uuid.Parse(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorClearCartResponse())
	}

	err = DeleteAllFromCart(nil, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorClearCartResponse())
	}

	return ctx.JSON(successClearCartResponse())
}

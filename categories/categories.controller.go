package categories

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var validate = validator.New()

func getCategories(c *fiber.Ctx) error {
	categories, err := getAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successGetCategoriesResponse(categories))
}

func createCategory(c *fiber.Ctx) error {
	category := new(Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	user := c.Locals("user").(jwt.MapClaims)
	userID, err := uuid.Parse(user["id"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Invalid user ID"))
	}

	category.CreatedBy = userID

	if err := validate.Struct(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	if err := create(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successCreateResponse(category))
}

func updateCategory(c *fiber.Ctx) error {
	category := new(Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	if err := validate.Struct(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	if err := update(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successUpdateResponse(category))
}

func deleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Category ID is required"))
	}

	categoryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Invalid category ID"))
	}

	category := &Category{ID: uint(categoryID)}
	if err := delete(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successDeleteResponse())
}

func getCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Category ID is required"))
	}

	categoryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Invalid category ID"))
	}

	category := &Category{ID: uint(categoryID)}
	if err := get(category); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(notFoundResponse())
	}
	return c.JSON(successGetCategoryResponse(category))
}
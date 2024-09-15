package products

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var validate = validator.New()

func createProduct(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	if err := validate.Struct(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	// Set CreatedBy from the authenticated user
	userID := c.Locals("user").(map[string]interface{})["id"].(string)
	product.CreatedBy = uuid.MustParse(userID)

	if err := create(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successCreateResponse(product))
}

func updateProduct(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	if err := validate.Struct(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse(err.Error()))
	}

	// Set CreatedBy from the authenticated user
	userID := c.Locals("user").(map[string]interface{})["id"].(string)
	product.CreatedBy = uuid.MustParse(userID)

	if err := update(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successUpdateResponse(product))
}

func deleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Product ID is required"))
	}

	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Invalid product ID"))
	}

	product := &Product{ID: uint(productID)}
	if err := delete(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successDeleteResponse())
}

func getProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Product ID is required"))
	}

	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Invalid product ID"))
	}

	product := &Product{ID: uint(productID)}
	if err := get(product); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(notFoundResponse())
	}
	return c.JSON(successGetProductResponse(product))
}

func getProductsByCategoryID(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	if categoryID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Category ID is required"))
	}

	categoryIDUint, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse("Invalid category ID"))
	}

	products, err := getProductsByOptions(&Product{CategoryID: uint(categoryIDUint)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successGetProductsResponse(products))
}

func getProductsBySearch(c *fiber.Ctx) error {
	// Parse query parameters
	name := c.Query("name")
	description := c.Query("description")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	categoryID := c.Query("category_id")
	inStock := c.Query("in_stock")

	criteria := make(map[string]interface{})

	if name != "" {
		criteria["name"] = name
	}
	if description != "" {
		criteria["description"] = description
	}
	if minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
			criteria["price_gte"] = price
		}
	}
	if maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			criteria["price_lte"] = price
		}
	}
	if categoryID != "" {
		if id, err := strconv.ParseUint(categoryID, 10, 64); err == nil {
			criteria["category_id"] = uint(id)
		}
	}
	if inStock != "" {
		if stock, err := strconv.ParseBool(inStock); err == nil {
			criteria["in_stock"] = stock
		}
	}

	products, err := getProductsBySearchCriteria(&criteria)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}

	return c.JSON(successSearchProductsResponse(products))
}

func getProducts(c *fiber.Ctx) error {
	products, err := getAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(internalServerErrorResponse())
	}
	return c.JSON(successGetProductsResponse(products))
}

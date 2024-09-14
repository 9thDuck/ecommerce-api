package common

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(message string, data interface{}) fiber.Map {
	return fiber.Map{
		"message": message,
		"data":    data,
	}
}

func FailedResponse(message string) fiber.Map {
	return fiber.Map{
		"message": message,
	}
}

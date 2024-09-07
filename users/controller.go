package users

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func signup(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Body())
	ctx.Status(http.StatusOK)
	return ctx.JSON(ctx.Body())
}

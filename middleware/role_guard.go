package middleware

import (
	"github.com/9thDuck/ecommerce-api.git/auth"
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/gofiber/fiber/v2"
)

const (
	ADMIN  Role = 0
	SELLER Role = 1
	BUYER  Role = 2
)

type Role int

func (r Role) IsAdmin() bool {
	return r == ADMIN
}

func (r Role) IsSeller() bool {
	return r == SELLER
}

func (r Role) IsBuyer() bool {
	return r == BUYER
}

func RoleGuard(role Role) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(auth.TokenClaims)
		userRole := Role(user.Role)

		if userRole > role {
			ctx.Status(fiber.StatusForbidden)
			return ctx.JSON(common.FailedResponse("Forbidden"))
		}

		return ctx.Next()
	}
}

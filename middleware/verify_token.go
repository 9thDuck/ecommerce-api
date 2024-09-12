package middleware

import (
	"net/http"
	"time"

	"github.com/9thDuck/ecommerce-api.git/auth"
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func VerifyToken(ctx *fiber.Ctx) error {
	accessToken := ctx.Cookies("access_token")
	refreshToken := ctx.Cookies("refresh_token")

	// Try to parse the access token
	claims, err := auth.ParseToken(accessToken)
	if err != nil {
		// Access token is invalid, try to use refresh token
		claims, err = auth.ParseToken(refreshToken)
		if err != nil {
			// Both tokens are invalid
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(common.Response{Message: "Unauthorized", Data: nil})
		}

		// Refresh token is valid, generate new access token
		id, _ := uuid.Parse(claims["id"].(string))
		role := int(claims["role"].(float64))
		newClaims := auth.TokenClaims{ID: id, Role: role}
		newAccessToken, _, err := auth.GenerateToken(newClaims)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(common.Response{Message: "Error refreshing token", Data: nil})
		}

		// Set new access token in cookie
		accessTokenCookie := common.MakeCookie("access_token", newAccessToken, time.Now().Add(common.APP_CONFIG.GetExpiryAccessTokenDurationInMinutes()), false, "/")
		ctx.Cookie(accessTokenCookie)

		// Update claims with new access token claims
		claims, _ = auth.ParseToken(newAccessToken)
	}

	ctx.Locals("user", claims)
	ctx.Next()
	return nil
}
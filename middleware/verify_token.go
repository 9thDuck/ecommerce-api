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
		// Refresh token is valid

		userId, _ := uuid.Parse(claims["id"].(string))

		userSession, err := getUserSession(refreshToken, userId)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(common.Response{Message: "Unauthorized", Data: nil})
		}
		//getting the existing user session to see if the user has been logged out of all devices. If that is the case, then all the refresh tokens given to user will be invalid, since we use user id to identify all the active sessions and logout all of them.
		if userSession.Session.LoggedOut {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(common.Response{Message: "Unauthorized", Data: nil})
		}
		//if the user is banned, then we don't want to issue new tokens
		if userSession.User.Banned {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(common.Response{Message: "Your account has been banned. If you believe this is a mistake, please contact support.", Data: nil})
		}

		role := int(claims["role"].(float64))
		newClaims := auth.TokenClaims{ID: userId, Role: role}
		newAccessToken, _, err := auth.GenerateToken(newClaims)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(common.Response{Message: "Error refreshing token", Data: nil})
		}

		// Set new access token in cookie
		accessTokenCookie := common.MakeCookie("access_token", newAccessToken, time.Now().Add(common.APP_CONFIG.GetExpiryAccessTokenDurationInMinutes()), false, "/")
		ctx.Cookie(accessTokenCookie)
	}

	ctx.Locals("user", claims)
	ctx.Next()
	return nil
}

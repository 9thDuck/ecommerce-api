package users

import (
	"net/http"
	"time"

	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var validate = validator.New()

func signup(ctx *fiber.Ctx) (err error) {
	userInput := createUserInput{}
	if err = ctx.BodyParser(&userInput); err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse("cannot parse the form details, please try again later"))
	}

	if err = validate.Struct(&userInput); err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse(err.Error()))
	}
	user := New(userInput.Username, userInput.Email, userInput.Password, userInput.Role)

	err = createUser(user)

	if err != nil {
		common.LogCustomError("failed to create user", err)
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse(err.Error()))
	}

	ctx.Status(http.StatusCreated)
	return ctx.JSON(successResponse("Successfully signed up", user))
}

func login(ctx *fiber.Ctx) (err error) {
	userInput := loginUserInput{}

	if err = ctx.BodyParser(&userInput); err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse("cannot parse the form details, please check the credentials and try again later"))
	}

	if err = validate.Struct(&userInput); err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse(err.Error()))
	}

	user := New("", userInput.Email, userInput.Password, 0)
	accessToken, refreshToken, err := loginUser(user)
	if err != nil {
		common.LogCustomError("login failed", err)
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse("bad credentials, please check the credentials and try again later"))
	}

	accessTokenCookie := common.MakeCookie("access_token", accessToken, time.Now().Add(common.APP_CONFIG.GetExpiryAccessTokenDurationInMinutes()), false, "/")
	ctx.Cookie(accessTokenCookie)
	refreshTokenCookie := common.MakeCookie("refresh_token", refreshToken, time.Now().Add(common.APP_CONFIG.GetExpiryRefreshTokenDurationInHours()), false, "/")
	ctx.Cookie(refreshTokenCookie)

	ctx.Status(http.StatusOK)
	return ctx.JSON(successResponse("Successfully logged in", user))
}

func getUserById(ctx *fiber.Ctx) error {
	ID, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse("given id is invalid"))
	}

	user := User{ID: ID}
	err = getUserDetails(&user)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return ctx.JSON(failedResopnse("Something went wrong"))
	}

	ctx.Status(http.StatusOK)
	return ctx.JSON(successResponse("Here is your data", &user))
}

package users

import (
	"net/http"
	"strconv"

	"github.com/9thDuck/ecommerce-api.git/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	user := New(userInput.Username, userInput.Email, userInput.Password)
	userRole, err := strconv.Atoi(userInput.Role)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return ctx.JSON(failedResopnse("Something went wrong. Try again later"))
	}
	user.Role = userRole

	err = createUser(user)

	if err != nil {
		utils.LogCustomError("failed to create user", err)
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

	user := New("", userInput.Email, userInput.Password)

	if err = loginUser(user); err != nil {
		utils.LogCustomError("login failed", err)
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(failedResopnse("bad credentials, please check the credentials and try again later"))
	}

	ctx.Status(http.StatusOK)
	return ctx.JSON(successResponse("Successfully logged in", user))
}

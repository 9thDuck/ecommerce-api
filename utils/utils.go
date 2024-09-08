package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetEnvOrThrow(varName string) string {
	val := os.Getenv(varName)

	if val == "" {
		log.Fatal(fmt.Errorf("\n\nerror: env variable missing\n%s\n ", varName))
	}

	return val
}

func LogFatalCustomError(msg string, err error) {
	if err == nil {
		return
	}

	log.Fatal(fmt.Errorf("\n\nerror: %s\n%v\n ", msg, err))
}

func LogCustomError(msg string, err error) {
	if err == nil {
		return
	}

	fmt.Print(fmt.Errorf("\n\nerror: %s\n%s\n ", msg, err))
}

func MakeCookie(name string, value string, expiry time.Time, secure bool, path string) *fiber.Cookie {
	if path == "" {
		path = "/"
	}

	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expiry
	cookie.HTTPOnly = true
	cookie.Secure = secure
	cookie.Domain = path

	return cookie
}

func EnvVarToTimeDuration(varName string, durationType time.Duration) (time.Duration, error) {
	durationStr := GetEnvOrThrow("EXPIRY_REFRESH_TOKEN_DURATION_IN_HOURS")

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return 0, err
	}

	return time.Duration(duration) * durationType, err
}

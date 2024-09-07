package utils

import (
	"fmt"
	"log"
	"os"
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

	log.Fatal(fmt.Errorf("\n\nerror:%s\n%v\n ", msg, err))
}

func LogCustomError(msg string, err error) {
	if err == nil {
		return
	}

	fmt.Print(fmt.Errorf("\n\nerror:%s\n%s\n ", msg, err))
}

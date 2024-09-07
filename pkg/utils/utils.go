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

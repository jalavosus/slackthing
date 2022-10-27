package utils

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

const isDevEnv string = "SLACKTHING_DEV"

func IsDev() bool {
	var isDev bool

	val, ok := os.LookupEnv(isDevEnv)
	if ok {
		b, err := strconv.ParseBool(val)
		if err != nil {
			isDev = false
		} else {
			isDev = b
		}
	} else {
		isDev = false
	}

	return isDev
}
package utils

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

const (
	isDevEnv      string = "SLACKTHING_DEV"
	slackTokenEnv string = "SLACK_OAUTH_TOKEN"
)

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

func SlackToken() string {
	var token string

	val, ok := os.LookupEnv(slackTokenEnv)
	if ok {
		token = val
	} else {
		panic(slackTokenEnv + " not found in environment")
	}

	return token
}

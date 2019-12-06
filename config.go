package main

import (
	"fmt"
	"os"
)

type Config struct {
	telegramToken string
	debug bool
}

func getConfig() (Config, error) {
	telegramTokenVarName := "TELEGRAM_TOKEN"
	telegramToken := os.Getenv(telegramTokenVarName)
	if telegramToken == "" {
		err := fmt.Errorf("environment variable '%s' is required", telegramTokenVarName)
		return Config{}, err
	}
	debug := os.Getenv("DEBUG") == "true"
	return Config{telegramToken: telegramToken, debug: debug}, nil
}

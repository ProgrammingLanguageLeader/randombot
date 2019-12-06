package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {
	config, err := getConfig()
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.telegramToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = config.debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updates, err := bot.GetUpdatesChan(update)
	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
		}
	}
}
package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func configureBot(config *Config) *tgbotapi.BotAPI {
	var err error
	var bot *tgbotapi.BotAPI
	if config.UseProxy {
		client := ConfigureClientProxy(config)
		bot, err = tgbotapi.NewBotAPIWithClient(config.TelegramToken, client)
	} else {
		bot, err = tgbotapi.NewBotAPI(config.TelegramToken)
	}
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = config.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func launchMessageProcessing(bot *tgbotapi.BotAPI) {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updates, err := bot.GetUpdatesChan(update)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
		}
	}
}

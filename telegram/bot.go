package telegram

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"randombot/app/service/message"
	"randombot/config"
	"randombot/net"
)

func ConfigureBot(config *config.Config) *tgbotapi.BotAPI {
	var err error
	var bot *tgbotapi.BotAPI
	if config.UseProxy {
		client := net.ConfigureClientProxy(config)
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

func LaunchMessageProcessing(bot *tgbotapi.BotAPI, service *message.Service) {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updates, err := bot.GetUpdatesChan(update)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		msg := update.Message
		if msg != nil {
			response, replyKeyboardMarkup := service.HandleMessage(msg)
			messageConfig := tgbotapi.NewMessage(msg.Chat.ID, response)
			if replyKeyboardMarkup != nil {
				messageConfig.ReplyMarkup = replyKeyboardMarkup
			} else {
				messageConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}
			sendMessage(bot, &messageConfig)
		}
	}
}

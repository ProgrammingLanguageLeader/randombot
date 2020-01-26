package telegram

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func sendMessage(bot *tgbotapi.BotAPI, message *tgbotapi.MessageConfig) {
	if _, err := bot.Send(message); err != nil {
		log.Println(err)
	}
}

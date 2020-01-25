package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"math/rand"
)

func sendMessage(bot *tgbotapi.BotAPI, message *tgbotapi.MessageConfig) {
	if _, err := bot.Send(message); err != nil {
		log.Println(err)
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message.IsCommand() {
		handleCommand(bot, message.Command(), message.Chat.ID)
	} else {
		responseMessage := tgbotapi.NewMessage(message.Chat.ID, "I work with commands only")
		sendMessage(bot, &responseMessage)
	}
}

func handleCommand(bot *tgbotapi.BotAPI, command string, chatId int64) {
	responseMessage := tgbotapi.NewMessage(chatId, "")
	switch command {
	case "help":
		responseMessage.Text = "type /flipcoin"
	case "flipcoin":
		if rand.Int()%2 == 0 {
			responseMessage.Text = "It's heads!"
		} else {
			responseMessage.Text = "It's tails!"
		}
	}
	sendMessage(bot, &responseMessage)
}

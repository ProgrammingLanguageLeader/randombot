package keyboard

import "github.com/Syfaro/telegram-bot-api"

const Exit string = "Exit"

var exitKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Exit),
	),
)

func GetExitKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	return &exitKeyboard
}

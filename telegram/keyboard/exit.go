package keyboard

import "github.com/Syfaro/telegram-bot-api"

const Exit string = "Exit"

var exitKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Exit),
	),
)

func GetExitKeyboard(lang string) *tgbotapi.ReplyKeyboardMarkup {
	return &exitKeyboard
}

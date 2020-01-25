package keyboard

import "github.com/Syfaro/telegram-bot-api"

const (
	EnglishLanguage string = "English"
	RussianLanguage string = "Russian"
)

var languageKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(EnglishLanguage),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(RussianLanguage),
	),
)

func GetLanguageKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	return &languageKeyboard
}

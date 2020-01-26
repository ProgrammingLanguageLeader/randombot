package keyboard

import "github.com/Syfaro/telegram-bot-api"

const (
	EnglishLanguage string = "English"
	RussianLanguage string = "Russian"
)

var languageSettingsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(EnglishLanguage),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(RussianLanguage),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Exit),
	),
)

func GetLanguageSettingsKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	return &languageSettingsKeyboard
}

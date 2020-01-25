package keyboard

import "github.com/Syfaro/telegram-bot-api"

const (
	RandomGeneratorSettings string = "Random generator settings"
	ChoiceSettings          string = "Choice settings"
	LanguageSettings        string = "Language settings"
)

var settingsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(RandomGeneratorSettings),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(ChoiceSettings),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(LanguageSettings),
	),
)

func GetSettingsKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	return &settingsKeyboard
}

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
		tgbotapi.NewKeyboardButton(ChoiceSettings),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(LanguageSettings),
		tgbotapi.NewKeyboardButton(Exit),
	),
)

func GetSettingsKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	return &settingsKeyboard
}

package keyboard

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	RandomGeneratorSettings = i18n.Message{
		ID:    "button_random-generator-settings",
		Other: "Random generator settings",
	}
	ChoiceSettings = i18n.Message{
		ID:    "button_choice-settings",
		Other: "Choice settings",
	}
	LanguageSettings = i18n.Message{
		ID:    "button_language-settings",
		Other: "Language settings",
	}
)

func GetSettingsKeyboard(lang string) *tgbotapi.ReplyKeyboardMarkup {
	settingsKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&RandomGeneratorSettings, lang)),
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&ChoiceSettings, lang)),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&LanguageSettings, lang)),
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&Exit, lang)),
		),
	)
	return &settingsKeyboard
}

package keyboard

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
	"github.com/Syfaro/telegram-bot-api"
)

const (
	EnglishLanguage string = "English"
	RussianLanguage string = "Русский"
)

func GetLanguageSettingsKeyboard(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(EnglishLanguage),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(RussianLanguage),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&Exit, lang)),
		),
	)
	return &keyboard
}

package keyboard

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var Exit = i18n.Message{
	ID:    "button_exit",
	Other: "Exit",
}

func GetExitKeyboard(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&Exit, lang)),
		),
	)
	return &keyboard
}

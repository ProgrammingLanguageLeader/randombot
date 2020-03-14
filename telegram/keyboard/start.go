package keyboard

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	FlipCoin = i18n.Message{
		ID:    "button_flip-coin",
		Other: "Flip a coin",
	}
	RollDice = i18n.Message{
		ID:    "button_roll-dice",
		Other: "Roll the dice",
	}
	RandomNumber = i18n.Message{
		ID:    "button_make-choice",
		Other: "Random number",
	}
	MakeChoice = i18n.Message{
		ID:    "button_make_choice",
		Other: "Make a choice",
	}
	Settings = i18n.Message{
		ID:    "button_settings",
		Other: "Settings",
	}
	Help = i18n.Message{
		ID:    "button_help",
		Other: "Help",
	}
)

func GetStartKeyboard(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&FlipCoin, lang)),
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&RollDice, lang)),
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&RandomNumber, lang)),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&MakeChoice, lang)),
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&Settings, lang)),
			tgbotapi.NewKeyboardButton(locale.LocalizeSimpleMessage(&Help, lang)),
		),
	)
	return &keyboard
}

package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/telegram/keyboard"
	"github.com/Syfaro/telegram-bot-api"
)

var stateToMarkupGetter = map[string]func(string) *tgbotapi.ReplyKeyboardMarkup{
	redis.StartMenu:        keyboard.GetStartKeyboard,
	redis.ChoiceMenu:       keyboard.GetExitKeyboard,
	redis.LanguageMenu:     keyboard.GetLanguageSettingsKeyboard,
	redis.RandomNumberMenu: keyboard.GetExitKeyboard,
	redis.SettingsMenu:     keyboard.GetSettingsKeyboard,
	redis.DefaultState: func(string) *tgbotapi.ReplyKeyboardMarkup {
		return nil
	},
}

func GetKeyboard(state string, lang string) *tgbotapi.ReplyKeyboardMarkup {
	getMarkup, contains := stateToMarkupGetter[state]
	if !contains {
		return nil
	}
	return getMarkup(lang)
}

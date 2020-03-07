package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/telegram/keyboard"
	"github.com/Syfaro/telegram-bot-api"
)

var stateToKeyboard = map[string]*tgbotapi.ReplyKeyboardMarkup{
	redis.StartMenu:        keyboard.GetStartKeyboard(),
	redis.ChoiceMenu:       keyboard.GetExitKeyboard(),
	redis.LanguageMenu:     keyboard.GetLanguageSettingsKeyboard(),
	redis.RandomNumberMenu: keyboard.GetExitKeyboard(),
	redis.SettingsMenu:     keyboard.GetSettingsKeyboard(),
	redis.DefaultState:     nil,
}

func GetKeyboardByState(state string) *tgbotapi.ReplyKeyboardMarkup {
	markup, contains := stateToKeyboard[state]
	if !contains {
		return nil
	}
	return markup
}

package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/redis/user"
	"github.com/Syfaro/telegram-bot-api"
)

func (service *Service) Exit(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	user.State = redis.StartMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(user)
	}
	response := locale.LocalizeSimpleMessage(&chooseOptionsMessage, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

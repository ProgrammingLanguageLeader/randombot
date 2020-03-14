package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/redis/user"
	"github.com/Syfaro/telegram-bot-api"
)

func (service *Service) Exit(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.StartMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Choose one of the following options", GetKeyboard(user.State, user.LanguageCode)
}

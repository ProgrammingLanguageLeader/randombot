package message

import (
	"github.com/Syfaro/telegram-bot-api"
	"randombot/redis"
	"randombot/redis/user"
)

func (service *Service) Exit(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.StartMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Choose one of the following options", GetKeyboardByState(user.State)
}

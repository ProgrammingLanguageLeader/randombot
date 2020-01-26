package message

import (
	"github.com/Syfaro/telegram-bot-api"
	"randombot/redis"
)

func (service *Service) GoToLanguageSettings(user *redis.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	user.State = redis.LanguageMenu
	err := service.userRepository.Set(user)
	if err != nil {
		service.ProcessError(redis.DefaultState)
	}
	return "Choose a language", GetKeyboardByState(user.State)
}

func (service *Service) GoToRandomGeneratorSettings(user *redis.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.RandomNumberMenu
	err := service.userRepository.Set(user)
	if err != nil {
		service.ProcessError(currentState)
	}
	return "Enter minimum and maximum numbers space separated", GetKeyboardByState(user.State)
}

func (service *Service) GoToChoiceSettings(user *redis.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.ChoiceMenu
	err := service.userRepository.Set(user)
	if err != nil {
		service.ProcessError(currentState)
	}
	return "Enter the choice variants. One item - one line", GetKeyboardByState(user.State)
}

func (service *Service) SwitchLanguage(
	languageCode string,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.StartMenu
	user.LanguageCode = languageCode
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Settings were successfully updated", GetKeyboardByState(user.State)
}

func (service *Service) ChangeRandomGeneratorSettings(
	minNumber int,
	maxNumber int,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.MinRandomNumber = minNumber
	user.MaxRandomNumber = maxNumber
	user.State = redis.StartMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Settings were successfully updated", GetKeyboardByState(user.State)
}

func (service *Service) ChangeChoiceSettings(
	variants []string,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.Variants = variants
	user.State = redis.StartMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Settings were successfully updated", GetKeyboardByState(user.State)
}

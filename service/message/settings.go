package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/redis/user"
	"github.com/Syfaro/telegram-bot-api"
)

func (service *Service) GoToLanguageSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	user.State = redis.LanguageMenu
	err := service.userRepository.Set(user)
	if err != nil {
		service.ProcessError(redis.DefaultState)
	}
	return "Choose a language", GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) GoToRandomGeneratorSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.RandomNumberMenu
	err := service.userRepository.Set(user)
	if err != nil {
		service.ProcessError(currentState)
	}
	return "Enter minimum and maximum numbers space separated", GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) GoToChoiceSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.ChoiceMenu
	err := service.userRepository.Set(user)
	if err != nil {
		service.ProcessError(currentState)
	}
	return "Enter the choice variants. One item - one line", GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) SwitchLanguage(
	languageCode string,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.StartMenu
	user.LanguageCode = languageCode
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Settings were successfully updated", GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) ChangeRandomGeneratorSettings(
	minNumber int,
	maxNumber int,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.MinRandomNumber = minNumber
	user.MaxRandomNumber = maxNumber
	user.State = redis.StartMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Settings were successfully updated", GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) ChangeChoiceSettings(
	variants []string,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.Variants = variants
	user.State = redis.StartMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Settings were successfully updated", GetKeyboard(user.State, user.LanguageCode)
}

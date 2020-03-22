package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/redis/user"
	"github.com/Syfaro/telegram-bot-api"
)

func (service *Service) GoToLanguageSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	err := service.UpdateUserState(user, redis.LanguageMenu)
	if err != nil {
		return service.ProcessError(user)
	}
	response := locale.LocalizeSimpleMessage(&chooseLanguageMessage, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) GoToRandomGeneratorSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	err := service.UpdateUserState(user, redis.RandomNumberMenu)
	if err != nil {
		return service.ProcessError(user)
	}
	response := locale.LocalizeSimpleMessage(&enterMinAndMaxNumbersMessage, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) GoToChoiceSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	err := service.UpdateUserState(user, redis.ChoiceMenu)
	if err != nil {
		return service.ProcessError(user)
	}
	response := locale.LocalizeSimpleMessage(&enterChoiceVariantsMessage, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) SwitchLanguage(lang string, user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	user.LanguageCode = lang
	err := service.UpdateUserState(user, redis.StartMenu)
	if err != nil {
		return service.ProcessError(user)
	}
	response := locale.LocalizeSimpleMessage(&settingsWereSuccessfullyUpdated, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) ChangeRandomGeneratorSettings(
	minNumber int,
	maxNumber int,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	user.MinRandomNumber = minNumber
	user.MaxRandomNumber = maxNumber
	err := service.UpdateUserState(user, redis.StartMenu)
	if err != nil {
		return service.ProcessError(user)
	}
	response := locale.LocalizeSimpleMessage(&settingsWereSuccessfullyUpdated, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) ChangeChoiceSettings(
	variants []string,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	user.Variants = variants
	err := service.UpdateUserState(user, redis.StartMenu)
	if err != nil {
		return service.ProcessError(user)
	}
	response := locale.LocalizeSimpleMessage(&settingsWereSuccessfullyUpdated, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

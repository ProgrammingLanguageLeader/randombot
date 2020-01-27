package message

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"math/rand"
	"randombot/redis"
	"randombot/redis/user"
	"strings"
)

func (service *Service) RegisterUser(tgUser *tgbotapi.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	langCode := strings.Split(tgUser.LanguageCode, "-")[0]
	userInstance := user.User{
		ID:              tgUser.ID,
		LanguageCode:    langCode,
		State:           redis.StartMenu,
		Variants:        nil,
		MinRandomNumber: 1,
		MaxRandomNumber: 10,
	}
	err := service.userRepository.Set(&userInstance)
	if err != nil {
		return service.ProcessError(redis.DefaultState)
	}
	return "Hello! You have successfully registered!", GetKeyboardByState(userInstance.State)
}

func (service *Service) FlipCoin(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	var flipResult string
	if rand.Int()%2 == 0 {
		flipResult = "it's heads!"
	} else {
		flipResult = "it's tails!"
	}
	response := fmt.Sprintf("You have flipped a coin: %s", flipResult)
	return response, GetKeyboardByState(user.State)
}

func (service *Service) RollDice(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	firstDie := rand.Int()%6 + 1
	secondDie := rand.Int()%6 + 1
	response := fmt.Sprintf("You have rolled the dice: %d and %d", firstDie, secondDie)
	return response, GetKeyboardByState(user.State)
}

func (service *Service) GetRandomNumber(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	randNum := rand.Intn(user.MaxRandomNumber-user.MinRandomNumber) + user.MinRandomNumber
	response := fmt.Sprintf(
		"Random number from the range [%d...%d]: %d",
		user.MinRandomNumber,
		user.MaxRandomNumber,
		randNum)
	return response, GetKeyboardByState(user.State)
}

func (service *Service) MakeChoice(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	if user.Variants == nil || len(user.Variants) == 0 {
		return service.GoToChoiceSettings(user)
	}
	choiceIndex := rand.Intn(len(user.Variants))
	choice := user.Variants[choiceIndex]
	response := fmt.Sprintf("Choice between %v: %s", user.Variants, choice)
	return response, GetKeyboardByState(user.State)
}

func (service *Service) GoToSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.SettingsMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	return "Choose one of the following options", GetKeyboardByState(user.State)
}

func (service *Service) GetAbout(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	return "Well, I hope it will be written soon :)", GetKeyboardByState(user.State)
}

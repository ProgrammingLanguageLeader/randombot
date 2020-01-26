package message

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"randombot/redis"
	"randombot/telegram/keyboard"
	"regexp"
	"strings"
)

func (service *Service) HandleMessage(message *tgbotapi.Message) (string, *tgbotapi.ReplyKeyboardMarkup) {
	user, err := service.userRepository.Get(message.From.ID)
	if err != nil {
		return service.ProcessError(redis.DefaultState)
	}
	if message.IsCommand() {
		return service.HandleCommand(message, user)
	}
	switch user.State {
	case redis.StartMenu:
		return service.HandleStartMenu(message, user)
	case redis.ChoiceMenu:
		return service.HandleChoiceMenu(message, user)
	case redis.SettingsMenu:
		return service.HandleSettingsMenu(message, user)
	case redis.RandomNumberMenu:
		return service.HandleRandomNumberMenu(message, user)
	case redis.LanguageMenu:
		return service.HandleLanguageMenu(message, user)
	}
	return service.Exit(user)
}

func (service *Service) HandleCommand(
	message *tgbotapi.Message,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	command := message.Command()
	switch command {
	case "start":
		return service.RegisterUser(message.From)
	case "help", "about":
		return service.GetAbout(user)
	case "flipcoin":
		return service.FlipCoin(user)
	}
	return "Unsupported command", nil
}

func (service *Service) HandleStartMenu(
	message *tgbotapi.Message,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	switch message.Text {
	case keyboard.FlipCoin:
		return service.FlipCoin(user)
	case keyboard.RollDice:
		return service.RollDice(user)
	case keyboard.RandomNumber:
		return service.GetRandomNumber(user)
	case keyboard.MakeChoice:
		return service.MakeChoice(user)
	case keyboard.Settings:
		return service.GoToSettings(user)
	case keyboard.About:
		return service.GetAbout(user)
	}
	return service.ProcessUserMistake(user.State)
}

func (service *Service) HandleChoiceMenu(
	message *tgbotapi.Message,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	if message.Text == keyboard.Exit {
		return service.Exit(user)
	}
	spaceRegexp := regexp.MustCompile(`\s+`)
	input := spaceRegexp.ReplaceAllString(message.Text, " ")
	variants := strings.Split(input, " ")
	currentState := user.State
	if !(2 <= len(variants) && len(variants) <= 16) {
		return "Incorrect input", GetKeyboardByState(currentState)
	}
	return service.ChangeChoiceSettings(variants, user)
}

func (service *Service) HandleSettingsMenu(
	message *tgbotapi.Message,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	switch message.Text {
	case keyboard.RandomGeneratorSettings:
		return service.GoToRandomGeneratorSettings(user)
	case keyboard.ChoiceSettings:
		return service.GoToChoiceSettings(user)
	case keyboard.LanguageSettings:
		return service.GoToLanguageSettings(user)
	case keyboard.Exit:
		return service.Exit(user)
	}
	return service.ProcessUserMistake(user.State)
}

func (service *Service) HandleRandomNumberMenu(
	message *tgbotapi.Message,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	input := message.Text
	if input == keyboard.Exit {
		return service.Exit(user)
	}
	var minRandomNumber int
	var maxRandomNumber int
	successfullyScanned, err := fmt.Sscanf(input, "%d %d", &minRandomNumber, &maxRandomNumber)
	if err != nil || successfullyScanned != 2 {
		return service.ProcessUserMistake(user.State)
	}
	return service.ChangeRandomGeneratorSettings(minRandomNumber, maxRandomNumber, user)
}

func (service *Service) HandleLanguageMenu(
	message *tgbotapi.Message,
	user *redis.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	input := message.Text
	var languageCode = ""
	switch input {
	case keyboard.EnglishLanguage:
		languageCode = "en"
	case keyboard.RussianLanguage:
		languageCode = "ru"
	case keyboard.Exit:
		return service.Exit(user)
	default:
		return service.ProcessUserMistake(user.State)
	}
	return service.SwitchLanguage(languageCode, user)

}

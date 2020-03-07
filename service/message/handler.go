package message

import (
	"fmt"
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/redis/user"
	"github.com/ProgrammingLanguageLeader/randombot/telegram/keyboard"
	"github.com/Syfaro/telegram-bot-api"
	"regexp"
	"strings"
)

func (service *Service) HandleMessage(message *tgbotapi.Message) (string, *tgbotapi.ReplyKeyboardMarkup) {
	userInstance, err := service.userRepository.Get(message.From.ID)
	if err != nil {
		switch err.(type) {
		case *user.DoesNotExist:
			return service.RegisterUser(message.From)
		default:
			return service.ProcessError(redis.DefaultState)
		}
	}
	if message.IsCommand() {
		return service.HandleCommand(message, userInstance)
	}
	switch userInstance.State {
	case redis.StartMenu:
		return service.HandleStartMenu(message, userInstance)
	case redis.ChoiceMenu:
		return service.HandleChoiceMenu(message, userInstance)
	case redis.SettingsMenu:
		return service.HandleSettingsMenu(message, userInstance)
	case redis.RandomNumberMenu:
		return service.HandleRandomNumberMenu(message, userInstance)
	case redis.LanguageMenu:
		return service.HandleLanguageMenu(message, userInstance)
	}
	return service.Exit(userInstance)
}

func (service *Service) HandleCommand(
	message *tgbotapi.Message,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	command := message.Command()
	switch command {
	case "help", "about":
		return service.GetAbout(user)
	case "flipcoin":
		return service.FlipCoin(user)
	}
	return "Unsupported command", nil
}

func (service *Service) HandleStartMenu(
	message *tgbotapi.Message,
	user *user.User,
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
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	if message.Text == keyboard.Exit {
		return service.Exit(user)
	}
	spaceRegexp := regexp.MustCompile(`[ \f\r\t\v]`)
	input := spaceRegexp.ReplaceAllString(message.Text, " ")
	variants := strings.Split(input, "\n")
	currentState := user.State
	if !(2 <= len(variants) && len(variants) <= 16) {
		return "Incorrect input", GetKeyboardByState(currentState)
	}
	return service.ChangeChoiceSettings(variants, user)
}

func (service *Service) HandleSettingsMenu(
	message *tgbotapi.Message,
	user *user.User,
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
	user *user.User,
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
	user *user.User,
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

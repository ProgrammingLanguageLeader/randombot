package message

import (
	"fmt"
	"github.com/ProgrammingLanguageLeader/randombot/locale"
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
	case "help":
		return service.GetHelp(user)
	case "flipcoin":
		return service.FlipCoin(user)
	}
	return "Unsupported command", GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) HandleStartMenu(
	message *tgbotapi.Message,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	switch message.Text {
	case locale.LocalizeSimpleMessage(&keyboard.FlipCoin, user.LanguageCode):
		return service.FlipCoin(user)

	case locale.LocalizeSimpleMessage(&keyboard.RollDice, user.LanguageCode):
		return service.RollDice(user)

	case locale.LocalizeSimpleMessage(&keyboard.RandomNumber, user.LanguageCode):
		return service.GetRandomNumber(user)

	case locale.LocalizeSimpleMessage(&keyboard.MakeChoice, user.LanguageCode):
		return service.MakeChoice(user)

	case locale.LocalizeSimpleMessage(&keyboard.Settings, user.LanguageCode):
		return service.GoToSettings(user)

	case locale.LocalizeSimpleMessage(&keyboard.Help, user.LanguageCode):
		return service.GetHelp(user)
	}
	return service.ProcessUserMistake(user.State)
}

func (service *Service) HandleChoiceMenu(
	message *tgbotapi.Message,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	localizedExit := locale.LocalizeSimpleMessage(&keyboard.Exit, user.LanguageCode)
	if localizedExit == message.Text {
		return service.Exit(user)
	}
	spaceRegexp := regexp.MustCompile(`[ \f\r\t\v]`)
	input := spaceRegexp.ReplaceAllString(message.Text, " ")
	variants := strings.Split(input, "\n")
	currentState := user.State
	if !(2 <= len(variants) && len(variants) <= 16) {
		return "Incorrect input", GetKeyboard(currentState, user.LanguageCode)
	}
	return service.ChangeChoiceSettings(variants, user)
}

func (service *Service) HandleSettingsMenu(
	message *tgbotapi.Message,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	switch message.Text {
	case locale.LocalizeSimpleMessage(&keyboard.RandomGeneratorSettings, user.LanguageCode):
		return service.GoToRandomGeneratorSettings(user)

	case locale.LocalizeSimpleMessage(&keyboard.ChoiceSettings, user.LanguageCode):
		return service.GoToChoiceSettings(user)

	case locale.LocalizeSimpleMessage(&keyboard.LanguageSettings, user.LanguageCode):
		return service.GoToLanguageSettings(user)

	case locale.LocalizeSimpleMessage(&keyboard.Exit, user.LanguageCode):
		return service.Exit(user)
	}
	return service.ProcessUserMistake(user.State)
}

func (service *Service) HandleRandomNumberMenu(
	message *tgbotapi.Message,
	user *user.User,
) (string, *tgbotapi.ReplyKeyboardMarkup) {
	input := message.Text
	localizedExit := locale.LocalizeSimpleMessage(&keyboard.Exit, user.LanguageCode)
	if localizedExit == input {
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

	case locale.LocalizeSimpleMessage(&keyboard.Exit, user.LanguageCode):
		return service.Exit(user)

	default:
		return service.ProcessUserMistake(user.State)
	}
	return service.SwitchLanguage(languageCode, user)

}

package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/redis/user"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"math/rand"
	"strings"
)

const choiceMessageTemplate = `Choice between
{{.Choices}}
===============
{{.Result}}
`

var (
	registrationMessage = i18n.Message{
		ID:    "message_register-success",
		Other: "Hello! You have successfully registered!",
	}

	flippingCoinHeadsMessage = i18n.Message{
		ID:    "message_flipping-coin-heads",
		Other: "it's heads!",
	}

	flippingCoinTailsMessage = i18n.Message{
		ID:    "message_flipping-coin-tails",
		Other: "it's tails!",
	}

	flippingCoinMessage = i18n.Message{
		ID:    "message_flipping-coin",
		Other: "You have flipped a coin: {{.Result}}",
	}

	rollingDiceMessage = i18n.Message{
		ID:    "message_rolling-dice",
		Other: "You have rolled the dice: {{.First}} and {{.Second}}",
	}

	randomizingMessage = i18n.Message{
		ID:    "message_randomizing",
		Other: "Random number from the range [{{.Min}}...{{.Max}}]: {{.Result}}",
	}

	makingChoiceMessage = i18n.Message{
		ID:    "message_making-choice",
		Other: choiceMessageTemplate,
	}

	settingsMessage = i18n.Message{
		ID:    "message_settings",
		Other: "Choose one of the following options",
	}

	helpMessage = i18n.Message{
		ID: "message_help",
		Other: "This is a Telegram bot that enables you to generate some kind of pseudorandom values. " +
			"For example, it has functionality for virtual \"rolling the dice\", \"flipping coin\", " +
			"choice between set of options, etc",
	}
)

func (service *Service) RegisterUser(tgUser *tgbotapi.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	langCode := "en"
	if tgUser.LanguageCode != "" {
		langCode = strings.Split(tgUser.LanguageCode, "-")[0]
	}
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
	response := locale.LocalizeSimpleMessage(&registrationMessage, userInstance.LanguageCode)
	keyboard := GetKeyboard(userInstance.State, userInstance.LanguageCode)
	return response, keyboard
}

func (service *Service) FlipCoin(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	var flipResult string
	if rand.Intn(2) == 0 {
		flipResult = locale.LocalizeSimpleMessage(&flippingCoinHeadsMessage, user.LanguageCode)
	} else {
		flipResult = locale.LocalizeSimpleMessage(&flippingCoinTailsMessage, user.LanguageCode)
	}
	response := locale.LocalizeMessage(&flippingCoinMessage, user.LanguageCode, map[string]interface{}{
		"Result": flipResult,
	})
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) RollDice(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	firstDie := rand.Intn(6) + 1
	secondDie := rand.Intn(6) + 1
	response := locale.LocalizeMessage(&rollingDiceMessage, user.LanguageCode, map[string]interface{}{
		"First":  firstDie,
		"Second": secondDie,
	})
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) GetRandomNumber(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	randNum := rand.Intn(user.MaxRandomNumber-user.MinRandomNumber) + user.MinRandomNumber
	response := locale.LocalizeMessage(&randomizingMessage, user.LanguageCode, map[string]interface{}{
		"Min":    user.MinRandomNumber,
		"Max":    user.MaxRandomNumber,
		"Result": randNum,
	})
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) MakeChoice(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	if user.Variants == nil || len(user.Variants) == 0 {
		return service.GoToChoiceSettings(user)
	}
	choiceIndex := rand.Intn(len(user.Variants))
	choice := user.Variants[choiceIndex]
	joinedVariants := strings.Join(user.Variants, ", \n")
	response := locale.LocalizeMessage(&makingChoiceMessage, user.LanguageCode, map[string]interface{}{
		"Choices": joinedVariants,
		"Result":  choice,
	})
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) GoToSettings(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentState := user.State
	user.State = redis.SettingsMenu
	err := service.userRepository.Set(user)
	if err != nil {
		return service.ProcessError(currentState)
	}
	response := locale.LocalizeSimpleMessage(&settingsMessage, user.LanguageCode)
	return response, GetKeyboard(user.State, user.LanguageCode)
}

func (service *Service) GetHelp(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	answer := locale.LocalizeSimpleMessage(&helpMessage, user.LanguageCode)
	return answer, GetKeyboard(user.State, user.LanguageCode)
}

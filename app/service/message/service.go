package message

import (
	"github.com/Syfaro/telegram-bot-api"
	. "github.com/go-redis/redis"
	"randombot/redis/user"
)

type Service struct {
	userRepository *user.Repository
}

func NewService(dbClient *Client) *Service {
	return &Service{
		userRepository: &user.Repository{
			DbClient: dbClient,
		},
	}
}

func (service *Service) ProcessError(currentState string) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentKeyboard := GetKeyboardByState(currentState)
	return "Something went wrong... Please, try again later", currentKeyboard
}

func (service *Service) ProcessUserMistake(currentState string) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentKeyboard := GetKeyboardByState(currentState)
	return "Sorry, I don't understand you", currentKeyboard
}

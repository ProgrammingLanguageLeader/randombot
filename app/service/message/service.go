package message

import (
	"github.com/Syfaro/telegram-bot-api"
	. "github.com/go-redis/redis"
	"randombot/redis"
)

type Service struct {
	userRepository *redis.UserRepository
}

func NewService(dbClient *Client) *Service {
	return &Service{
		userRepository: &redis.UserRepository{
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

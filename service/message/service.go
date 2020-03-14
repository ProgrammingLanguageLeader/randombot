package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/redis/user"
	"github.com/Syfaro/telegram-bot-api"
	. "github.com/go-redis/redis"
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
	// TODO: pass language code here
	currentKeyboard := GetKeyboard(currentState, "en")
	return "Something went wrong... Please, try again later", currentKeyboard
}

func (service *Service) ProcessUserMistake(currentState string) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentKeyboard := GetKeyboard(currentState, "en")
	return "Sorry, I don't understand you", currentKeyboard
}

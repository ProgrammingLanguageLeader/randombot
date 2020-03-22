package message

import (
	"github.com/ProgrammingLanguageLeader/randombot/locale"
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

func (service *Service) ProcessError(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentKeyboard := GetKeyboard(user.State, user.LanguageCode)
	response := locale.LocalizeSimpleMessage(&somethingWentWrongMessage, user.LanguageCode)
	return response, currentKeyboard
}

func (service *Service) ProcessUserMistake(user *user.User) (string, *tgbotapi.ReplyKeyboardMarkup) {
	currentKeyboard := GetKeyboard(user.State, user.LanguageCode)
	response := locale.LocalizeSimpleMessage(&dontUnderstandMessage, user.LanguageCode)
	return response, currentKeyboard
}

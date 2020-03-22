package message

import "github.com/ProgrammingLanguageLeader/randombot/redis/user"

func (service *Service) UpdateUserState(user *user.User, nextState string) error {
	prevState := user.State
	user.State = nextState
	err := service.userRepository.Set(user)
	if err != nil {
		user.State = prevState
		return err
	}
	return nil
}

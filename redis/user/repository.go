package user

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

type Repository struct {
	DbClient *redis.Client
}

func (repository *Repository) Get(id int) (*User, error) {
	userKey := strconv.Itoa(id)
	userString, err := repository.DbClient.Get(userKey).Result()
	if err == redis.Nil {
		return nil, &DoesNotExist{}
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var user User
	err = json.Unmarshal([]byte(userString), &user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (repository *Repository) Set(user *User) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return err
	}
	userKey := strconv.Itoa(user.ID)
	repository.DbClient.Set(userKey, userJson, 0)
	return nil
}

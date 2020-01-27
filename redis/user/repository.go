package user

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
)

type Repository struct {
	DbClient *redis.Client
}

func (repository *Repository) Get(id int) (*User, error) {
	userString, err := repository.DbClient.Get(string(id)).Result()
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
	repository.DbClient.Set(string(user.ID), userJson, 0)
	return nil
}

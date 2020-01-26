package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
)

type User struct {
	ID              int      `json:"id"`
	LanguageCode    string   `json:"languageCode"`
	State           string   `json:"state"`
	Variants        []string `json:"variants"`
	MinRandomNumber int      `json:"minRandomNumber"`
	MaxRandomNumber int      `json:"maxRandomNumber"`
}

type UserRepository struct {
	DbClient *redis.Client
}

func (repository *UserRepository) Get(id int) (*User, error) {
	userString, err := repository.DbClient.Get(string(id)).Result()
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

func (repository *UserRepository) Set(user *User) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return err
	}
	repository.DbClient.Set(string(user.ID), userJson, 0)
	return nil
}

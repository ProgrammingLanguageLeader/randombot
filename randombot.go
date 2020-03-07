package main

import (
	"github.com/ProgrammingLanguageLeader/randombot/config"
	"github.com/ProgrammingLanguageLeader/randombot/redis"
	"github.com/ProgrammingLanguageLeader/randombot/service/message"
	"github.com/ProgrammingLanguageLeader/randombot/telegram"
)

func main() {
	conf := config.LoadConfig()
	redisClient := redis.CreateRedisClient(conf)
	redis.CheckConnection(redisClient, true)
	messageService := message.NewService(redisClient)

	bot := telegram.ConfigureBot(conf)
	telegram.LaunchMessageProcessing(bot, messageService)
}

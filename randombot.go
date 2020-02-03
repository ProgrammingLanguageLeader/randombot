package main

import (
	"randombot/config"
	"randombot/redis"
	"randombot/service/message"
	"randombot/telegram"
)

func main() {
	conf := config.LoadConfig()
	redisClient := redis.CreateRedisClient(conf)
	redis.CheckConnection(redisClient, true)
	messageService := message.NewService(redisClient)

	bot := telegram.ConfigureBot(conf)
	telegram.LaunchMessageProcessing(bot, messageService)
}

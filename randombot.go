package main

import (
	"randombot/app/service/message"
	"randombot/config"
	"randombot/redis"
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

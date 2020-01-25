package main

import (
	"randombot/config"
	"randombot/telegram"
)

func main() {
	conf := config.LoadConfig()
	//redisClient := redis.CreateRedisClient(conf)
	//redis.CheckConnection(redisClient, true)

	bot := telegram.ConfigureBot(conf)
	telegram.LaunchMessageProcessing(bot)
}

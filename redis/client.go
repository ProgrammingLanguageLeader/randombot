package redis

import (
	"github.com/ProgrammingLanguageLeader/randombot/config"
	"github.com/go-redis/redis"
	"log"
)

func CreateRedisClient(config *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	return client
}

func CheckConnection(client *redis.Client, fatal bool) bool {
	_, err := client.Ping().Result()
	successful := err == nil
	if successful {
		log.Println("Redis is successfully connected to the application")
	} else if fatal {
		log.Fatalf("Redis is not available: %s\n", err)
	}
	return successful
}

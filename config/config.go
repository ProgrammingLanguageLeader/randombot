package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	TelegramToken          string `split_words:"true" required:"true"`
	Debug                  bool   `split_words:"true" default:"false"`
	UseProxy               bool   `split_words:"true" default:"false"`
	ProxyURL               string `split_words:"true"`
	ProxyUsername          string `split_words:"true"`
	ProxyPassword          string `split_words:"true"`
	ProxyTransportProtocol string `split_words:"true" default:"tcp"`
	RedisURL               string `split_words:"true" default:"127.0.0.1:6379"`
	RedisPassword          string `split_words:"true" default:""`
	RedisDB                int    `split_words:"true" default:"0"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load enviromnent variables from file: %s\n", err.Error())
	}
	config := Config{}
	err = envconfig.Process("random_bot", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &config
}

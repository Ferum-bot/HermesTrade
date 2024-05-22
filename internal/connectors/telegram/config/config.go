package config

import (
	"errors"
	"os"
)

type config struct {
}

func NewConfig() Telegram {
	return &config{}
}

func (c *config) GetToken() (string, error) {
	token, exists := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !exists {
		return "", errors.New("environment variable \"TELEGRAM_BOT_TOKEN\" not provided")
	}

	return token, nil
}

func (c *config) GetMongoUrl() (string, error) {
	mongoUrl, exists := os.LookupEnv("MONGODB_URL")
	if !exists {
		return "", errors.New("environment variable \"MONGODB_URL\" not provided")
	}

	return mongoUrl, nil
}

func (c *config) GetMongoDatabase() (string, error) {
	database, exists := os.LookupEnv("MONGODB_DATABASE")
	if !exists {
		return "", errors.New("environment variable \"MONGODB_DATABASE\" not provided")
	}

	return database, nil
}

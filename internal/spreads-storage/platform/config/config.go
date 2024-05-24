package config

import (
	"errors"
	"os"
)

type config struct {
}

func NewConfig() SpreadsStorage {
	return &config{}
}

func (c *config) GetMongoUrl() (string, error) {
	url, exists := os.LookupEnv("SPREADS_STORAGE_MONGODB_URL")
	if !exists {
		return "", errors.New("environment variable \"SPREADS_STORAGE_MONGODB_URL\" not provided")
	}

	return url, nil
}

func (c *config) GetMongoDatabase() (string, error) {
	database, exists := os.LookupEnv("SPREADS_STORAGE_MONGODB_DATABASE")
	if !exists {
		return "", errors.New("environment variable \"SPREADS_STORAGE_MONGODB_DATABASE\" not provided")
	}

	return database, nil
}

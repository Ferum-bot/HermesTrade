package config

import (
	"errors"
	"os"
)

type config struct {
}

func NewConfig() AssetsStorage {
	return &config{}
}

func (c *config) GetMongoUrl() (string, error) {
	url, exists := os.LookupEnv("ASSETS_STORAGE_MONGODB_URL")
	if !exists {
		return "", errors.New("environment variable \"ASSETS_STORAGE_MONGODB_URL\" not provided")
	}

	return url, nil
}

func (c *config) GetMongoDatabase() (string, error) {
	url, exists := os.LookupEnv("ASSETS_STORAGE_MONGODB_DATABASE")
	if !exists {
		return "", errors.New("environment variable \"ASSETS_STORAGE_MONGODB_DATABASE\" not provided")
	}

	return url, nil
}

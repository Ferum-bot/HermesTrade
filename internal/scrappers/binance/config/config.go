package config

import (
	"errors"
	"os"
)

type Config struct {
}

func New() *Config {
	return &Config{}
}

func (config *Config) GetToken() (string, error) {
	token, exists := os.LookupEnv("BINANCE_API_TOKEN")
	if !exists {
		return "", errors.New("environment variable \"BINANCE_API_TOKEN\" not provided")
	}

	return token, nil
}

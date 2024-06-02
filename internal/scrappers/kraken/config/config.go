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
	token, exists := os.LookupEnv("KRAKEN_API_TOKEN")
	if !exists {
		return "", errors.New("environment variable \"KRAKEN_API_TOKEN\" not provided")
	}

	return token, nil
}

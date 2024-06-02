package config

import (
	"errors"
	"os"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetKafkaUrl() (string, error) {
	url, exists := os.LookupEnv("ASSET_SPREAD_HUNTER_KAFKA_URL")
	if !exists {
		return "", errors.New("environment variable \"ASSET_SPREAD_HUNTER_KAFKA_URL\" not provided")
	}

	return url, nil
}

func (c *Config) GetKafkaTopicFoundSpreads() (string, error) {
	url, exists := os.LookupEnv("ASSET_SPREAD_HUNTER_KAFKA_TOPIC_FOUND_SPREADS")
	if !exists {
		return "", errors.New("environment variable \"ASSET_SPREAD_HUNTER_KAFKA_TOPIC_FOUND_SPREADS\" not provided")
	}

	return url, nil
}

package config

type AssetSpreadHunter interface {
	GetKafkaUrl() (string, error)
	GetKafkaTopicFoundSpreads() (string, error)
}

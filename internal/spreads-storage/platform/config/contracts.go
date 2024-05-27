package config

type SpreadsStorage interface {
	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)

	GetKafkaUrl() (string, error)
	GetKafkaTopicFoundSpreads() (string, error)
	GetConsumerGroup() (string, error)
}

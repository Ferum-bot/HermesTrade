package config

type Telegram interface {
	GetToken() (string, error)

	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)
}

package config

type SpreadsStorage interface {
	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)
}

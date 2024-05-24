package config

type AssetsStorage interface {
	GetMongoUrl() (string, error)
	GetMongoDatabase() (string, error)
}

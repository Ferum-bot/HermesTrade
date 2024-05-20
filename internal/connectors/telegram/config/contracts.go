package config

type Telegram interface {
	GetToken() (string, error)
}

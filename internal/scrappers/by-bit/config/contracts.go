package config

type Scrapper interface {
	GetToken() (string, error)
}

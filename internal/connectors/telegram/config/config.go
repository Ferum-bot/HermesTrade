package config

type config struct {
}

func NewConfig() Telegram {
	return &config{}
}

func (c config) GetToken() (string, error) {
	//TODO implement me
	panic("implement me")
}

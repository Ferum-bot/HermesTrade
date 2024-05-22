package config

type config struct {
}

func NewConfig() Telegram {
	return &config{}
}

func (c *config) GetToken() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *config) GetMongoUrl() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *config) GetMongoDatabase() (string, error) {
	//TODO implement me
	panic("implement me")
}

package spread_link_builder

import "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) ProvideLinks(spread model.Spread) model.SpreadWithLink {
	//TODO implement me
	panic("implement me")
}

package converter

import (
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	model2 "github.com/Ferum-Bot/HermesTrade/internal/scrappers/binance/model"
)

type Converter struct {
}

func New() *Converter {
	return &Converter{}
}

func (converter *Converter) Convert(
	assets []model2.AssetCurrencyPair,
) ([]model.AssetCurrencyPair, error) {
	//TODO implement me
	panic("implement me")
}

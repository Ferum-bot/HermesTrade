package currency_ratio

import (
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"math/rand"
)

func ProvideCurrencyRatio() model.AssetsCurrencyRatio {
	precision := rand.Int63() % 3
	value := rand.Int63()%100000 + precision
	return model.AssetsCurrencyRatio{
		Precision: precision,
		Value:     value,
	}
}

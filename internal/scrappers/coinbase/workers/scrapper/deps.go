package scrapper

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/coinbase/model"
)

type assetsStorageSender interface {
	SaveNewAssets(
		ctx context.Context,
		assets []model.AssetCurrencyPair,
	) error
}

type exchangeParser interface {
	ParseNewAssetsPairs(
		ctx context.Context,
	) ([]model.AssetCurrencyPair, error)
}

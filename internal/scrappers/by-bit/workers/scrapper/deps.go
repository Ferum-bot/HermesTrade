package scrapper

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/by-bit/model"
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

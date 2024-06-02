package spread_hunter

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type assetsRetrieverService interface {
	RetrieveNewAssets(
		ctx context.Context,
	) ([]model.AssetCurrencyPair, error)
}

type spreadsHunterService interface {
	FindSpreads(
		ctx context.Context,
		assetPairs []model.AssetCurrencyPair,
	) ([]model.Spread, error)
}

type foundSpreadsProducer interface {
	Produce(
		ctx context.Context,
		spreads []model.Spread,
	) error
}

package spread_hunter

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type spreadHunterAlgorithm interface {
	SearchSpread(
		ctx context.Context,
		assetPairs []model.AssetCurrencyPair,
		searchSettings model.SpreadSearchSettings,
	) ([]model.Spread, error)
}

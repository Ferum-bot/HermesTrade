package spread_hunter

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"

type SyncSpreadHunter interface {
	SearchSpread(
		assetPairs []model.AssetCurrencyPair,
		searchSettings model.SpreadSearchSettings,
	) ([]model.Spread, error)
}

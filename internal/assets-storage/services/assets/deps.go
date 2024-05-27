package assets

import (
	"context"
	model "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
)

type assetsStorage interface {
	SaveAssetsPairs(
		ctx context.Context,
		assetsPairs []model.AssetCurrencyPair,
	) ([]model.AssetCurrencyPair, error)

	SearchAssetsPairs(
		ctx context.Context,
		filters model.AssetFilters,
		offset, limit int64,
	) ([]model.AssetCurrencyPair, error)
}

package get_assets

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
)

type assetsService interface {
	GetAssets(
		ctx context.Context,
		filter model.AssetFilters,
		offset, limit int64,
	) ([]model.AssetCurrencyPair, error)
}

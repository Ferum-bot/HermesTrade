package add_assets

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
)

type assetsService interface {
	AddAssets(
		ctx context.Context,
		assets []model.AddAssetCurrencyPairData,
	) ([]model.AssetCurrencyPair, error)
}

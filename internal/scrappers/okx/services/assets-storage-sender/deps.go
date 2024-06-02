package assets_storage_sender

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	model2 "github.com/Ferum-Bot/HermesTrade/internal/scrappers/okx/model"
)

type assetsStorageClient interface {
	Save(
		ctx context.Context,
		assets []model.AssetCurrencyPair,
	) error
}

type assetsConverter interface {
	Convert(
		assets []model2.AssetCurrencyPair,
	) ([]model.AssetCurrencyPair, error)
}

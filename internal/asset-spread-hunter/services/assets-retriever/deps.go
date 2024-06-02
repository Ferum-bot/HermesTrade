package assets_retriever

import (
	"context"
	assets_storage "github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/clients/assets-storage"
)

type assetsStorageClient interface {
	GetAssets(
		ctx context.Context,
		assetsFilter assets_storage.AssetsFilter,
		offset, limit int64,
	) ([]assets_storage.AssetCurrencyPair, error)
}

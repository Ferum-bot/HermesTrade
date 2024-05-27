package assets

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/google/uuid"
)

type Service struct {
	storage assetsStorage
}

func NewAssetsService(
	storage assetsStorage,
) *Service {
	return &Service{
		storage: storage,
	}
}

func (service *Service) AddAssets(
	ctx context.Context,
	assets []model.AddAssetCurrencyPairData,
) ([]model.AssetCurrencyPair, error) {
	assetsWithIDs := make([]model.AssetCurrencyPair, 0, len(assets))
	for _, asset := range assets {
		assetsWithIDs = append(assetsWithIDs, model.AssetCurrencyPair{
			Identifier:    model.AssetCurrencyPairIdentifier(uuid.New().String()),
			BaseAsset:     asset.BaseAsset,
			QuotedAsset:   asset.QuotedAsset,
			CurrencyRatio: asset.CurrencyRatio,
		})
	}

	addedAssets, err := service.storage.SaveAssetsPairs(ctx, assetsWithIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.storage.SaveAssetsPairs")
	}

	return addedAssets, nil
}

func (service *Service) GetAssets(
	ctx context.Context,
	filter model.AssetFilters,
	offset, limit int64,
) ([]model.AssetCurrencyPair, error) {
	foundAssets, err := service.storage.SearchAssetsPairs(ctx, filter, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "service.storage.SearchAssetsPairs")
	}

	return foundAssets, nil
}

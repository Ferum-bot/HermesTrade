package assets_retriever

import (
	"context"
	assets_storage "github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/clients/assets-storage"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"time"
)

type Service struct {
	logger              logger.Logger
	assetsStorageClient assetsStorageClient
}

func NewService(
	logger logger.Logger,
	assetsStorageClient assetsStorageClient,
) *Service {
	return &Service{
		logger:              logger,
		assetsStorageClient: assetsStorageClient,
	}
}

func (service *Service) RetrieveNewAssets(
	ctx context.Context,
) ([]model.AssetCurrencyPair, error) {
	const retriesCount = 5

	const maxAssetsPerRequest = 100
	const maxTotalAssetsCount = 1000

	const requestsCount = maxTotalAssetsCount/maxAssetsPerRequest + 1

	allAssetsPairs := make([]model.AssetCurrencyPair, 0, maxTotalAssetsCount)

	var assetsStartTime = time.Now().Add(-5 * time.Minute)
	var assetsEndTime = time.Now()
	assetsFilter := assets_storage.AssetsFilter{
		TimeFilter: &assets_storage.AssetsTimeFilter{
			StartTime: &assetsStartTime,
			EndTime:   &assetsEndTime,
		},
	}

	currentOffset := int64(0)
	for i := 0; i < requestsCount; i++ {
		var assets []model.AssetCurrencyPair

		for retryAttempt := 1; retryAttempt <= retriesCount; retryAttempt++ {
			receivedAssets, err := service.assetsStorageClient.GetAssets(
				ctx, assetsFilter, currentOffset, maxAssetsPerRequest,
			)
			if err != nil {
				service.logger.Warnf(
					"Received an error retrieving assets(attempt %d): %v", retryAttempt, err,
				)

				if retryAttempt == retriesCount {
					return nil, errors.Wrap(err, "Failed to retrieve assets after all attempts")
				}

				time.Sleep(10 * time.Millisecond)
				continue
			}

			assets = convertAssets(receivedAssets)
		}

		allAssetsPairs = append(allAssetsPairs, assets...)

		currentOffset += int64(len(assets))
		if len(assets) == 0 {
			break
		}
	}

	return allAssetsPairs, nil
}

func convertAssets(assets []assets_storage.AssetCurrencyPair) []model.AssetCurrencyPair {
	resultAssets := make([]model.AssetCurrencyPair, 0, len(assets))

	for _, asset := range assets {
		resultAssets = append(resultAssets, model.AssetCurrencyPair{
			Identifier:    model.AssetPairIdentifier(asset.Identifier),
			BaseAsset:     model.Asset{},
			QuotedAsset:   model.Asset{},
			CurrencyRatio: model.AssetsCurrencyRatio{},
		})
	}

	return resultAssets
}

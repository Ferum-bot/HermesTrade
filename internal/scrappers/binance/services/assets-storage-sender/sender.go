package assets_storage_sender

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/binance/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type AssetsStorageSender struct {
	assetsStorageClient assetsStorageClient
	assetsConverter     assetsConverter
}

func New(
	assetsStorageClient assetsStorageClient,
	assetsConverter assetsConverter,
) *AssetsStorageSender {
	return &AssetsStorageSender{
		assetsStorageClient: assetsStorageClient,
		assetsConverter:     assetsConverter,
	}
}

func (sender *AssetsStorageSender) SaveNewAssets(
	ctx context.Context,
	assets []model.AssetCurrencyPair,
) error {
	convertedAssets, err := sender.assetsConverter.Convert(assets)
	if err != nil {
		return errors.Wrap(err, "sender.assetsConverter.Convert")
	}

	err = sender.assetsStorageClient.Save(ctx, convertedAssets)
	if err != nil {
		return errors.Wrap(err, "sender.assetsStorageClient.Save")
	}

	return nil
}

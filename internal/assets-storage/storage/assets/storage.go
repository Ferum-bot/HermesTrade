package assets

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	"go.mongodb.org/mongo-driver/mongo"
	"math"
)

const collectionName = "assetsCurrencyPairs"

type Storage struct {
	collection  *mongo.Collection
	savedAssets []model.AssetCurrencyPair
}

func NewAssetsStorage(
	database *mongo.Database,
) *Storage {
	return &Storage{
		collection: database.Collection(collectionName),
	}
}

func (storage *Storage) SaveAssetsPairs(
	ctx context.Context,
	assetsPairs []model.AssetCurrencyPair,
) ([]model.AssetCurrencyPair, error) {
	storage.savedAssets = append(storage.savedAssets, assetsPairs...)
	return assetsPairs, nil
}

func (storage *Storage) SearchAssetsPairs(
	ctx context.Context,
	filters model.AssetFilters,
	offset, limit int64,
) ([]model.AssetCurrencyPair, error) {
	right := math.Min(float64(offset+limit+1), float64(len(storage.savedAssets)))
	return storage.savedAssets[offset:int(right)], nil
}

package assets

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "assetsCurrencyPairs"

type Storage struct {
	collection *mongo.Collection
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
	//TODO implement me
	panic("implement me")
}

func (storage *Storage) SearchAssetsPairs(
	ctx context.Context,
	filters model.AssetFilters,
	offset, limit int64,
) ([]model.AssetCurrencyPair, error) {
	//TODO implement me
	panic("implement me")
}

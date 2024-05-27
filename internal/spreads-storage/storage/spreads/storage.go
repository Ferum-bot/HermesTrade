package spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "spreads"

type Storage struct {
	collection *mongo.Collection
}

func NewStorage(
	database *mongo.Database,
) *Storage {
	return &Storage{
		collection: database.Collection(collectionName),
	}
}

func (s *Storage) GetSpreadsByIDs(
	ctx context.Context,
	spreadIDs []model.SpreadIdentifier,
) ([]model.SpreadWithLink, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) SearchSpreads(
	ctx context.Context,
	filter model.SpreadsFilter,
	offset, limit int64,
) ([]model.Spread, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) SaveSpreads(
	ctx context.Context,
	spreads []model.Spread,
) ([]model.Spread, error) {
	//TODO implement me
	panic("implement me")
}

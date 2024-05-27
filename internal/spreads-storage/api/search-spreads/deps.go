package search_spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
)

type spreadService interface {
	SearchSpreads(
		ctx context.Context,
		filter model.SpreadsFilter,
		offset, limit int64,
	) ([]model.Spread, error)
}

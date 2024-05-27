package spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
)

type spreadsStorage interface {
	GetSpreadsByIDs(
		ctx context.Context,
		spreadIDs []model.SpreadIdentifier,
	) ([]model.SpreadWithLink, error)

	SearchSpreads(
		ctx context.Context,
		filter model.SpreadsFilter,
		offset, limit int64,
	) ([]model.Spread, error)

	SaveSpreads(
		ctx context.Context,
		spreads []model.Spread,
	) ([]model.Spread, error)
}

type spreadLinkBuilder interface {
	ProvideLinks(spread model.Spread) model.SpreadWithLink
}

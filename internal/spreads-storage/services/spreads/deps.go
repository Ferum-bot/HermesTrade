package spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
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
	ProvideLinks(
		spread model.Spread,
	) model.SpreadWithLink

	ProvideLinksForAssetPair(
		assetPair model2.AssetCurrencyPair,
	) model.AssetPairWithLinks
}

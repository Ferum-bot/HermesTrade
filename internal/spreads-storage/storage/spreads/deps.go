package spreads

import (
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type spreadLinkBuilder interface {
	ProvideLinksForAssetPair(
		assetPair model2.AssetCurrencyPair,
	) model.AssetPairWithLinks
}

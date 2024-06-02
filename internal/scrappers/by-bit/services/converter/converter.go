package converter

import (
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	model2 "github.com/Ferum-Bot/HermesTrade/internal/scrappers/by-bit/model"
)

type Converter struct {
}

func New() *Converter {
	return &Converter{}
}

func (converter *Converter) Convert(
	assets []model2.AssetCurrencyPair,
) ([]model.AssetCurrencyPair, error) {
	result := make([]model.AssetCurrencyPair, 0, len(assets))

	for _, asset := range assets {
		result = append(result, model.AssetCurrencyPair{
			Identifier: model.AssetCurrencyPairIdentifier(asset.Identifier),
			BaseAsset: model.Asset{
				SourceIdentifier:    model.AssetSourceIdentifier(asset.BaseAsset.SourceIdentifier),
				UniversalIdentifier: model.AssetUniversalIdentifier(asset.BaseAsset.UniversalIdentifier),
				ExternalIdentifier:  model.AssetExternalIdentifier(asset.BaseAsset.ExternalIdentifier),
			},
			QuotedAsset: model.Asset{
				SourceIdentifier:    model.AssetSourceIdentifier(asset.QuotedAsset.SourceIdentifier),
				UniversalIdentifier: model.AssetUniversalIdentifier(asset.QuotedAsset.UniversalIdentifier),
				ExternalIdentifier:  model.AssetExternalIdentifier(asset.QuotedAsset.ExternalIdentifier),
			},
			CurrencyRatio: model.AssetCurrencyRatio{
				Value:     asset.CurrencyRatio.Value,
				Precision: asset.CurrencyRatio.Precision,
			},
		})
	}

	return result, nil
}

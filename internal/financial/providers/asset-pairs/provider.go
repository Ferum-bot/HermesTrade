package asset_pairs

import (
	"fmt"
	"github.com/Ferum-Bot/HermesTrade/internal/financial/constants"
	currency_ratio "github.com/Ferum-Bot/HermesTrade/internal/financial/providers/currency-ratio"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"math/rand"
)

func ProvideAllAvailableCurrencyPairs() []model.AssetCurrencyPair {
	assetPairs := make([]model.AssetCurrencyPair, 0)

	for sourceIdentifier := range constants.SourceNameByIdentifier {
		sourceAssetPairs := ProvideAllSourceAvailableCurrencyPairs(model.AssetSourceIdentifier(sourceIdentifier))

		assetPairs = append(assetPairs, sourceAssetPairs...)
	}

	return assetPairs
}

func ProvideAllSourceAvailableCurrencyPairs(sourceIdentifier model.AssetSourceIdentifier) []model.AssetCurrencyPair {
	assetPairs := make([]model.AssetCurrencyPair, 0)

	for firstAssetIdentifier := range constants.AssetNameByIdentifier {
		for secondAssetIdentifier := range constants.AssetNameByIdentifier {

			if firstAssetIdentifier == secondAssetIdentifier {
				continue
			}

			firstAsset := model.Asset{
				UniversalIdentifier: model.AssetUniversalIdentifier(firstAssetIdentifier),
				ExternalIdentifier:  model.AssetExternalIdentifier(rand.Int63()),
				SourceIdentifier:    sourceIdentifier,
			}
			secondAsset := model.Asset{
				UniversalIdentifier: model.AssetUniversalIdentifier(secondAssetIdentifier),
				ExternalIdentifier:  model.AssetExternalIdentifier(rand.Int63()),
				SourceIdentifier:    sourceIdentifier,
			}

			currencyRatio := currency_ratio.ProvideCurrencyRatio()

			firstAssetName := constants.AssetNameByIdentifier[firstAssetIdentifier]
			secondAssetName := constants.AssetNameByIdentifier[secondAssetIdentifier]
			assetPairIdentifier := fmt.Sprintf("%s/%s", firstAssetName, secondAssetName)

			assetPair := model.AssetCurrencyPair{
				Identifier:    model.AssetPairIdentifier(assetPairIdentifier),
				BaseAsset:     firstAsset,
				QuotedAsset:   secondAsset,
				CurrencyRatio: currencyRatio,
			}

			assetPairs = append(assetPairs, assetPair)
		}
	}

	return assetPairs
}

func ProvideAllNextAssetPair(
	allAssetPairs []model.AssetCurrencyPair,
	currentAssetPair model.AssetCurrencyPair,
) []model.AssetCurrencyPair {
	currentSourceIdentifier := currentAssetPair.BaseAsset.SourceIdentifier

	baseUniversalIdentifier := currentAssetPair.BaseAsset.UniversalIdentifier
	quotedUniversalIdentifier := currentAssetPair.QuotedAsset.UniversalIdentifier

	result := make([]model.AssetCurrencyPair, 0)

	for _, assetPair := range allAssetPairs {

		if assetPair.BaseAsset.SourceIdentifier == currentSourceIdentifier {
			continue
		}

		if assetPair.BaseAsset.UniversalIdentifier != quotedUniversalIdentifier {
			continue
		}
		if assetPair.QuotedAsset.UniversalIdentifier == baseUniversalIdentifier {
			continue
		}

		result = append(result, assetPair)
	}

	return result
}

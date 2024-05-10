package asset_graph_builder_test

import (
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	graph_builders "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-builders"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"strconv"
	"testing"
)

func buildVertexesForOneAssetPair(
	t *testing.T,
	assetPair model2.AssetCurrencyPair,
) []model.GraphVertex {
	result := make([]model.GraphVertex, 0, 1)

	baseVertex := buildVertexFromAsset(assetPair.BaseAsset)
	quotedVertex := buildVertexFromAsset(assetPair.QuotedAsset)

	commonPrecision := calculateCommonPrecision([]model2.AssetCurrencyPair{assetPair})
	edgeWeight, err := calculateAssetRationFromEdgeWeight(assetPair.CurrencyRatio, commonPrecision)
	if err != nil {
		t.Fatal(err)
	}

	graph_builders.BuildEdgeWithWeight(baseVertex, quotedVertex, edgeWeight)

	result = append(result, *baseVertex)
	result = append(result, *quotedVertex)

	return result
}

func buildVertexesForAssetsInOneSource(
	t *testing.T,
	assetPairs []model2.AssetCurrencyPair,
) []model.GraphVertex {
	result := make([]model.GraphVertex, 0, 1)

	commonPrecision := calculateCommonPrecision(assetPairs)

	for _, assetPair := range assetPairs {
		baseVertex := buildVertexFromAsset(assetPair.BaseAsset)
		quotedVertex := buildVertexFromAsset(assetPair.QuotedAsset)

		edgeWeight, err := calculateAssetRationFromEdgeWeight(assetPair.CurrencyRatio, commonPrecision)
		if err != nil {
			t.Fatal(err)
		}

		graph_builders.BuildEdgeWithWeight(baseVertex, quotedVertex, edgeWeight)

		result = append(result, *baseVertex)
		result = append(result, *quotedVertex)
	}

	return result
}

func buildVertexesForTwoAssetsInDifferentSources(
	t *testing.T,
	firstPair model2.AssetCurrencyPair,
	secondPair model2.AssetCurrencyPair,
) []model.GraphVertex {
	result := make([]model.GraphVertex, 0, 1)

	commonPrecision := calculateCommonPrecision([]model2.AssetCurrencyPair{firstPair, secondPair})

	firstBaseVertex := buildVertexFromAsset(firstPair.BaseAsset)
	firstQuotedVertex := buildVertexFromAsset(firstPair.QuotedAsset)

	firstEdgeWeight, err := calculateAssetRationFromEdgeWeight(firstPair.CurrencyRatio, commonPrecision)
	if err != nil {
		t.Fatal(err)
	}

	graph_builders.BuildEdgeWithWeight(firstBaseVertex, firstQuotedVertex, firstEdgeWeight)

	secondBaseVertex := buildVertexFromAsset(secondPair.BaseAsset)
	secondQuotedVertex := buildVertexFromAsset(secondPair.QuotedAsset)

	secondEdgeWeight, err := calculateAssetRationFromEdgeWeight(secondPair.CurrencyRatio, commonPrecision)
	if err != nil {
		t.Fatal(err)
	}

	graph_builders.BuildEdgeWithWeight(secondBaseVertex, secondQuotedVertex, secondEdgeWeight)

	if firstPair.QuotedAsset.UniversalIdentifier == secondPair.BaseAsset.UniversalIdentifier {
		graph_builders.BuildEdge(firstQuotedVertex, secondBaseVertex)
	}

	if secondPair.QuotedAsset.UniversalIdentifier == firstPair.BaseAsset.UniversalIdentifier {
		graph_builders.BuildEdge(secondQuotedVertex, firstBaseVertex)
	}

	result = append(result, *firstBaseVertex)
	result = append(result, *firstQuotedVertex)
	result = append(result, *secondBaseVertex)
	result = append(result, *secondQuotedVertex)

	return result
}

func buildVertexesForCycleAssetPairs(
	t *testing.T,
	assetPairs []model2.AssetCurrencyPair,
) []model.GraphVertex {
	vertexes := make([]*model.GraphVertex, 0, 2*len(assetPairs))
	vertexByExternalIdentifier := make(map[model2.AssetExternalIdentifier]*model.GraphVertex, 2*len(assetPairs))

	commonPrecision := calculateCommonPrecision(assetPairs)

	for _, assetPair := range assetPairs {
		baseVertex := vertexByExternalIdentifier[assetPair.BaseAsset.ExternalIdentifier]
		quotedVertex := vertexByExternalIdentifier[assetPair.QuotedAsset.ExternalIdentifier]

		if baseVertex == nil {
			baseVertex = buildVertexFromAsset(assetPair.BaseAsset)
			vertexes = append(vertexes, baseVertex)
			vertexByExternalIdentifier[assetPair.BaseAsset.ExternalIdentifier] = baseVertex
		}
		if quotedVertex == nil {
			quotedVertex = buildVertexFromAsset(assetPair.QuotedAsset)
			vertexes = append(vertexes, quotedVertex)
			vertexByExternalIdentifier[assetPair.QuotedAsset.ExternalIdentifier] = quotedVertex
		}

		firstEdgeWeight, err := calculateAssetRationFromEdgeWeight(assetPair.CurrencyRatio, commonPrecision)
		if err != nil {
			t.Fatal(err)
		}

		graph_builders.BuildEdgeWithWeight(baseVertex, quotedVertex, firstEdgeWeight)
	}

	for i := 0; i < len(assetPairs); i++ {
		for j := 0; j < len(assetPairs); j++ {
			if i == j {
				continue
			}

			firstPair := assetPairs[i]
			secondPair := assetPairs[j]

			if firstPair.QuotedAsset.SourceIdentifier != secondPair.BaseAsset.SourceIdentifier {
				if firstPair.QuotedAsset.UniversalIdentifier == secondPair.BaseAsset.UniversalIdentifier {

					quotedVertex := vertexByExternalIdentifier[firstPair.QuotedAsset.ExternalIdentifier]
					baseVertex := vertexByExternalIdentifier[secondPair.BaseAsset.ExternalIdentifier]

					graph_builders.BuildEdgeWithWeight(quotedVertex, baseVertex, model.EdgeWeight(1))
				}
			}
		}
	}

	result := make([]model.GraphVertex, 0, len(vertexes))
	for _, vertex := range vertexes {
		result = append(result, *vertex)
	}

	return result
}

func buildVertexFromAsset(asset model2.Asset) *model.GraphVertex {
	return &model.GraphVertex{
		Identifier: model.GraphVertexIdentifier(asset.ExternalIdentifier),
		Edges:      make([]model.VertexEdge, 0),
	}
}

func calculateCommonPrecision(assets []model2.AssetCurrencyPair) int64 {
	const maxPrecision = 10

	commonPrecision := int64(0)

	for _, asset := range assets {
		currencyRatio := asset.CurrencyRatio

		if currencyRatio.Precision > commonPrecision {
			commonPrecision = currencyRatio.Precision
		}
	}

	if commonPrecision > maxPrecision {
		commonPrecision = maxPrecision
	}

	return commonPrecision
}

func calculateAssetRationFromEdgeWeight(
	assetRation model2.AssetsCurrencyRatio,
	commonPrecision int64,
) (model.EdgeWeight, error) {
	ratioValueString := strconv.FormatInt(assetRation.Value, 10)

	if commonPrecision < assetRation.Precision {
		needlessValuesCount := assetRation.Precision - commonPrecision
		rationValueLen := int64(len(ratioValueString))

		ratioValueString = ratioValueString[:(rationValueLen - needlessValuesCount)]
	}

	if commonPrecision > assetRation.Precision {
		zerosToAdd := commonPrecision - assetRation.Precision

		for i := 0; i < int(zerosToAdd); i++ {
			ratioValueString += "0"
		}
	}

	weight, err := strconv.ParseInt(ratioValueString, 10, 64)
	if err != nil {
		return model.EdgeWeight(0), errors.Wrap(err, "strconv.ParseInt")
	}

	return model.EdgeWeight(weight), nil
}

package asset_graph_builder

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	graph_builders "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-builders"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"strconv"
)

type graphVertexesByIdentifier map[model2.AssetExternalIdentifier]*model.GraphVertex

const sourceEdgeWeight = model.EdgeWeight(1)
const maxAvailablePrecision = 10

type AssetGraphBuilder struct {
}

func NewAssetGraphBuilder() *AssetGraphBuilder {
	return &AssetGraphBuilder{}
}

func (graphBuilder *AssetGraphBuilder) BuildGraphFromAssets(
	ctx context.Context,
	assets []model2.AssetCurrencyPair,
) (model.Graph, error) {
	vertexes := graphBuilder.createEmptyVertexes(assets)

	err := graphBuilder.addAllEdges(assets, vertexes)
	if err != nil {
		return model.Graph{}, errors.Wrap(err, "graphBuilder.addAllEdges")
	}

	graph := model.Graph{
		Vertexes: make([]model.GraphVertex, 0, len(vertexes)),
	}

	for _, vertex := range vertexes {
		graph.Vertexes = append(graph.Vertexes, *vertex)
	}

	return graph, nil
}

func (graphBuilder *AssetGraphBuilder) createEmptyVertexes(
	assets []model2.AssetCurrencyPair,
) graphVertexesByIdentifier {
	vertexes := make(map[model2.AssetExternalIdentifier]*model.GraphVertex, 2*len(assets))

	for _, asset := range assets {
		baseAsset := asset.BaseAsset
		quotedAsset := asset.QuotedAsset

		baseVertex := model.GraphVertex{
			Identifier: model.GraphVertexIdentifier(baseAsset.ExternalIdentifier),
			Edges:      make([]model.VertexEdge, 0),
		}
		quotedVertex := model.GraphVertex{
			Identifier: model.GraphVertexIdentifier(quotedAsset.ExternalIdentifier),
			Edges:      make([]model.VertexEdge, 0),
		}

		vertexes[baseAsset.ExternalIdentifier] = &baseVertex
		vertexes[quotedAsset.ExternalIdentifier] = &quotedVertex
	}

	return vertexes
}

func (graphBuilder *AssetGraphBuilder) addAllEdges(
	assets []model2.AssetCurrencyPair,
	vertexes graphVertexesByIdentifier,
) error {
	err := graphBuilder.addAssetCurrencyPairsEdges(assets, vertexes)
	if err != nil {
		return errors.Wrap(err, "graphBuilder.addAssetCurrencyPairsEdges")
	}

	graphBuilder.addAssetSourcesEdges(assets, vertexes)

	return nil
}

func (graphBuilder *AssetGraphBuilder) addAssetCurrencyPairsEdges(
	assets []model2.AssetCurrencyPair,
	vertexes graphVertexesByIdentifier,
) error {
	commonPrecision := graphBuilder.calculateCommonPrecision(assets)

	for _, asset := range assets {
		baseAsset := asset.BaseAsset
		quotedAsset := asset.QuotedAsset

		baseVertex := vertexes[baseAsset.ExternalIdentifier]
		quotedVertex := vertexes[quotedAsset.ExternalIdentifier]

		edgeWeight, err := graphBuilder.convertAssetRationToEdgeWeight(asset.CurrencyRatio, commonPrecision)
		if err != nil {
			return errors.Wrap(err, "graphBuilder.convertAssetRationToEdgeWeight")
		}

		baseVertex.Edges = append(baseVertex.Edges, model.VertexEdge{
			TargetVertex: quotedVertex,
			Weight:       edgeWeight,
		})
	}

	return nil
}

func (graphBuilder *AssetGraphBuilder) addAssetSourcesEdges(
	assetPairs []model2.AssetCurrencyPair,
	vertexes graphVertexesByIdentifier,
) {

	for i := range assetPairs {
		for j := range assetPairs {
			if i == j {
				continue
			}

			firstQuotedAsset := assetPairs[i].QuotedAsset
			secondBaseAsset := assetPairs[j].BaseAsset

			if firstQuotedAsset.SourceIdentifier == secondBaseAsset.SourceIdentifier {
				continue
			}

			if firstQuotedAsset.UniversalIdentifier == secondBaseAsset.UniversalIdentifier {
				sourceVertex := vertexes[firstQuotedAsset.ExternalIdentifier]
				targetVertex := vertexes[secondBaseAsset.ExternalIdentifier]

				graph_builders.BuildEdgeWithWeight(sourceVertex, targetVertex, sourceEdgeWeight)
			}
		}
	}
}

func (graphBuilder *AssetGraphBuilder) calculateCommonPrecision(
	assets []model2.AssetCurrencyPair,
) int64 {
	commonPrecision := int64(0)

	for _, asset := range assets {
		currencyRatio := asset.CurrencyRatio

		if currencyRatio.Precision > commonPrecision {
			commonPrecision = currencyRatio.Precision
		}
	}

	if commonPrecision > maxAvailablePrecision {
		commonPrecision = maxAvailablePrecision
	}

	return commonPrecision
}

func (graphBuilder *AssetGraphBuilder) convertAssetRationToEdgeWeight(
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

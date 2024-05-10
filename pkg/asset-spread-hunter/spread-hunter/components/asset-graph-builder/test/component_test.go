package asset_graph_builder_test

import (
	"context"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	asset_graph_builder "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/components/asset-graph-builder"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssetGraphBuilder_BuildGraphFromAssets(t *testing.T) {
	type testCase struct {
		inputAssetsPairs []model.AssetCurrencyPair
		expectedVertexes []model2.GraphVertex
		expectedErr      error
	}

	assetsPairs := []model.AssetCurrencyPair{

		// First source
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  11,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  12,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     123,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  14,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  13,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1123,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 5,
				ExternalIdentifier:  15,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 6,
				ExternalIdentifier:  16,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 0,
				Value:     90,
			},
		},

		// Second source
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  21,
				SourceIdentifier:    2,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  22,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     123,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  23,
				SourceIdentifier:    2,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  24,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     123,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 6,
				ExternalIdentifier:  25,
				SourceIdentifier:    2,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 5,
				ExternalIdentifier:  26,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     123,
			},
		},
	}

	cycleAssetsPairs := []model.AssetCurrencyPair{
		// First Source
		{
			Identifier: model.AssetPairIdentifier("EUR/USD"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  102,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  101,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1123,
			},
		},
		{
			Identifier: model.AssetPairIdentifier("USD/EUR"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  101,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  102,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     8,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 52,
				ExternalIdentifier:  202,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 51,
				ExternalIdentifier:  201,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1234,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 51,
				ExternalIdentifier:  201,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 52,
				ExternalIdentifier:  202,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1234,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 62,
				ExternalIdentifier:  302,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 61,
				ExternalIdentifier:  301,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1234,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 61,
				ExternalIdentifier:  301,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 62,
				ExternalIdentifier:  302,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1234,
			},
		},

		// Second Source
		{
			Identifier: model.AssetPairIdentifier("USD/RUB"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  902,
				SourceIdentifier:    2,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  901,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 0,
				Value:     90,
			},
		},
		{
			Identifier: model.AssetPairIdentifier("RUB/USD"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  901,
				SourceIdentifier:    2,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  902,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1,
			},
		},

		// Third Source
		{
			Identifier: model.AssetPairIdentifier("RUB/GBP"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  801,
				SourceIdentifier:    3,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  802,
				SourceIdentifier:    3,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     1,
			},
		},
		{
			Identifier: model.AssetPairIdentifier("GBP/RUB"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  802,
				SourceIdentifier:    3,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  801,
				SourceIdentifier:    3,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 0,
				Value:     10,
			},
		},

		// Fours source
		{
			Identifier: model.AssetPairIdentifier("GBP/EUR"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  702,
				SourceIdentifier:    4,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  701,
				SourceIdentifier:    4,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     1,
			},
		},
		{
			Identifier: model.AssetPairIdentifier("EUR/GBP"),
			BaseAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  701,
				SourceIdentifier:    4,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  702,
				SourceIdentifier:    4,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 0,
				Value:     1,
			},
		},
	}

	tests := map[string]testCase{
		"empty_assets_pairs": {
			inputAssetsPairs: []model.AssetCurrencyPair{},
			expectedVertexes: []model2.GraphVertex{},
		},
		"one_asset_pair": {
			inputAssetsPairs: assetsPairs[0:1],
			expectedVertexes: buildVertexesForOneAssetPair(t, assetsPairs[0]),
		},
		"three_asset_pairs_in_one_source": {
			inputAssetsPairs: assetsPairs[0:3],
			expectedVertexes: buildVertexesForAssetsInOneSource(t, assetsPairs[0:3]),
		},
		"two_asset_pairs_in_different_sources": {
			inputAssetsPairs: append(assetsPairs[0:1], assetsPairs[3:4]...),
			expectedVertexes: buildVertexesForTwoAssetsInDifferentSources(t, assetsPairs[0], assetsPairs[3]),
		},
		"many_cycles_and_many_asset_pairs": {
			inputAssetsPairs: cycleAssetsPairs,
			expectedVertexes: buildVertexesForCycleAssetPairs(t, cycleAssetsPairs),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			graphBuilder := asset_graph_builder.NewAssetGraphBuilder()

			actualGraph, actualErr := graphBuilder.BuildGraphFromAssets(ctx, test.inputAssetsPairs)
			if test.expectedErr != nil {
				assert.Error(t, actualErr)
				assert.Equal(t, test.expectedErr, actualErr)
			}

			assert.NoError(t, actualErr)
			assert.Equal(t, len(test.expectedVertexes), len(actualGraph.Vertexes))

			for _, actualVertex := range actualGraph.Vertexes {
				assertContainsVertex(t, test.expectedVertexes, actualVertex)
			}
		})
	}
}

func assertContainsVertex(
	t *testing.T,
	allVertexes []model2.GraphVertex,
	targetVertex model2.GraphVertex,
) {
	for _, vertex := range allVertexes {
		if vertexIsEqual(vertex, targetVertex) {
			return
		}
	}

	t.Fatalf("all vertexes(%v) not contains target vertex(%v)", allVertexes, targetVertex)
}

func vertexIsEqual(
	firstVertex model2.GraphVertex,
	secondVertex model2.GraphVertex,
) bool {
	if firstVertex.Identifier != secondVertex.Identifier {
		return false
	}

	firstEdges := firstVertex.Edges
	secondEdges := secondVertex.Edges

	if len(firstEdges) != len(secondEdges) {
		return false
	}

	for _, firstEdge := range firstEdges {
		foundEqualEdges := false

		for _, secondEdge := range secondEdges {
			if edgesIsEqual(firstEdge, secondEdge) {
				foundEqualEdges = true
				break
			}
		}

		if !foundEqualEdges {
			return false
		}
	}

	return true
}

func edgesIsEqual(
	firstEdge model2.VertexEdge,
	secondEdge model2.VertexEdge,
) bool {
	if firstEdge.Weight != secondEdge.Weight {
		return false
	}

	if firstEdge.TargetVertex.Identifier != secondEdge.TargetVertex.Identifier {
		return false
	}

	return true
}

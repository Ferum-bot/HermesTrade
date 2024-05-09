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
		inputAssetsPair  []model.AssetCurrencyPair
		expectedVertexes []model2.GraphVertex
		expectedErr      error
	}

	tests := map[string]testCase{
		"empty_assets_pairs": {
			inputAssetsPair:  []model.AssetCurrencyPair{},
			expectedVertexes: []model2.GraphVertex{},
		},
		"one_asset_pair": {
			inputAssetsPair: []model.AssetCurrencyPair{
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
			},
			expectedVertexes: buildVertexesForOneAssetPair(),
		},
		"three_asset_pairs_in_one_source":      {},
		"two_asset_pairs_in_different_sources": {},
		"small_cycle_of_asset_pairs":           {},
		"many_cycles_and_many_asset_pairs":     {},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			graphBuilder := asset_graph_builder.NewAssetGraphBuilder()

			actualGraph, actualErr := graphBuilder.BuildGraphFromAssets(ctx, test.inputAssetsPair)
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

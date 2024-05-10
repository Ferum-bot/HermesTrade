package cycles_searcher_test

import (
	"context"
	cycles_searcher "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/algoritms/cycles-searcher"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	graph_builders "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-builders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultAlgorithm_SearchAllCycles(t *testing.T) {
	type testCase struct {
		inputGraph     model.Graph
		expectedCycles []model.GraphCycle
		expectedErr    error
	}

	tests := map[string]testCase{
		"empty_graph": {
			inputGraph: model.Graph{
				Vertexes: []model.GraphVertex{},
			},
			expectedCycles: []model.GraphCycle{},
		},
		"graph_with_one_vertex": {
			inputGraph: model.Graph{
				Vertexes: []model.GraphVertex{
					{
						Identifier: model.GraphVertexIdentifier(1),
					},
				},
			},
			expectedCycles: []model.GraphCycle{},
		},
		"graph_with_two_vertex_without_edge": {
			inputGraph: model.Graph{
				Vertexes: []model.GraphVertex{
					{
						Identifier: model.GraphVertexIdentifier(1),
					},
					{
						Identifier: model.GraphVertexIdentifier(2),
					},
				},
			},
			expectedCycles: []model.GraphCycle{},
		},
		"graph_with_two_vertex_with_edge": {
			inputGraph:     buildGraphWithTwoVertexWithEdge(),
			expectedCycles: []model.GraphCycle{},
		},
		"graph_without_cycles": {
			inputGraph:     buildGraphWithoutCycles(),
			expectedCycles: []model.GraphCycle{},
		},
		"forest_graph_without_cycles": {
			inputGraph:     buildForestGraphWithoutCycles(),
			expectedCycles: []model.GraphCycle{},
		},
		"simple_graph_with_one_cycle": {
			inputGraph: buildSimpleGraphWithOneCycle(),
			expectedCycles: []model.GraphCycle{
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(0),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(0),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(0),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(0),
						},
					},
				},
			},
		},
		"forest_graph_with_two_cycles": {
			inputGraph: buildForestGraphWithTwoCycles(),
			expectedCycles: []model.GraphCycle{
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(5),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(5),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(5),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(5),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(6),
							TargetVertex: graph_builders.BuildEmptyVertex(7),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(7),
							TargetVertex: graph_builders.BuildEmptyVertex(8),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(8),
							TargetVertex: graph_builders.BuildEmptyVertex(9),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(9),
							TargetVertex: graph_builders.BuildEmptyVertex(10),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(10),
							TargetVertex: graph_builders.BuildEmptyVertex(6),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(6),
							TargetVertex: graph_builders.BuildEmptyVertex(10),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(10),
							TargetVertex: graph_builders.BuildEmptyVertex(9),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(9),
							TargetVertex: graph_builders.BuildEmptyVertex(8),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(8),
							TargetVertex: graph_builders.BuildEmptyVertex(7),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(7),
							TargetVertex: graph_builders.BuildEmptyVertex(6),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(11),
							TargetVertex: graph_builders.BuildEmptyVertex(12),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(12),
							TargetVertex: graph_builders.BuildEmptyVertex(13),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(13),
							TargetVertex: graph_builders.BuildEmptyVertex(14),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(14),
							TargetVertex: graph_builders.BuildEmptyVertex(15),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(15),
							TargetVertex: graph_builders.BuildEmptyVertex(11),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(11),
							TargetVertex: graph_builders.BuildEmptyVertex(15),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(15),
							TargetVertex: graph_builders.BuildEmptyVertex(14),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(14),
							TargetVertex: graph_builders.BuildEmptyVertex(13),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(13),
							TargetVertex: graph_builders.BuildEmptyVertex(12),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(12),
							TargetVertex: graph_builders.BuildEmptyVertex(11),
						},
					},
				},
			},
		},
		"medium_graph_with_three_cycles": {
			inputGraph: buildMediumGraphWithThreeCycles(),
			expectedCycles: []model.GraphCycle{
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(8),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(8),
							TargetVertex: graph_builders.BuildEmptyVertex(7),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(7),
							TargetVertex: graph_builders.BuildEmptyVertex(5),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(5),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(8),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(8),
							TargetVertex: graph_builders.BuildEmptyVertex(7),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(7),
							TargetVertex: graph_builders.BuildEmptyVertex(5),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(5),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(5),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(5),
							TargetVertex: graph_builders.BuildEmptyVertex(7),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(7),
							TargetVertex: graph_builders.BuildEmptyVertex(8),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(8),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(3),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(3),
							TargetVertex: graph_builders.BuildEmptyVertex(2),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(2),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: graph_builders.BuildEmptyVertex(1),
							TargetVertex: graph_builders.BuildEmptyVertex(5),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(5),
							TargetVertex: graph_builders.BuildEmptyVertex(7),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(7),
							TargetVertex: graph_builders.BuildEmptyVertex(8),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(8),
							TargetVertex: graph_builders.BuildEmptyVertex(4),
						},
						{
							SourceVertex: graph_builders.BuildEmptyVertex(4),
							TargetVertex: graph_builders.BuildEmptyVertex(1),
						},
					},
				},
			},
		},
		"big_forest_graph_with_many_cycles": {},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			cyclesSearcher := cycles_searcher.NewDefaultAlgorithm()

			actualCycles, actualErr := cyclesSearcher.SearchAllCycles(ctx, test.inputGraph)
			if test.expectedErr != nil {
				require.Error(t, actualErr)
				require.Contains(t, actualErr, test.expectedErr)
			}

			require.NoError(t, actualErr)
			require.Equal(t, len(test.expectedCycles), len(actualCycles))
			for _, actualCycle := range actualCycles {
				assertContainsCycle(t, test.expectedCycles, actualCycle)
			}
		})
	}
}

func TestDefaultAlgorithm_SearchAllCycles_BigForestGraphWithManyCycles(t *testing.T) {
	inputGraph := buildBigForestGraphWithManyCycles()
	expectedGraphCount := 72

	ctx := context.Background()
	cyclesSearcher := cycles_searcher.NewDefaultAlgorithm()

	actualCycles, actualErr := cyclesSearcher.SearchAllCycles(ctx, inputGraph)

	assert.NoError(t, actualErr)
	assert.Equal(t, expectedGraphCount, len(actualCycles))
}

func assertContainsCycle(t *testing.T, cycles []model.GraphCycle, targetCycle model.GraphCycle) {
	for _, cycle := range cycles {
		if cyclesIsEqual(cycle, targetCycle) {
			return
		}
	}

	t.Fatalf("failed to find cycle %v", targetCycle)
}

func cyclesIsEqual(firstCycle model.GraphCycle, secondCycle model.GraphCycle) bool {
	if len(firstCycle.Edges) != len(secondCycle.Edges) {
		return false
	}

	for _, edge := range firstCycle.Edges {
		if !cycleContainsEdge(secondCycle, edge) {
			return false
		}
	}

	return true
}

func cycleContainsEdge(cycle model.GraphCycle, targetEdge model.Edge) bool {
	for _, edge := range cycle.Edges {

		if edge.SourceVertex.Identifier == targetEdge.SourceVertex.Identifier &&
			edge.TargetVertex.Identifier == targetEdge.TargetVertex.Identifier {
			return true
		}
	}

	return false
}

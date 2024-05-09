package cycles_searcher_test

import (
	"context"
	cycles_searcher "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/algoritms/cycles-searcher"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
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
							SourceVertex: buildEmptyVertex(0),
							TargetVertex: buildEmptyVertex(1),
						},
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(0),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(0),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(1),
						},
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(0),
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
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(5),
						},
						{
							SourceVertex: buildEmptyVertex(5),
							TargetVertex: buildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(5),
						},
						{
							SourceVertex: buildEmptyVertex(5),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(6),
							TargetVertex: buildEmptyVertex(7),
						},
						{
							SourceVertex: buildEmptyVertex(7),
							TargetVertex: buildEmptyVertex(8),
						},
						{
							SourceVertex: buildEmptyVertex(8),
							TargetVertex: buildEmptyVertex(9),
						},
						{
							SourceVertex: buildEmptyVertex(9),
							TargetVertex: buildEmptyVertex(10),
						},
						{
							SourceVertex: buildEmptyVertex(10),
							TargetVertex: buildEmptyVertex(6),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(6),
							TargetVertex: buildEmptyVertex(10),
						},
						{
							SourceVertex: buildEmptyVertex(10),
							TargetVertex: buildEmptyVertex(9),
						},
						{
							SourceVertex: buildEmptyVertex(9),
							TargetVertex: buildEmptyVertex(8),
						},
						{
							SourceVertex: buildEmptyVertex(8),
							TargetVertex: buildEmptyVertex(7),
						},
						{
							SourceVertex: buildEmptyVertex(7),
							TargetVertex: buildEmptyVertex(6),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(11),
							TargetVertex: buildEmptyVertex(12),
						},
						{
							SourceVertex: buildEmptyVertex(12),
							TargetVertex: buildEmptyVertex(13),
						},
						{
							SourceVertex: buildEmptyVertex(13),
							TargetVertex: buildEmptyVertex(14),
						},
						{
							SourceVertex: buildEmptyVertex(14),
							TargetVertex: buildEmptyVertex(15),
						},
						{
							SourceVertex: buildEmptyVertex(15),
							TargetVertex: buildEmptyVertex(11),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(11),
							TargetVertex: buildEmptyVertex(15),
						},
						{
							SourceVertex: buildEmptyVertex(15),
							TargetVertex: buildEmptyVertex(14),
						},
						{
							SourceVertex: buildEmptyVertex(14),
							TargetVertex: buildEmptyVertex(13),
						},
						{
							SourceVertex: buildEmptyVertex(13),
							TargetVertex: buildEmptyVertex(12),
						},
						{
							SourceVertex: buildEmptyVertex(12),
							TargetVertex: buildEmptyVertex(11),
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
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(8),
						},
						{
							SourceVertex: buildEmptyVertex(8),
							TargetVertex: buildEmptyVertex(7),
						},
						{
							SourceVertex: buildEmptyVertex(7),
							TargetVertex: buildEmptyVertex(5),
						},
						{
							SourceVertex: buildEmptyVertex(5),
							TargetVertex: buildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(8),
						},
						{
							SourceVertex: buildEmptyVertex(8),
							TargetVertex: buildEmptyVertex(7),
						},
						{
							SourceVertex: buildEmptyVertex(7),
							TargetVertex: buildEmptyVertex(5),
						},
						{
							SourceVertex: buildEmptyVertex(5),
							TargetVertex: buildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(5),
						},
						{
							SourceVertex: buildEmptyVertex(5),
							TargetVertex: buildEmptyVertex(7),
						},
						{
							SourceVertex: buildEmptyVertex(7),
							TargetVertex: buildEmptyVertex(8),
						},
						{
							SourceVertex: buildEmptyVertex(8),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(3),
						},
						{
							SourceVertex: buildEmptyVertex(3),
							TargetVertex: buildEmptyVertex(2),
						},
						{
							SourceVertex: buildEmptyVertex(2),
							TargetVertex: buildEmptyVertex(1),
						},
					},
				},
				{
					Edges: []model.Edge{
						{
							SourceVertex: buildEmptyVertex(1),
							TargetVertex: buildEmptyVertex(5),
						},
						{
							SourceVertex: buildEmptyVertex(5),
							TargetVertex: buildEmptyVertex(7),
						},
						{
							SourceVertex: buildEmptyVertex(7),
							TargetVertex: buildEmptyVertex(8),
						},
						{
							SourceVertex: buildEmptyVertex(8),
							TargetVertex: buildEmptyVertex(4),
						},
						{
							SourceVertex: buildEmptyVertex(4),
							TargetVertex: buildEmptyVertex(1),
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

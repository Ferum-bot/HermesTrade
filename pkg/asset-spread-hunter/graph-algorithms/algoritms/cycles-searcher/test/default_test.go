package cycles_searcher_test

import (
	"context"
	cycles_searcher "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/algoritms/cycles-searcher"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
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
		"forest_graph_without_cycles":       {},
		"simple_graph_with_one_cycle":       {},
		"forest_graph_with_two_cycles":      {},
		"medium_graph_with_three_cycles":    {},
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

		if edge.SourceVertex == targetEdge.SourceVertex &&
			edge.TargetVertex == targetEdge.TargetVertex &&
			edge.Weight == targetEdge.Weight {

			return true
		}
	}

	return false
}

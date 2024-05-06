package graph_algorithms

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
)

type GraphCyclesSearcher interface {
	SearchAllCycles(
		ctx context.Context,
		graph model.Graph,
	) ([]model.GraphCycle, error)
}

type BFS interface {
	Run(
		ctx context.Context,
		graph model.Graph,
		action OnVertexAction,
	) error
}

type DFS interface {
	Run(
		ctx context.Context,
		startVertex model.GraphVertex,
		graph model.Graph,
		action OnVertexManagedAction,
	) error
}

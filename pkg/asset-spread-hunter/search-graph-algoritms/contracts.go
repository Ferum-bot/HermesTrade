package search_graph_algoritms

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/search-graph-algoritms/model"
)

type GraphCyclesSearcher interface {
	SearchAllCycles(
		ctx context.Context,
		graph model.Graph,
	) ([]model.GraphCycle, error)
}

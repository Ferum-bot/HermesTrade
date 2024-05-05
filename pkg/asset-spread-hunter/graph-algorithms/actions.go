package graph_algorithms

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
)

type OnVertexAction interface {
	OnVertex(ctx context.Context, vertex model.GraphVertex, graph model.Graph)
}

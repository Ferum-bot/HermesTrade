package dfs

import (
	graph_algorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
)

type dfsContext struct {
	graph  model.Graph
	action graph_algorithms.OnVertexManagedAction
}

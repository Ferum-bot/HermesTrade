package cycles_searcher

import (
	collection_algorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/collection-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
)

type vertexStatus int32

const (
	statusNotVisited vertexStatus = iota
	statusInProgress
	statusHandled
)

type searchContext struct {
	graph            model.Graph
	vertexStatuses   map[model.GraphVertexIdentifier]vertexStatus
	foundCycles      []model.GraphCycle
	currentEdgeChain collection_algorithms.CopyableStack[model.Edge]
}

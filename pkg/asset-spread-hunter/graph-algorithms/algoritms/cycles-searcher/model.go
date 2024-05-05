package cycles_searcher

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"

type vertexStatus int32

const (
	StatusNotVisited vertexStatus = iota
	StatusInProgress
	StatusHandled
)

type searchContext struct {
	vertexStatuses map[model.GraphVertexIdentifier]vertexStatus
	foundCycles    []model.GraphCycle
}

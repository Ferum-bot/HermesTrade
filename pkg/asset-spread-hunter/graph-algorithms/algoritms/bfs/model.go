package bfs

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"

type bfsContext struct {
	visitedVertexes map[model.GraphVertexIdentifier]bool
}

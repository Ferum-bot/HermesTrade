package bfs

import (
	"context"
	graph_algorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
)

type defaultBFS struct {
	context bfsContext
}

func NewDefaultBFS() graph_algorithms.BFS {
	return &defaultBFS{}
}

func (d defaultBFS) Run(
	ctx context.Context,
	graph model.Graph,
	action graph_algorithms.OnVertexAction,
) error {
	//TODO implement me
	panic("implement me")
}

package bfs

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/collection-algorithms/queue"
	graphalgorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type defaultBFS struct {
	context *bfsContext
}

func NewDefaultBFS() graphalgorithms.BFS {
	return &defaultBFS{
		context: &bfsContext{},
	}
}

func (d *defaultBFS) Run(
	ctx context.Context,
	graph model.Graph,
	action graphalgorithms.OnVertexAction,
) error {
	d.initContext(graph)

	currentVertexes := queue.NewDefaultQueue[model.GraphVertex]()
	for !currentVertexes.IsEmpty() {
		vertex, err := currentVertexes.Pop()
		if err != nil {
			return errors.Wrap(err, "currentVertexes.Pop()")
		}

		d.markVertexIsVisited(vertex.Identifier)
		action.OnVertex(ctx, *vertex, graph)

		for _, edge := range vertex.Edges {
			targetVertex := edge.TargetVertex
			if targetVertex != nil && !d.isVertexVisited(targetVertex.Identifier) {
				currentVertexes.Push(*targetVertex)
			}
		}
	}

	d.clearContext()
	return nil
}

func (d *defaultBFS) initContext(graph model.Graph) {
	d.context.vertexesQueue = make([]model.GraphVertex, 0, len(graph.Vertexes))

	d.context.visitedVertexes = make(map[model.GraphVertexIdentifier]bool, len(graph.Vertexes))
	for _, vertex := range graph.Vertexes {
		d.context.visitedVertexes[vertex.Identifier] = false
	}
}

func (d *defaultBFS) isVertexVisited(vertex model.GraphVertexIdentifier) bool {
	return d.context.visitedVertexes[vertex]
}

func (d *defaultBFS) markVertexIsVisited(vertex model.GraphVertexIdentifier) {
	d.context.visitedVertexes[vertex] = true
}

func (d *defaultBFS) clearContext() {
	d.context = &bfsContext{}
}

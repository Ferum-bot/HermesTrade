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

func (algorithm *defaultBFS) Run(
	ctx context.Context,
	graph model.Graph,
	action graphalgorithms.OnVertexAction,
) error {
	algorithm.initContext(graph)
	defer algorithm.clearContext()

	currentVertexes := queue.NewDefaultQueue[model.GraphVertex]()
	if len(graph.Vertexes) > 0 {
		currentVertexes.Push(graph.Vertexes[0])
	}

	for !currentVertexes.IsEmpty() {
		vertex, err := currentVertexes.Pop()
		if err != nil {
			return errors.Wrap(err, "currentVertexes.Pop()")
		}

		algorithm.markVertexIsVisited(vertex.Identifier)
		action.OnVertex(ctx, *vertex, graph)

		for _, edge := range vertex.Edges {
			targetVertex := edge.TargetVertex
			if targetVertex != nil && !algorithm.isVertexVisited(targetVertex.Identifier) {
				currentVertexes.Push(*targetVertex)
			}
		}
	}

	return nil
}

func (algorithm *defaultBFS) initContext(graph model.Graph) {
	algorithm.context.visitedVertexes = make(map[model.GraphVertexIdentifier]bool, len(graph.Vertexes))
	for _, vertex := range graph.Vertexes {
		algorithm.context.visitedVertexes[vertex.Identifier] = false
	}
}

func (algorithm *defaultBFS) isVertexVisited(vertex model.GraphVertexIdentifier) bool {
	return algorithm.context.visitedVertexes[vertex]
}

func (algorithm *defaultBFS) markVertexIsVisited(vertex model.GraphVertexIdentifier) {
	algorithm.context.visitedVertexes[vertex] = true
}

func (algorithm *defaultBFS) clearContext() {
	algorithm.context = &bfsContext{}
}

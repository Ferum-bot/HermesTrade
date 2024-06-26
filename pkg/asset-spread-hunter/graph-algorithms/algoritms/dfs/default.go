package dfs

import (
	"context"
	"fmt"
	graph_algorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type defaultDFS struct {
	context *dfsContext
}

func NewDefaultDFS() graph_algorithms.DFS {
	return &defaultDFS{
		context: &dfsContext{},
	}
}

func (algorithm *defaultDFS) Run(
	ctx context.Context,
	startVertex model.GraphVertex,
	graph model.Graph,
	action graph_algorithms.OnVertexManagedAction,
) error {
	algorithm.context.action = action
	algorithm.context.graph = graph

	err := algorithm.dfs(ctx, startVertex, nil)
	if err != nil {
		return errors.Wrap(err, "algorithm.dfs()")
	}

	return nil
}

func (algorithm *defaultDFS) dfs(
	ctx context.Context,
	currentVertex model.GraphVertex,
	sourceEdge *model.Edge,
) error {
	manageType := algorithm.beforeVertexManager(ctx, currentVertex, sourceEdge)

	if manageType == graph_algorithms.NotVisitChildrenManageType {
		algorithm.onVertexExit(ctx, currentVertex, sourceEdge)
		return nil
	}

	for _, edge := range currentVertex.Edges {
		targetVertex := *edge.TargetVertex

		if targetVertex.Identifier == currentVertex.Identifier {
			continue
		}

		if sourceEdge != nil && sourceEdge.SourceVertex.Identifier == targetVertex.Identifier {
			continue
		}

		currentEdge := model.Edge{
			SourceVertex: &currentVertex,
			TargetVertex: edge.TargetVertex,
			Weight:       edge.Weight,
		}

		err := algorithm.dfs(ctx, targetVertex, &currentEdge)
		if err != nil {
			errMessage := fmt.Sprintf(
				"recursive algorithm.dfs() on edge: %algorithm --> %algorithm",
				currentVertex.Identifier,
				targetVertex.Identifier,
			)
			return errors.Wrap(err, errMessage)
		}
	}

	algorithm.afterVertexManaged(ctx, currentVertex, sourceEdge)

	return nil
}

func (algorithm *defaultDFS) beforeVertexManager(
	ctx context.Context,
	currentVertex model.GraphVertex,
	sourceEdge *model.Edge,
) graph_algorithms.VertexManageType {
	return algorithm.context.action.BeforeVertexManaged(
		ctx, currentVertex, sourceEdge, algorithm.context.graph,
	)
}

func (algorithm *defaultDFS) afterVertexManaged(
	ctx context.Context,
	currentVertex model.GraphVertex,
	sourceEdge *model.Edge,
) {
	algorithm.context.action.AfterVertexManaged(
		ctx, currentVertex, sourceEdge, algorithm.context.graph,
	)
}

func (algorithm *defaultDFS) onVertexExit(
	ctx context.Context,
	currentVertex model.GraphVertex,
	sourceEdge *model.Edge,
) {
	algorithm.context.action.OnVertexExit(
		ctx, currentVertex, sourceEdge, algorithm.context.graph,
	)
}

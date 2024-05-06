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

func (d *defaultDFS) Run(
	ctx context.Context,
	graph model.Graph,
	action graph_algorithms.OnVertexManagedAction,
) error {
	rootVertex := model.GraphVertex{}
	if len(graph.Vertexes) > 0 {
		rootVertex = graph.Vertexes[0]
	} else {
		return nil
	}

	d.context.action = action
	d.context.graph = graph

	err := d.dfs(ctx, rootVertex, nil)
	if err != nil {
		return errors.Wrap(err, "d.dfs()")
	}

	return nil
}

func (d *defaultDFS) dfs(
	ctx context.Context,
	currentVertex model.GraphVertex,
	sourceEdge *model.Edge,
) error {
	manageType := d.context.action.OnVertexManaged(
		ctx, currentVertex, sourceEdge, d.context.graph,
	)

	if manageType == graph_algorithms.NotVisitChildrenManageType {
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

		err := d.dfs(ctx, targetVertex, sourceEdge)
		if err != nil {
			errMessage := fmt.Sprintf(
				"recursive d.dfs() on edge: %d --> %d",
				currentVertex.Identifier,
				targetVertex.Identifier,
			)
			return errors.Wrap(err, errMessage)
		}
	}

	return nil
}

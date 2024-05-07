package cycles_searcher

import (
	"context"
	collection_algorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/collection-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/collection-algorithms/stack"
	graphalgorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	dfs2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/algoritms/dfs"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type defaultAlgorithm struct {
	context *searchContext
}

func NewDefaultAlgorithm() graphalgorithms.GraphCyclesSearcher {
	return &defaultAlgorithm{
		context: &searchContext{},
	}
}

func (algorithm *defaultAlgorithm) SearchAllCycles(
	ctx context.Context,
	graph model.Graph,
) ([]model.GraphCycle, error) {
	algorithm.initContext(graph)
	defer algorithm.clearContext()

	for _, vertex := range graph.Vertexes {
		if algorithm.isVertexNotVisited(vertex) {
			err := algorithm.searchCycles(ctx, vertex)
			if err != nil {
				return nil, errors.Wrap(err, "algorithm.searchCycles()")
			}
		}
	}

	return algorithm.context.foundCycles, nil
}

func (algorithm *defaultAlgorithm) searchCycles(
	ctx context.Context,
	vertex model.GraphVertex,
) error {
	dfs := dfs2.NewDefaultDFS()

	err := dfs.Run(ctx, vertex, algorithm.context.graph, algorithm)
	if err != nil {
		return errors.Wrap(err, "dfs.Run()")
	}

	return nil
}

func (algorithm *defaultAlgorithm) BeforeVertexManaged(
	ctx context.Context,
	targetVertex model.GraphVertex,
	edge *model.Edge,
	graph model.Graph,
) graphalgorithms.VertexManageType {
	if edge != nil {
		algorithm.context.currentEdgeChain.Push(*edge)
	}

	if algorithm.isVertexInProgress(targetVertex) {
		algorithm.addCycle(targetVertex)
		return graphalgorithms.NotVisitChildrenManageType
	} else {
		algorithm.markVertexInProgress(targetVertex)
	}

	return graphalgorithms.VisitChildrenManageType
}

func (algorithm *defaultAlgorithm) AfterVertexManaged(
	ctx context.Context,
	targetVertex model.GraphVertex,
	edge *model.Edge,
	graph model.Graph,
) {
	if !algorithm.context.currentEdgeChain.IsEmpty() {
		_, _ = algorithm.context.currentEdgeChain.Pop()
	}

	algorithm.markVertexHandled(targetVertex)
}

func (algorithm *defaultAlgorithm) isVertexInProgress(vertex model.GraphVertex) bool {
	return algorithm.context.vertexStatuses[vertex.Identifier] == statusInProgress
}

func (algorithm *defaultAlgorithm) isVertexNotVisited(vertex model.GraphVertex) bool {
	return algorithm.context.vertexStatuses[vertex.Identifier] == statusNotVisited
}

func (algorithm *defaultAlgorithm) markVertexNotVisited(vertex model.GraphVertex) {
	algorithm.context.vertexStatuses[vertex.Identifier] = statusNotVisited
}

func (algorithm *defaultAlgorithm) markVertexInProgress(vertex model.GraphVertex) {
	algorithm.context.vertexStatuses[vertex.Identifier] = statusInProgress
}

func (algorithm *defaultAlgorithm) markVertexHandled(vertex model.GraphVertex) {
	algorithm.context.vertexStatuses[vertex.Identifier] = statusHandled
}

func (algorithm *defaultAlgorithm) addCycle(targetVertex model.GraphVertex) {
	allCurrentEdges := algorithm.context.currentEdgeChain.MakeCopy().(collection_algorithms.CopyableStack[model.Edge])

	currentCycleEdges := make([]model.Edge, 0)

	for !allCurrentEdges.IsEmpty() {
		edge, _ := allCurrentEdges.Pop()
		currentCycleEdges = append(currentCycleEdges, *edge)

		if edge.SourceVertex.Identifier == targetVertex.Identifier {
			break
		}
	}

	cycle := model.GraphCycle{
		Edges: currentCycleEdges,
	}
	algorithm.context.foundCycles = append(algorithm.context.foundCycles, cycle)
}

func (algorithm *defaultAlgorithm) initContext(graph model.Graph) {
	algorithm.context.graph = graph
	algorithm.context.foundCycles = make([]model.GraphCycle, 0)
	algorithm.context.currentEdgeChain = stack.NewDefaultStack[model.Edge]()

	algorithm.context.vertexStatuses = make(map[model.GraphVertexIdentifier]vertexStatus, len(graph.Vertexes))
	for _, vertex := range graph.Vertexes {
		algorithm.markVertexNotVisited(vertex)
	}
}

func (algorithm *defaultAlgorithm) clearContext() {
	algorithm.context = &searchContext{}
}

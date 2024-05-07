package cycles_searcher_test

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"

func buildGraphWithTwoVertexWithEdge() model.Graph {
	firstVertex := model.GraphVertex{
		Identifier: model.GraphVertexIdentifier(1),
	}
	secondVertex := model.GraphVertex{
		Identifier: model.GraphVertexIdentifier(2),
	}

	firstVertex.Edges = make([]model.VertexEdge, 0)
	secondVertex.Edges = make([]model.VertexEdge, 0)

	firstVertex.Edges = append(firstVertex.Edges, model.VertexEdge{
		TargetVertex: &secondVertex,
		Weight:       model.EdgeWeight(1),
	})
	secondVertex.Edges = append(secondVertex.Edges, model.VertexEdge{
		TargetVertex: &firstVertex,
		Weight:       model.EdgeWeight(1),
	})

	return model.Graph{
		Vertexes: []model.GraphVertex{firstVertex, secondVertex},
	}
}

func buildGraphWithoutCycles() model.Graph {
	const treeDepth = 5
	const edgesPerVertexCount = 3

	identifierCounter := int64(1)
	allVertexes := make([]*model.GraphVertex, 0, treeDepth*edgesPerVertexCount)

	var treeBuilder func(currentDepth int, sourceVertex *model.GraphVertex)
	treeBuilder = func(currentDepth int, sourceVertex *model.GraphVertex) {
		if currentDepth > treeDepth {
			return
		}

		targetVertex := model.GraphVertex{
			Identifier: model.GraphVertexIdentifier(identifierCounter),
			Edges:      make([]model.VertexEdge, 0),
		}
		identifierCounter++

		allVertexes = append(allVertexes, &targetVertex)
		sourceVertex.Edges = append(sourceVertex.Edges, model.VertexEdge{
			TargetVertex: &targetVertex,
			Weight:       model.EdgeWeight(1),
		})
		targetVertex.Edges = append(targetVertex.Edges, model.VertexEdge{
			TargetVertex: sourceVertex,
			Weight:       model.EdgeWeight(1),
		})

		for i := 0; i < edgesPerVertexCount; i++ {
			treeBuilder(currentDepth+1, &targetVertex)
		}
	}

	rootVertex := model.GraphVertex{
		Identifier: 0,
		Edges:      make([]model.VertexEdge, 0),
	}
	allVertexes = append(allVertexes, &rootVertex)
	treeBuilder(1, &rootVertex)

	graph := model.Graph{Vertexes: make([]model.GraphVertex, 0, len(allVertexes))}
	for _, vertex := range allVertexes {
		graph.Vertexes = append(graph.Vertexes, *vertex)
	}

	return graph
}

func buildForestGraphWithoutCycles() model.Graph {
	const treeDepth = 3
	const edgesPerVertexCount = 5
	const forestCount = 4

	identifierCounter := int64(0)
	allVertexes := make([]*model.GraphVertex, 0, treeDepth*edgesPerVertexCount)

	var treeBuilder func(currentDepth int, sourceVertex *model.GraphVertex)
	treeBuilder = func(currentDepth int, sourceVertex *model.GraphVertex) {
		if currentDepth > treeDepth {
			return
		}

		targetVertex := model.GraphVertex{
			Identifier: model.GraphVertexIdentifier(identifierCounter),
			Edges:      make([]model.VertexEdge, 0),
		}
		identifierCounter++

		allVertexes = append(allVertexes, &targetVertex)
		sourceVertex.Edges = append(sourceVertex.Edges, model.VertexEdge{
			TargetVertex: &targetVertex,
			Weight:       model.EdgeWeight(1),
		})
		targetVertex.Edges = append(targetVertex.Edges, model.VertexEdge{
			TargetVertex: sourceVertex,
			Weight:       model.EdgeWeight(1),
		})

		for i := 0; i < edgesPerVertexCount; i++ {
			treeBuilder(currentDepth+1, &targetVertex)
		}
	}

	for i := 0; i < forestCount; i++ {
		forestRoot := model.GraphVertex{Identifier: model.GraphVertexIdentifier(identifierCounter)}
		identifierCounter++

		allVertexes = append(allVertexes, &forestRoot)
		treeBuilder(1, &forestRoot)
	}

	graph := model.Graph{Vertexes: make([]model.GraphVertex, 0, len(allVertexes))}
	for _, vertex := range allVertexes {
		graph.Vertexes = append(graph.Vertexes, *vertex)
	}

	return graph
}

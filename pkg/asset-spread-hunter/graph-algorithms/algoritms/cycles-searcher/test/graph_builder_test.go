package cycles_searcher_test

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"

func buildGraphWithTwoVertexWithEdge() model.Graph {
	firstVertex := buildEmptyVertex(1)
	secondVertex := buildEmptyVertex(2)

	firstVertex.Edges = make([]model.VertexEdge, 0)
	secondVertex.Edges = make([]model.VertexEdge, 0)

	buildEdge(firstVertex, secondVertex)

	return model.Graph{
		Vertexes: []model.GraphVertex{*firstVertex, *secondVertex},
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

		targetVertex := buildEmptyVertex(identifierCounter)
		identifierCounter++

		allVertexes = append(allVertexes, targetVertex)

		buildEdge(sourceVertex, targetVertex)

		for i := 0; i < edgesPerVertexCount; i++ {
			treeBuilder(currentDepth+1, targetVertex)
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

		targetVertex := buildEmptyVertex(identifierCounter)
		identifierCounter++

		allVertexes = append(allVertexes, targetVertex)
		buildEdge(sourceVertex, targetVertex)

		for i := 0; i < edgesPerVertexCount; i++ {
			treeBuilder(currentDepth+1, targetVertex)
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

func buildSimpleGraphWithOneCycle() model.Graph {
	const cycleLength = 4
	const nonCycleEdgesCount = 3

	identifierCounter := int64(0)
	allVertexes := make([]*model.GraphVertex, 0, cycleLength)

	var cycleBuilder func(rootVertex *model.GraphVertex, prevVertex *model.GraphVertex, vertexNumber int)
	cycleBuilder = func(rootVertex *model.GraphVertex, prevVertex *model.GraphVertex, vertexNumber int) {
		if vertexNumber > cycleLength {
			buildEdge(rootVertex, prevVertex)

			return
		}

		identifierCounter++
		newVertex := buildEmptyVertex(identifierCounter)
		allVertexes = append(allVertexes, newVertex)

		buildEdge(prevVertex, newVertex)

		cycleBuilder(rootVertex, newVertex, vertexNumber+1)
	}

	rootVertex := buildEmptyVertex(identifierCounter)
	allVertexes = append(allVertexes, rootVertex)
	cycleBuilder(rootVertex, rootVertex, 1)

	for i := range allVertexes {
		for j := 0; j < nonCycleEdgesCount; j++ {
			identifierCounter++
			newVertex := buildEmptyVertex(identifierCounter)
			allVertexes = append(allVertexes, newVertex)
			cycleVertex := allVertexes[i]

			buildEdge(cycleVertex, newVertex)
		}
	}

	graph := model.Graph{Vertexes: make([]model.GraphVertex, 0, len(allVertexes))}
	for _, vertex := range allVertexes {
		graph.Vertexes = append(graph.Vertexes, *vertex)
	}

	return graph
}

func buildForestGraphWithTwoCycles() model.Graph {
	const cycleLength = 4
	const nonCycleEdgesCount = 1
	const forestCount = 3

	identifierCounter := int64(0)
	allVertexes := make([]*model.GraphVertex, 0, cycleLength)

	var cycleBuilder func(rootVertex *model.GraphVertex, prevVertex *model.GraphVertex, vertexNumber int)
	cycleBuilder = func(rootVertex *model.GraphVertex, prevVertex *model.GraphVertex, vertexNumber int) {
		if vertexNumber > cycleLength {
			buildEdge(rootVertex, prevVertex)
			return
		}

		identifierCounter++
		newVertex := buildEmptyVertex(identifierCounter)
		allVertexes = append(allVertexes, newVertex)

		buildEdge(prevVertex, newVertex)

		cycleBuilder(rootVertex, newVertex, vertexNumber+1)
	}

	for i := 0; i < forestCount; i++ {
		identifierCounter++
		rootVertex := buildEmptyVertex(identifierCounter)
		allVertexes = append(allVertexes, rootVertex)

		cycleBuilder(rootVertex, rootVertex, 1)
	}

	for i := range allVertexes {
		for j := 0; j < nonCycleEdgesCount; j++ {
			identifierCounter++
			newVertex := buildEmptyVertex(identifierCounter)
			allVertexes = append(allVertexes, newVertex)
			cycleVertex := allVertexes[i]

			buildEdge(cycleVertex, newVertex)
		}
	}

	graph := model.Graph{Vertexes: make([]model.GraphVertex, 0, len(allVertexes))}
	for _, vertex := range allVertexes {
		graph.Vertexes = append(graph.Vertexes, *vertex)
	}

	return graph
}

func buildMediumGraphWithThreeCycles() model.Graph {
	const vertexCount = 11

	identifierCounter := int64(0)
	allVertexes := make([]*model.GraphVertex, 0, vertexCount)

	for i := 0; i < vertexCount; i++ {
		vertex := buildEmptyVertex(identifierCounter)
		identifierCounter++

		allVertexes = append(allVertexes, vertex)
	}

	buildEdge(allVertexes[0], allVertexes[1])

	buildEdge(allVertexes[1], allVertexes[2])
	buildEdge(allVertexes[2], allVertexes[3])
	buildEdge(allVertexes[3], allVertexes[4])
	buildEdge(allVertexes[4], allVertexes[1])

	buildEdge(allVertexes[3], allVertexes[9])
	buildEdge(allVertexes[9], allVertexes[10])

	buildEdge(allVertexes[4], allVertexes[8])
	buildEdge(allVertexes[8], allVertexes[7])
	buildEdge(allVertexes[7], allVertexes[5])
	buildEdge(allVertexes[5], allVertexes[1])

	buildEdge(allVertexes[5], allVertexes[6])

	graph := model.Graph{Vertexes: make([]model.GraphVertex, 0, len(allVertexes))}
	for _, vertex := range allVertexes {
		graph.Vertexes = append(graph.Vertexes, *vertex)
	}

	return graph
}

func buildEmptyVertex(identifier int64) *model.GraphVertex {
	return &model.GraphVertex{
		Identifier: model.GraphVertexIdentifier(identifier),
		Edges:      make([]model.VertexEdge, 0),
	}
}

func buildEdge(sourceVertex *model.GraphVertex, targetVertex *model.GraphVertex) {
	sourceVertex.Edges = append(sourceVertex.Edges, model.VertexEdge{
		TargetVertex: targetVertex,
		Weight:       model.EdgeWeight(1),
	})
	targetVertex.Edges = append(targetVertex.Edges, model.VertexEdge{
		TargetVertex: sourceVertex,
		Weight:       model.EdgeWeight(1),
	})
}

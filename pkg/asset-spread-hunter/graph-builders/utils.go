package graph_builders

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"

func BuildEmptyVertex(identifier int64) *model.GraphVertex {
	return &model.GraphVertex{
		Identifier: model.GraphVertexIdentifier(identifier),
		Edges:      make([]model.VertexEdge, 0),
	}
}

func BuildEdge(sourceVertex *model.GraphVertex, targetVertex *model.GraphVertex) {
	sourceVertex.Edges = append(sourceVertex.Edges, model.VertexEdge{
		TargetVertex: targetVertex,
		Weight:       model.EdgeWeight(1),
	})
	targetVertex.Edges = append(targetVertex.Edges, model.VertexEdge{
		TargetVertex: sourceVertex,
		Weight:       model.EdgeWeight(1),
	})
}

func BuildEdgeWithWeight(
	sourceVertex *model.GraphVertex,
	targetVertex *model.GraphVertex,
	weight model.EdgeWeight,
) {
	sourceVertex.Edges = append(sourceVertex.Edges, model.VertexEdge{
		TargetVertex: targetVertex,
		Weight:       weight,
	})
}

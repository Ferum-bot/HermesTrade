package model

type GraphVertexIdentifier int64
type EdgeWeight int64

type GraphVertex struct {
	Identifier GraphVertexIdentifier
	Edges      []VertexEdge
}

type VertexEdge struct {
	TargetVertex *GraphVertex
	Weight       EdgeWeight
}

type Edge struct {
	SourceVertex *GraphVertex
	TargetVertex *GraphVertex
	Weight       EdgeWeight
}

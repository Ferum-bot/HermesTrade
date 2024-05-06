package graph_algorithms

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
)

type VertexManageType bool

const (
	VisitChildrenManageType    VertexManageType = true
	NotVisitChildrenManageType VertexManageType = false
)

type OnVertexAction interface {
	OnVertex(
		ctx context.Context,
		vertex model.GraphVertex,
		graph model.Graph,
	)
}

type OnVertexManagedAction interface {
	BeforeVertexManaged(
		ctx context.Context,
		targetVertex model.GraphVertex,
		edge *model.Edge,
		graph model.Graph,
	) VertexManageType

	AfterVertexManaged(
		ctx context.Context,
		targetVertex model.GraphVertex,
		edge *model.Edge,
		graph model.Graph,
	)
}

package model

type Graph struct {
	Vertexes []GraphVertex
}

type GraphCycle struct {
	Edges       []Edge
	CycleLength int64
}

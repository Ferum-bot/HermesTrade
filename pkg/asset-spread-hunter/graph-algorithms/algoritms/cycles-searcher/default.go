package cycles_searcher

import (
	"context"
	"errors"
	searchgraphalgoritms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
)

type defaultAlgorithm struct {
}

func NewDefaultAlgorithm() searchgraphalgoritms.GraphCyclesSearcher {
	return &defaultAlgorithm{}
}

func (algorithm defaultAlgorithm) SearchAllCycles(
	ctx context.Context,
	graph model.Graph,
) ([]model.GraphCycle, error) {
	return nil, errors.New("not implemented yet")
}

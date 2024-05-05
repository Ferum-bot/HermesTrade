package cycles_searcher

import (
	"context"
	"errors"
	graphalgorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
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
	return nil, errors.New("not implemented yet")
}

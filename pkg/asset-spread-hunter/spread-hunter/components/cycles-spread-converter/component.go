package cycles_spread_converter

import (
	"context"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type CyclesSpreadConverter struct {
}

func NewCyclesSpreadConverter() *CyclesSpreadConverter {
	return &CyclesSpreadConverter{}
}

func (c *CyclesSpreadConverter) ConvertCyclesToSpreads(
	ctx context.Context,
	cycles []model2.GraphCycle,
) ([]model.Spread, error) {
	//TODO implement me
	panic("implement me")
}

package asset_graph_builder

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type AssetGraphBuilder struct {
}

func NewAssetGraphBuilder() *AssetGraphBuilder {
	return &AssetGraphBuilder{}
}

func (graphBuilder *AssetGraphBuilder) BuildGraphFromAssets(
	ctx context.Context,
	assets []model2.AssetCurrencyPair,
) (model.Graph, error) {
	//TODO implement me
	panic("implement me")
}

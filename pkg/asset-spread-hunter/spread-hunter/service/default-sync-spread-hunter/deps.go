package default_sync_spread_hunter

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

//go:generate mockgen -source $GOFILE -destination test/service_mocks_test.go -package default_sync_spread_hunter_test

type assetGraphBuilder interface {
	BuildGraphFromAssets(
		ctx context.Context,
		assets []model2.AssetCurrencyPair,
	) (model.Graph, error)
}

type cyclesSpreadConverter interface {
	ConvertCyclesToSpreads(
		ctx context.Context,
		cycles []model.GraphCycle,
	) ([]model2.Spread, error)
}

type profitabilityComparer interface {
	ProfitabilityIsLessThan(
		source model2.SpreadProfitability,
		than model2.SpreadProfitability,
	) bool

	ProfitabilityIsGreaterThan(
		source model2.SpreadProfitability,
		than model2.SpreadProfitability,
	) bool
}

type graphCyclesSearcher interface {
	SearchAllCycles(
		ctx context.Context,
		graph model.Graph,
	) ([]model.GraphCycle, error)
}

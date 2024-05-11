package default_sync_spread_hunter

import (
	"context"
	graphalgorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms"
	cycles_searcher "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/algoritms/cycles-searcher"
	errors2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter"
	asset_graph_builder "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/components/asset-graph-builder"
	cycles_spread_converter "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/components/cycles-spread-converter"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/components/profitability"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type defaultSpreadHunter struct {
	graphBuilder          assetGraphBuilder
	cyclesConverter       cyclesSpreadConverter
	cyclesSearcher        graphCyclesSearcher
	profitabilityComparer profitabilityComparer
}

func NewDefaultSyncSpreadHunter() spread_hunter.SyncSpreadHunter {
	graphBuilder := asset_graph_builder.NewAssetGraphBuilder()
	cyclesConverter := cycles_spread_converter.NewCyclesSpreadConverter()
	cyclesSearcher := cycles_searcher.NewDefaultAlgorithm()
	comparer := profitability.NewComparer()

	return NewSyncSpreadHunter(
		graphBuilder,
		cyclesConverter,
		cyclesSearcher,
		comparer,
	)
}

func NewSyncSpreadHunter(
	graphBuilder assetGraphBuilder,
	cyclesConverter cyclesSpreadConverter,
	cyclesSearcher graphalgorithms.GraphCyclesSearcher,
	profitabilityComparer profitabilityComparer,
) spread_hunter.SyncSpreadHunter {
	return &defaultSpreadHunter{
		graphBuilder:          graphBuilder,
		cyclesConverter:       cyclesConverter,
		cyclesSearcher:        cyclesSearcher,
		profitabilityComparer: profitabilityComparer,
	}
}

func (spreadHunter *defaultSpreadHunter) SearchSpread(
	ctx context.Context,
	assetPairs []model.AssetCurrencyPair,
	searchSettings model.SpreadSearchSettings,
) ([]model.Spread, error) {
	assetsGraph, err := spreadHunter.graphBuilder.BuildGraphFromAssets(ctx, assetPairs)
	if err != nil {
		return nil, errors2.Wrap(err, "spreadHunter.graphBuilder.BuildGraphFromAssets")
	}

	foundCycles, err := spreadHunter.cyclesSearcher.SearchAllCycles(ctx, assetsGraph)
	if err != nil {
		return nil, errors2.Wrap(err, "spreadHunter.graphCyclesSearcher.SearchAllCycles")
	}

	spreads, err := spreadHunter.cyclesConverter.ConvertCyclesToSpreads(ctx, foundCycles, assetPairs)
	if err != nil {
		return nil, errors2.Wrap(err, "spreadHunter.cyclesConverter.ConvertCyclesToSpreads")
	}

	filteredSpreads := spreadHunter.filterFoundSpreadsBySettings(spreads, searchSettings)
	return filteredSpreads, nil
}

func (spreadHunter *defaultSpreadHunter) filterFoundSpreadsBySettings(
	foundSpreads []model.Spread,
	settings model.SpreadSearchSettings,
) []model.Spread {
	filteredSpreads := make([]model.Spread, 0, len(foundSpreads))

	for _, spread := range foundSpreads {
		if spreadHunter.spreadIsMatchSettings(spread, settings) {
			filteredSpreads = append(filteredSpreads, spread)
		}
	}

	return filteredSpreads
}

func (spreadHunter *defaultSpreadHunter) spreadIsMatchSettings(
	spread model.Spread,
	settings model.SpreadSearchSettings,
) bool {
	maxSpreadLength := model.SpreadLength(model.DefaultMaxSpreadLength)
	if settings.MaxSpreadLength != nil {
		maxSpreadLength = model.SpreadLength(*settings.MaxSpreadLength)
	}

	minSpreadLength := model.SpreadLength(model.DefaultMinSpreadLength)
	if settings.MinSpreadLength != nil {
		minSpreadLength = model.SpreadLength(*settings.MinSpreadLength)
	}

	minProfitability := settings.MinSearchProfitabilityRatio
	maxProfitability := settings.MaxSearchProfitabilityRatio

	if spread.MetaInformation.Length > maxSpreadLength {
		return false
	}
	if spread.MetaInformation.Length < minSpreadLength {
		return false
	}

	if maxProfitability != nil {
		profitabilityIsGrater := spreadHunter.profitabilityComparer.ProfitabilityIsGreaterThan(
			spread.MetaInformation.ProfitabilityPercent, *maxProfitability,
		)
		if profitabilityIsGrater {
			return false
		}
	}

	if minProfitability != nil {
		profitabilityIsLess := spreadHunter.profitabilityComparer.ProfitabilityIsLessThan(
			spread.MetaInformation.ProfitabilityPercent, *minProfitability,
		)
		if profitabilityIsLess {
			return false
		}
	}

	return true
}

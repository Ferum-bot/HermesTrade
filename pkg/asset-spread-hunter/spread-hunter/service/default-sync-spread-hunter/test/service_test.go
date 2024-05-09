package default_sync_spread_hunter_test

import (
	"context"
	"errors"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/pointers"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	defaultsyncspreadhunter "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/service/default-sync-spread-hunter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestDefaultSpreadHunter_SearchSpread(t *testing.T) {
	type testCase struct {
		initCyclesSearcher        func(mock *MockgraphCyclesSearcher)
		initAssetGraphBuilder     func(mock *MockassetGraphBuilder)
		initCyclesConverter       func(mock *MockcyclesSpreadConverter)
		initProfitabilityComparer func(mock *MockprofitabilityComparer)
		assetPairsIn              []model.AssetCurrencyPair
		searchSettingsIn          model.SpreadSearchSettings
		expectedSpreads           []model.Spread
		expectedErr               error
	}

	assetPairs := []model.AssetCurrencyPair{
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  1,
				SourceIdentifier:    2,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  2,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     101,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  1,
				SourceIdentifier:    1,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  2,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 2,
				Value:     111,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  1,
				SourceIdentifier:    5,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  2,
				SourceIdentifier:    5,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 2,
				Value:     112,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 5,
				ExternalIdentifier:  1,
				SourceIdentifier:    4,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 6,
				ExternalIdentifier:  2,
				SourceIdentifier:    4,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 2,
				Value:     121,
			},
		},
		{
			Identifier: model.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model.Asset{
				UniversalIdentifier: 7,
				ExternalIdentifier:  1,
				SourceIdentifier:    3,
			},
			QuotedAsset: model.Asset{
				UniversalIdentifier: 8,
				ExternalIdentifier:  2,
				SourceIdentifier:    3,
			},
			CurrencyRatio: model.AssetsCurrencyRatio{
				Precision: 1,
				Value:     211,
			},
		},
	}
	nullSearchSettings := model.SpreadSearchSettings{
		MaxSpreadLength:             nil,
		MinSpreadLength:             nil,
		MinSearchProfitabilityRatio: nil,
		MaxSearchProfitabilityRatio: nil,
	}
	maxSpreadLength := int64(10)
	minSpreadLength := int64(0)
	minProfitability := model.SpreadProfitability{
		Precision: 0,
		Value:     0,
	}
	maxProfitability := model.SpreadProfitability{
		Precision: 0,
		Value:     100000,
	}
	searchSettings := model.SpreadSearchSettings{
		MaxSpreadLength:             &maxSpreadLength,
		MinSpreadLength:             &minSpreadLength,
		MinSearchProfitabilityRatio: &minProfitability,
		MaxSearchProfitabilityRatio: &maxProfitability,
	}
	vertexes := []model2.GraphVertex{
		{
			Identifier: 1,
			Edges:      []model2.VertexEdge{},
		},
		{
			Identifier: 2,
			Edges:      []model2.VertexEdge{},
		},
		{
			Identifier: 3,
			Edges:      []model2.VertexEdge{},
		},
		{
			Identifier: 4,
			Edges:      []model2.VertexEdge{},
		},
		{
			Identifier: 5,
			Edges:      []model2.VertexEdge{},
		},
	}
	vertexes[0].Edges = append(vertexes[0].Edges, model2.VertexEdge{
		TargetVertex: &vertexes[1],
		Weight:       1,
	})
	vertexes[1].Edges = append(vertexes[1].Edges, model2.VertexEdge{
		TargetVertex: &vertexes[2],
		Weight:       1,
	})
	vertexes[2].Edges = append(vertexes[2].Edges, model2.VertexEdge{
		TargetVertex: &vertexes[3],
		Weight:       1,
	})
	vertexes[3].Edges = append(vertexes[3].Edges, model2.VertexEdge{
		TargetVertex: &vertexes[4],
		Weight:       1,
	})
	vertexes[4].Edges = append(vertexes[4].Edges, model2.VertexEdge{
		TargetVertex: &vertexes[0],
		Weight:       1,
	})
	graph := model2.Graph{
		Vertexes: vertexes,
	}
	emptyGraph := model2.Graph{
		Vertexes: []model2.GraphVertex{},
	}
	emptyAssetPairs := []model.AssetCurrencyPair{}
	cycles := []model2.GraphCycle{
		{
			Edges: []model2.Edge{
				{
					SourceVertex: &vertexes[0],
					TargetVertex: &vertexes[1],
				},
			},
		},
	}
	emptyCycles := []model2.GraphCycle{}
	spreadProfitability := model.SpreadProfitability{
		Precision: 2,
		Value:     811,
	}
	spreads := []model.Spread{
		{
			MetaInformation: model.SpreadMetaInformation{
				Length:        3,
				Profitability: minProfitability,
				CreatedAt:     time.Now(),
			},
			Identifier: model.SpreadIdentifier(uuid.New().String()),
		},
		{
			MetaInformation: model.SpreadMetaInformation{
				Length:        3,
				Profitability: maxProfitability,
				CreatedAt:     time.Now(),
			},
			Identifier: model.SpreadIdentifier(uuid.New().String()),
		},
		{
			MetaInformation: model.SpreadMetaInformation{
				Length:        0,
				Profitability: spreadProfitability,
				CreatedAt:     time.Now(),
			},
			Identifier: model.SpreadIdentifier(uuid.New().String()),
		},
		{
			MetaInformation: model.SpreadMetaInformation{
				Length:        8,
				Profitability: spreadProfitability,
				CreatedAt:     time.Now(),
			},
			Identifier: model.SpreadIdentifier(uuid.New().String()),
		},
	}

	tests := map[string]testCase{
		"success_empty_currency_pairs": {
			initAssetGraphBuilder: func(mock *MockassetGraphBuilder) {
				mock.EXPECT().
					BuildGraphFromAssets(gomock.Any(), emptyAssetPairs).
					Times(1).
					Return(emptyGraph, nil)
			},
			initCyclesSearcher: func(mock *MockgraphCyclesSearcher) {
				mock.EXPECT().
					SearchAllCycles(gomock.Any(), emptyGraph).
					Times(1).
					Return(emptyCycles, nil)
			},
			initCyclesConverter: func(mock *MockcyclesSpreadConverter) {
				mock.EXPECT().
					ConvertCyclesToSpreads(gomock.Any(), emptyCycles).
					Times(1).
					Return([]model.Spread{}, nil)
			},
			assetPairsIn:     emptyAssetPairs,
			searchSettingsIn: nullSearchSettings,
			expectedSpreads:  []model.Spread{},
		},
		"success_several_currency_pairs": {
			initAssetGraphBuilder: func(mock *MockassetGraphBuilder) {
				mock.EXPECT().
					BuildGraphFromAssets(gomock.Any(), assetPairs).
					Times(1).
					Return(graph, nil)
			},
			initCyclesSearcher: func(mock *MockgraphCyclesSearcher) {
				mock.EXPECT().
					SearchAllCycles(gomock.Any(), graph).
					Times(1).
					Return(cycles, nil)
			},
			initCyclesConverter: func(mock *MockcyclesSpreadConverter) {
				mock.EXPECT().
					ConvertCyclesToSpreads(gomock.Any(), cycles).
					Times(1).
					Return(spreads, nil)
			},
			initProfitabilityComparer: func(mock *MockprofitabilityComparer) {
				mock.EXPECT().
					ProfitabilityIsLessThan(spreadProfitability, minProfitability).
					Times(2).
					Return(false)
				mock.EXPECT().
					ProfitabilityIsGreaterThan(spreadProfitability, maxProfitability).
					Times(2).
					Return(false)
				mock.EXPECT().
					ProfitabilityIsGreaterThan(minProfitability, maxProfitability).
					Times(1).
					Return(false)
				mock.EXPECT().
					ProfitabilityIsLessThan(minProfitability, minProfitability).
					Times(1).
					Return(false)
				mock.EXPECT().
					ProfitabilityIsGreaterThan(maxProfitability, maxProfitability).
					Times(1).
					Return(false)
				mock.EXPECT().
					ProfitabilityIsLessThan(maxProfitability, minProfitability).
					Times(1).
					Return(false)
			},
			assetPairsIn:     assetPairs,
			searchSettingsIn: searchSettings,
			expectedSpreads:  spreads,
		},
		"success_several_currency_pairs_filtered_by": {
			initAssetGraphBuilder: func(mock *MockassetGraphBuilder) {
				mock.EXPECT().
					BuildGraphFromAssets(gomock.Any(), assetPairs).
					Times(1).
					Return(graph, nil)
			},
			initCyclesSearcher: func(mock *MockgraphCyclesSearcher) {
				mock.EXPECT().
					SearchAllCycles(gomock.Any(), graph).
					Times(1).
					Return(cycles, nil)
			},
			initCyclesConverter: func(mock *MockcyclesSpreadConverter) {
				mock.EXPECT().
					ConvertCyclesToSpreads(gomock.Any(), cycles).
					Times(1).
					Return(spreads, nil)
			},
			initProfitabilityComparer: func(mock *MockprofitabilityComparer) {
				mock.EXPECT().
					ProfitabilityIsGreaterThan(minProfitability, maxProfitability).
					Times(1).
					Return(false)
				mock.EXPECT().
					ProfitabilityIsLessThan(minProfitability, minProfitability).
					Times(1).
					Return(true)
				mock.EXPECT().
					ProfitabilityIsGreaterThan(maxProfitability, maxProfitability).
					Times(1).
					Return(true)
			},
			assetPairsIn: assetPairs,
			searchSettingsIn: model.SpreadSearchSettings{
				MaxSpreadLength:             pointers.Int64Pointer(7),
				MinSpreadLength:             pointers.Int64Pointer(1),
				MinSearchProfitabilityRatio: &minProfitability,
				MaxSearchProfitabilityRatio: &maxProfitability,
			},
			expectedSpreads: []model.Spread{},
		},
		"error_graph_builder": {
			initAssetGraphBuilder: func(mock *MockassetGraphBuilder) {
				mock.EXPECT().
					BuildGraphFromAssets(gomock.Any(), assetPairs).
					Times(1).
					Return(model2.Graph{}, errors.New("something went wrong"))
			},
			searchSettingsIn: nullSearchSettings,
			assetPairsIn:     assetPairs,
			expectedErr:      errors.New("something went wrong"),
		},
		"error_found_cycles": {
			initAssetGraphBuilder: func(mock *MockassetGraphBuilder) {
				mock.EXPECT().
					BuildGraphFromAssets(gomock.Any(), assetPairs).
					Times(1).
					Return(graph, nil)
			},
			initCyclesSearcher: func(mock *MockgraphCyclesSearcher) {
				mock.EXPECT().
					SearchAllCycles(gomock.Any(), graph).
					Times(1).
					Return(nil, errors.New("something went wrong"))
			},
			searchSettingsIn: nullSearchSettings,
			assetPairsIn:     assetPairs,
			expectedErr:      errors.New("something went wrong"),
		},
		"error_cycles_converter": {
			initAssetGraphBuilder: func(mock *MockassetGraphBuilder) {
				mock.EXPECT().
					BuildGraphFromAssets(gomock.Any(), assetPairs).
					Times(1).
					Return(graph, nil)
			},
			initCyclesSearcher: func(mock *MockgraphCyclesSearcher) {
				mock.EXPECT().
					SearchAllCycles(gomock.Any(), graph).
					Times(1).
					Return(cycles, nil)
			},
			initCyclesConverter: func(mock *MockcyclesSpreadConverter) {
				mock.EXPECT().
					ConvertCyclesToSpreads(gomock.Any(), cycles).
					Times(1).
					Return(nil, errors.New("something went wrong"))
			},
			searchSettingsIn: nullSearchSettings,
			assetPairsIn:     assetPairs,
			expectedErr:      errors.New("something went wrong"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			cyclesSearcher := NewMockgraphCyclesSearcher(ctrl)
			if test.initCyclesSearcher != nil {
				test.initCyclesSearcher(cyclesSearcher)
			}

			graphBuilder := NewMockassetGraphBuilder(ctrl)
			if test.initAssetGraphBuilder != nil {
				test.initAssetGraphBuilder(graphBuilder)
			}

			cyclesConverter := NewMockcyclesSpreadConverter(ctrl)
			if test.initCyclesConverter != nil {
				test.initCyclesConverter(cyclesConverter)
			}

			profitabilityComparer := NewMockprofitabilityComparer(ctrl)
			if test.initProfitabilityComparer != nil {
				test.initProfitabilityComparer(profitabilityComparer)
			}

			ctx := context.Background()
			spreadHunter := defaultsyncspreadhunter.NewSyncSpreadHunter(
				graphBuilder, cyclesConverter, cyclesSearcher, profitabilityComparer,
			)

			actualSpreads, err := spreadHunter.SearchSpread(ctx, test.assetPairsIn, test.searchSettingsIn)

			if test.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, len(test.expectedSpreads), len(actualSpreads))
			for _, actualSpread := range actualSpreads {
				assertContainsSpread(t, test.expectedSpreads, actualSpread)
			}
		})
	}
}

func assertContainsSpread(t *testing.T, allSpreads []model.Spread, targetSpread model.Spread) {
	for _, actualSpread := range allSpreads {
		if spreadsIsEqual(actualSpread, targetSpread) {
			return
		}
	}

	t.Fatalf("all spreads(%v) not contains target spread(%v)", allSpreads, targetSpread)
}

func spreadsIsEqual(firstSpread model.Spread, secondSpread model.Spread) bool {
	if firstSpread.MetaInformation.Length != secondSpread.MetaInformation.Length {
		return false
	}
	if firstSpread.MetaInformation.Profitability != secondSpread.MetaInformation.Profitability {
		return false
	}

	firstSpreadElement := &firstSpread.Head
	secondSpreadElement := &secondSpread.Head

	for firstSpreadElement != nil {
		if firstSpreadElement.AssetPair != secondSpreadElement.AssetPair {
			return false
		}

		firstSpreadElement = firstSpreadElement.NextElement
		secondSpreadElement = secondSpreadElement.NextElement
	}

	return true
}

package spreads

import (
	"context"
	"fmt"
	"github.com/Ferum-Bot/HermesTrade/internal/financial/constants"
	asset_pairs "github.com/Ferum-Bot/HermesTrade/internal/financial/providers/asset-pairs"
	currency_ratio "github.com/Ferum-Bot/HermesTrade/internal/financial/providers/currency-ratio"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
)

const collectionName = "spreads"

type Storage struct {
	collection        *mongo.Collection
	assetPairs        []model2.AssetCurrencyPair
	spreadLinkBuilder spreadLinkBuilder
}

func NewStorage(
	database *mongo.Database,
	spreadLinkBuilder spreadLinkBuilder,
) *Storage {
	return &Storage{
		collection:        database.Collection(collectionName),
		assetPairs:        asset_pairs.ProvideAllAvailableCurrencyPairs(),
		spreadLinkBuilder: spreadLinkBuilder,
	}
}

func (s *Storage) GetSpreadsByIDs(
	ctx context.Context,
	spreadIDs []model.SpreadIdentifier,
) ([]model.SpreadWithLink, error) {
	result := make([]model.SpreadWithLink, 0)

	lengthFilter := model.SpreadsLengthFilter{
		MinLength: 2,
		MaxLength: 10,
	}
	profitabilityFilter := model.SpreadsProfitabilityFilter{
		MinProfitability: model.SpreadProfitabilityPercent{
			Precision: 1,
			Value:     5,
		},
		MaxProfitability: model.SpreadProfitabilityPercent{
			Precision: 0,
			Value:     20,
		},
	}

	for _, spreadID := range spreadIDs {
		spread := s.getSpreadFromDatabase(lengthFilter, profitabilityFilter)

		spread.Identifier = spreadID

		spreadWithLink := model.SpreadWithLink{
			Identifier:      spreadID,
			Head:            s.convertSpreadElementToSpreadWithLink(spread.Head),
			MetaInformation: spread.MetaInformation,
		}

		result = append(result, spreadWithLink)
	}

	return result, nil
}

func (s *Storage) SearchSpreads(
	ctx context.Context,
	filter model.SpreadsFilter,
	offset, limit int64,
) ([]model.Spread, error) {
	spreads := s.getSpreadsFromDatabase(filter, limit*3)
	return spreads[offset : offset+limit], nil
}

func (s *Storage) SaveSpreads(
	ctx context.Context,
	spreads []model.Spread,
) ([]model.Spread, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) getSpreadsFromDatabase(
	filter model.SpreadsFilter,
	count int64,
) []model.Spread {
	result := make([]model.Spread, 0)

	for i := int64(0); i < count; i++ {
		lengthFilter := model.SpreadsLengthFilter{
			MinLength: 2,
			MaxLength: 10,
		}
		profitabilityFilter := model.SpreadsProfitabilityFilter{
			MinProfitability: model.SpreadProfitabilityPercent{
				Precision: 1,
				Value:     5,
			},
			MaxProfitability: model.SpreadProfitabilityPercent{
				Precision: 0,
				Value:     20,
			},
		}

		if filter.LengthFilter != nil {
			lengthFilter = *filter.LengthFilter
		}
		if filter.ProfitabilityFilter != nil {
			profitabilityFilter = *filter.ProfitabilityFilter
		}

		result = append(result, s.getSpreadFromDatabase(lengthFilter, profitabilityFilter))
	}

	return result
}

func (s *Storage) getSpreadFromDatabase(
	lengthFilter model.SpreadsLengthFilter,
	profitabilityFilter model.SpreadsProfitabilityFilter,
) model.Spread {
	spreadLength := rand.Int63()%(lengthFilter.MaxLength-lengthFilter.MinLength+1) + lengthFilter.MinLength

	spread := model.Spread{
		Identifier: model.SpreadIdentifier(uuid.New().String()),
		MetaInformation: model.SpreadMetaInformation{
			Length:    model.SpreadLength(spreadLength),
			CreatedAt: time.Now(),
		},
	}

	var rootElement *model.SpreadElement
	var prevElement *model.SpreadElement
	prevElement = nil
	rootElement = nil

	startIndex := rand.Int63() % int64(len(s.assetPairs))
	currentAssetPair := s.assetPairs[startIndex]

	for i := int64(0); i < spreadLength; i++ {
		newElement := &model.SpreadElement{
			AssetPair: currentAssetPair,
		}

		if prevElement != nil {
			prevElement.NextElement = newElement
			prevElement = newElement
		} else {
			rootElement = newElement
			prevElement = newElement
		}

		if i == spreadLength-2 {
			firstAssetPair := s.assetPairs[startIndex]
			prevAssetPair := currentAssetPair

			firstName := constants.AssetNameByIdentifier[int64(prevAssetPair.QuotedAsset.UniversalIdentifier)]
			secondName := constants.AssetNameByIdentifier[int64(firstAssetPair.BaseAsset.UniversalIdentifier)]

			currentAssetPair = model2.AssetCurrencyPair{
				Identifier:    model2.AssetPairIdentifier(fmt.Sprintf("%s/%s", firstName, secondName)),
				BaseAsset:     prevAssetPair.QuotedAsset,
				QuotedAsset:   firstAssetPair.BaseAsset,
				CurrencyRatio: currency_ratio.ProvideCurrencyRatio(),
			}

		} else {
			nextAssetPairs := asset_pairs.ProvideAllNextAssetPair(s.assetPairs, currentAssetPair)
			nextIndex := rand.Int63() % int64(len(nextAssetPairs))

			currentAssetPair = nextAssetPairs[nextIndex]
		}
	}

	spread.Head = *rootElement

	spread.MetaInformation.ProfitabilityPercent = model.SpreadProfitabilityPercent{
		Precision: 3,
		Value:     rand.Int63()%(profitabilityFilter.MaxProfitability.Value-profitabilityFilter.MinProfitability.Value+1) + profitabilityFilter.MinProfitability.Value,
	}

	return spread
}

func (s *Storage) convertSpreadElementToSpreadWithLink(
	spreadElement model.SpreadElement,
) model.SpreadElementWithLink {
	root := &model.SpreadElementWithLink{}
	root = nil

	current := &spreadElement
	currentWithLink := root
	for current != nil {
		assetPairWithLink := s.spreadLinkBuilder.ProvideLinksForAssetPair(current.AssetPair)

		if root == nil {
			root = &model.SpreadElementWithLink{
				AssetPair: assetPairWithLink,
			}
			currentWithLink = root
		} else {
			newCurrentWithLink := &model.SpreadElementWithLink{
				AssetPair: assetPairWithLink,
			}
			currentWithLink.NextElement = newCurrentWithLink
			currentWithLink = newCurrentWithLink
		}

		current = current.NextElement
	}

	return *root
}

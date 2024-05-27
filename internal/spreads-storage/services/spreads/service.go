package spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type Service struct {
	storage     spreadsStorage
	linkBuilder spreadLinkBuilder
}

func NewService(
	storage spreadsStorage,
	linkBuilder spreadLinkBuilder,
) *Service {
	return &Service{
		storage:     storage,
		linkBuilder: linkBuilder,
	}
}

func (service *Service) SearchSpreads(
	ctx context.Context,
	filter model.SpreadsFilter,
	offset, limit int64,
) ([]model.Spread, error) {
	foundSpreads, err := service.storage.SearchSpreads(ctx, filter, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "service.storage.SearchSpreads")
	}

	return foundSpreads, nil
}

func (service *Service) GetSpreadsWithLinks(
	ctx context.Context,
	spreadIDs []model.SpreadIdentifier,
) ([]model.SpreadWithLink, error) {
	spreads, err := service.storage.GetSpreadsByIDs(ctx, spreadIDs)
	if err != nil {
		return nil, errors.Wrap(err, "service.storage.GetSpreads")
	}

	return spreads, nil
}

func (service *Service) SaveSpreads(
	ctx context.Context,
	spreads []model.Spread,
) ([]model.SpreadWithLink, error) {
	spreads, err := service.storage.SaveSpreads(ctx, spreads)
	if err != nil {
		return nil, errors.Wrap(err, "service.storage.SaveSpreads")
	}

	spreadsWithLink := make([]model.SpreadWithLink, 0, len(spreads))
	for _, spread := range spreads {
		spreadsWithLink = append(spreadsWithLink, service.linkBuilder.ProvideLinks(spread))
	}

	return spreadsWithLink, nil
}

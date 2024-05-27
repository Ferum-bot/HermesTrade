package get_spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
)

type spreadService interface {
	GetSpreadsWithLinks(
		ctx context.Context,
		spreadIDs []model.SpreadIdentifier,
	) ([]model.SpreadWithLink, error)
}

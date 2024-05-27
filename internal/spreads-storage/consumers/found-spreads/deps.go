package found_spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
)

type spreadsService interface {
	SaveSpreads(
		ctx context.Context,
		spreads []model.Spread,
	) ([]model.SpreadWithLink, error)
}

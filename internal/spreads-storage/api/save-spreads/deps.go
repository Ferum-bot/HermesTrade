package save_spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
)

type spreadService interface {
	SaveSpreads(
		ctx context.Context,
		spreads []model.Spread,
	) ([]model.SpreadWithLink, error)
}

package profitability

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"

type Comparer struct {
}

func NewComparer() *Comparer {
	return &Comparer{}
}

func (c *Comparer) ProfitabilityIsLessThan(
	source model.SpreadProfitability,
	than model.SpreadProfitability,
) bool {
	//TODO implement me
	panic("implement me")
}

func (c *Comparer) ProfitabilityIsGreaterThan(
	source model.SpreadProfitability,
	than model.SpreadProfitability,
) bool {
	//TODO implement me
	panic("implement me")
}

package spread_link_builder

import (
	"github.com/Ferum-Bot/HermesTrade/internal/financial/constants"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) ProvideLinks(
	spread model.Spread,
) model.SpreadWithLink {
	//TODO implement me
	panic("implement me")
}

func (service *Service) ProvideLinksForAssetPair(
	assetPair model2.AssetCurrencyPair,
) model.AssetPairWithLinks {
	assetPairWithLinks := model.AssetPairWithLinks{}

	assetPairWithLinks.AssetPair = assetPair

	switch assetPair.BaseAsset.SourceIdentifier {
	case model2.AssetSourceIdentifier(constants.BinanceSourceIdentifier):
		assetPairWithLinks.SourceLink = binanceSourceLink
	case model2.AssetSourceIdentifier(constants.ByBitSourceIdentifier):
		assetPairWithLinks.SourceLink = byBitSourceLink
	case model2.AssetSourceIdentifier(constants.CoinbaseSourceIdentifier):
		assetPairWithLinks.SourceLink = coinbaseSourceLink
	case model2.AssetSourceIdentifier(constants.KrakenSourceIdentifier):
		assetPairWithLinks.SourceLink = krakenSourceLink
	case model2.AssetSourceIdentifier(constants.OKXSourceIdentifier):
		assetPairWithLinks.SourceLink = okxSourceLink
	case model2.AssetSourceIdentifier(constants.UpBitSourceIdentifier):
		assetPairWithLinks.SourceLink = upBitSourceLink
	default:
		assetPairWithLinks.SourceLink = "unknown source"
	}

	assetPairWithLinks.PairLink = service.provideAssetPairLink(assetPair)

	return assetPairWithLinks
}

func (service *Service) provideAssetPairLink(
	assetPair model2.AssetCurrencyPair,
) string {
	switch assetPair.BaseAsset.SourceIdentifier {
	case model2.AssetSourceIdentifier(constants.BinanceSourceIdentifier):
		return "https://www.binance.com/en/trade/ETH_USDT?_from=markets&type=spot"
	case model2.AssetSourceIdentifier(constants.ByBitSourceIdentifier):
		return "https://www.bybit.com/en/trade/spot/BTC/USDT"
	case model2.AssetSourceIdentifier(constants.CoinbaseSourceIdentifier):
		return "https://www.coinbase.com/price/usdc"
	case model2.AssetSourceIdentifier(constants.KrakenSourceIdentifier):
		return "https://www.kraken.com/prices/solana"
	case model2.AssetSourceIdentifier(constants.OKXSourceIdentifier):
		return "https://www.okx.com/ru/price/bnb-bnb"
	case model2.AssetSourceIdentifier(constants.UpBitSourceIdentifier):
		return "https://sg.upbit.com/exchange?code=CRIX.UPBIT.SGD-USDT"
	}

	return "unknown asset pair"
}

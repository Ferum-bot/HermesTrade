package constants

const (
	BinanceSourceIdentifier  int64 = 1
	ByBitSourceIdentifier    int64 = 2
	CoinbaseSourceIdentifier int64 = 3
	KrakenSourceIdentifier   int64 = 4
	OKXSourceIdentifier      int64 = 5
	UpBitSourceIdentifier    int64 = 6
)

var SourceNameByIdentifier = map[int64]string{
	BinanceSourceIdentifier:  "Binance",
	ByBitSourceIdentifier:    "ByBit",
	CoinbaseSourceIdentifier: "Coinbase",
	KrakenSourceIdentifier:   "Kraken",
	OKXSourceIdentifier:      "OKX",
	UpBitSourceIdentifier:    "UpBit",
}

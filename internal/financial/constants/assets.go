package constants

const (
	USDUniversalIdentifier int64 = 1
	EURUniversalIdentifier int64 = 2
	RUBUniversalIdentifier int64 = 3
	AUDUniversalIdentifier int64 = 4
	GBPUniversalIdentifier int64 = 5
	TRYUniversalIdentifier int64 = 6
	HUFUniversalIdentifier int64 = 7

	BTCUniversalIdentifier  int64 = 8
	ETHUniversalIdentifier  int64 = 9
	BNBUniversalIdentifier  int64 = 10
	SOLUniversalIdentifier  int64 = 11
	USDTUniversalIdentifier int64 = 12
	USDCUniversalIdentifier int64 = 13
	ADAUniversalIdentifier  int64 = 14
	XRPUniversalIdentifier  int64 = 15

	AEDUniversalIdentifier int64 = 16
	CHFUniversalIdentifier int64 = 17
	JPYUniversalIdentifier int64 = 18
)

var AssetNameByIdentifier = map[int64]string{
	USDUniversalIdentifier: "USD",
	EURUniversalIdentifier: "EUR",
	RUBUniversalIdentifier: "RUB",
	AUDUniversalIdentifier: "AUD",
	GBPUniversalIdentifier: "GBP",
	TRYUniversalIdentifier: "TRY",
	HUFUniversalIdentifier: "HUF",

	BTCUniversalIdentifier:  "BTC",
	ETHUniversalIdentifier:  "ETH",
	BNBUniversalIdentifier:  "BNB",
	SOLUniversalIdentifier:  "SOL",
	USDTUniversalIdentifier: "USDT",
	USDCUniversalIdentifier: "USDC",
	ADAUniversalIdentifier:  "ADA",
	XRPUniversalIdentifier:  "XRP",

	AEDUniversalIdentifier: "AED",
	CHFUniversalIdentifier: "CHF",
	JPYUniversalIdentifier: "JPY",
}

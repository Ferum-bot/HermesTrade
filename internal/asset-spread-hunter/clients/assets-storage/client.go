package assets_storage

import "context"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) GetAssets(
	ctx context.Context,
	assetsFilter AssetsFilter,
	offset, limit int64,
) ([]AssetCurrencyPair, error) {
	//TODO implement me
	panic("implement me")
}

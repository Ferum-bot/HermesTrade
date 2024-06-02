package upbit

import "context"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) GetAssetPairs(
	ctx context.Context,
	filter string,
	offset, limit int64,
) ([]ExchangeData, error) {
	//TODO implement me
	panic("implement me")
}

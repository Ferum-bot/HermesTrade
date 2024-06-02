package assets_storage

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Save(
	ctx context.Context,
	assets []model.AssetCurrencyPair,
) error {
	//TODO implement me
	panic("implement me")
}

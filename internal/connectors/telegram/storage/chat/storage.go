package chat

import (
	"context"
	"errors"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	errors2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "chats"

type Storage struct {
	collection *mongo.Collection
}

func NewChatStorage(
	database *mongo.Database,
) *Storage {
	return &Storage{
		collection: database.Collection(collectionName),
	}
}

func (storage *Storage) GetChatByChatID(
	ctx context.Context,
	chatID model.ChatID,
) (*model.Chat, error) {
	result := storage.collection.FindOne(
		ctx,
		bson.M{
			chatFieldChatID: chatID,
		},
	)

	err := result.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		} else {
			return nil, errors2.Wrap(err, "storage.collection.FindOne")
		}
	}

	foundChat := chatRow{}
	err = result.Decode(&foundChat)
	if err != nil {
		return nil, errors2.Wrap(err, "result.Decode")
	}

	return &model.Chat{
		ChatID:            model.ChatID(foundChat.ChatID),
		ProfitabilityType: parseProfitability(foundChat.ProfitabilitySettingsType),
	}, nil
}

func (storage *Storage) UpdateChatProfitability(
	ctx context.Context,
	chatID model.ChatID,
	newProfitability model.ProfitabilitySettingsType,
) error {

	updateResult, err := storage.collection.UpdateMany(
		ctx,
		bson.D{
			{
				chatFieldChatID,
				bson.M{
					"$in": chatID,
				},
			},
		},
		bson.D{
			{
				"$set",
				bson.D{
					{chatFieldChatID, newStatus},
					{chatFieldProfitabilitySettingsType, convertProfitability(newProfitability)},
				},
			},
		},
	)
}

func (storage *Storage) CreateChat(
	ctx context.Context,
	chat model.Chat,
) error {
	//TODO implement me
	panic("implement me")
}

func (storage *Storage) DeleteChat(
	ctx context.Context,
	chatID model.ChatID,
) error {
	//TODO implement me
	panic("implement me")
}

func parseProfitability(profitability string) model.ProfitabilitySettingsType {
	switch profitability {
	case profitabilityTypeAll:
		return model.ProfitabilityAll
	case profitabilityType1Percent:
		return model.ProfitabilityPercent1
	case profitabilityType5Percent:
		return model.ProfitabilityPercent5
	case profitabilityType20Percent:
		return model.ProfitabilityPercent20
	}

	return model.ProfitabilityPercent20
}

func convertProfitability(profitability model.ProfitabilitySettingsType) string {
	switch profitability {
	case model.ProfitabilityAll:
		return profitabilityTypeAll
	case model.ProfitabilityPercent1:
		return profitabilityType1Percent
	case model.ProfitabilityPercent5:
		return profitabilityType5Percent
	case model.ProfitabilityPercent20:
		return profitabilityType20Percent
	}

	return profitabilityType20Percent
}

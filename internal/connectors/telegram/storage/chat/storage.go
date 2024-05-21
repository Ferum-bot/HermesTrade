package chat

import (
	"context"
	"errors"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	errors2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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
	now := toDatabaseTimeFormat(time.Now())

	_, err := storage.collection.UpdateOne(
		ctx,
		bson.D{
			{
				chatFieldChatID,
				chatID,
			},
		},
		bson.D{
			{
				"$set",
				bson.D{
					{chatFieldProfitabilitySettingsType, convertProfitability(newProfitability)},
					{chatFieldUpdatedAt, now},
				},
			},
		},
	)

	if err != nil {
		return errors2.Wrap(err, "storage.collection.UpdateOne")
	}

	return nil
}

func (storage *Storage) CreateChat(
	ctx context.Context,
	chat model.Chat,
) error {
	chatRow := chatRow{
		ChatID:                    int64(chat.ChatID),
		ProfitabilitySettingsType: convertProfitability(chat.ProfitabilityType),
		CreatedAt:                 toDatabaseTimeFormat(time.Now()),
		UpdatedAt:                 toDatabaseTimeFormat(time.Now()),
	}

	_, err := storage.collection.InsertOne(ctx, chatRow)
	if err != nil {
		return errors2.Wrap(err, "storage.collection.InsertOne")
	}

	return nil
}

func (storage *Storage) DeleteChat(
	ctx context.Context,
	chatID model.ChatID,
) error {
	_, err := storage.collection.DeleteOne(
		ctx,
		bson.D{
			{
				chatFieldChatID,
				chatID,
			},
		},
	)

	if err != nil {
		return errors2.Wrap(err, "storage.collection.DeleteOne")
	}

	return nil
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

func toDatabaseTimeFormat(value time.Time) int64 {
	return value.UnixMilli()
}

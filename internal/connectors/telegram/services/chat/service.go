package chat

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type Service struct {
	storage chatStorage
}

func NewChatService(
	storage chatStorage,
) *Service {
	return &Service{
		storage: storage,
	}
}

func (service *Service) UpdateChatSettings(
	ctx context.Context,
	chatID model.ChatID,
	newType model.ProfitabilitySettingsType,
) error {
	targetChat, err := service.storage.GetChatByChatID(ctx, chatID)
	if err != nil {
		return errors.Wrap(err, "service.storage.GetChatByChatID")
	}

	if targetChat == nil {
		newChat := model.Chat{
			ChatID:            chatID,
			ProfitabilityType: newType,
		}

		err = service.storage.CreateChat(ctx, newChat)
		if err != nil {
			return errors.Wrap(err, "service.storage.GetChatByChatID")
		}

		return nil
	}

	err = service.storage.UpdateChatProfitability(ctx, chatID, newType)
	if err != nil {
		return errors.Wrap(err, "service.storage.UpdateChatProfitability")
	}

	return nil
}

func (service *Service) RemoveChat(
	ctx context.Context,
	chatID model.ChatID,
) error {
	err := service.storage.DeleteChat(ctx, chatID)
	if err != nil {
		return errors.Wrap(err, "service.storage.DeleteChat")
	}

	return nil
}

func (service *Service) GetChatsCountWithProfitability(
	ctx context.Context,
	profitability model.ProfitabilitySettingsType,
) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (service *Service) GetChatsWithProfitability(
	ctx context.Context,
	profitability model.ProfitabilitySettingsType,
	pageNumber, pageSize int64,
) ([]model.Chat, error) {
	//TODO implement me
	panic("implement me")
}

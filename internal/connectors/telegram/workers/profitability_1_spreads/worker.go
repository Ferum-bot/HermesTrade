package profitability_1_spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"time"
)

type Worker struct {
	logger                 logger.Logger
	chatService            chatService
	spreadsService         spreadsService
	spreadMessageConverter spreadMessageConverter
	telegramBot            telegramBot
}

func NewWorker(
	logger logger.Logger,
	chatService chatService,
	spreadsService spreadsService,
	spreadMessageConverter spreadMessageConverter,
	telegramBot telegramBot,
) *Worker {
	return &Worker{
		logger:                 logger,
		chatService:            chatService,
		spreadsService:         spreadsService,
		spreadMessageConverter: spreadMessageConverter,
		telegramBot:            telegramBot,
	}
}

func (worker *Worker) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := worker.work(ctx)
			if err != nil {
				worker.logger.Errorf("Received error in profitability_1_spreads worker: %s", err)

				time.Sleep(10 * time.Second)
				continue
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func (worker *Worker) work(ctx context.Context) error {
	const targetProfitability = model.ProfitabilityPercent1
	const spreadMaxLength = 6
	const pageSize = 100

	chatsCount, err := worker.chatService.GetChatsCountWithProfitability(ctx, targetProfitability)
	if err != nil {
		return errors.Wrap(err, "worker.chatService.GetChatsCountWithProfitability")
	}

	iterationsCount := chatsCount / pageSize
	if chatsCount%pageSize != 0 {
		iterationsCount++
	}

	spreadParameters := model.SpreadParameters{
		MaxLength: spreadMaxLength,
		FoundTime: time.Now().Add(-time.Minute),
	}
	newSpreads, err := worker.spreadsService.GetSpreads(ctx, targetProfitability, spreadParameters)
	if err != nil {
		return errors.Wrap(err, "worker.spreadsService.GetSpreads")
	}

	for iteration := int64(0); iteration < iterationsCount; iteration++ {
		targetChats, err := worker.chatService.GetChatsWithProfitability(
			ctx, targetProfitability, iteration, pageSize,
		)
		if err != nil {
			return errors.Wrap(err, "worker.chatService.GetChatsWithProfitability")
		}

		err = worker.sendSpreadsToChats(ctx, newSpreads, targetChats)
		if err != nil {
			return errors.Wrap(err, "worker.sendSpreadsToChats")
		}
	}

	return nil
}

func (worker *Worker) sendSpreadsToChats(
	ctx context.Context,
	spreads []model2.Spread,
	chats []model.Chat,
) error {
	for _, chat := range chats {
		message := worker.spreadMessageConverter.ConvertSpreadsDataInOneMessage(spreads)

		err := worker.telegramBot.SendMessage(ctx, chat.ChatID, message)
		if err != nil {
			return errors.Wrap(err, "telegramBot.SendMessage")
		}
	}

	return nil
}

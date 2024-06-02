package found_spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	kafkaWriter *kafka.Writer
}

func NewProducer(
	kafkaWriter *kafka.Writer,
) *Producer {
	return &Producer{
		kafkaWriter: kafkaWriter,
	}
}

func (producer *Producer) Produce(
	ctx context.Context,
	spreads []model.Spread,
) error {
	//TODO implement me
	panic("implement me")
}

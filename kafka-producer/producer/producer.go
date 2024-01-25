package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer() *Producer {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "test",
	}
	return &Producer{
		Writer: writer,
	}
}

func (k *Producer) WriteMessage(ctx context.Context, messages chan kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case m := <-messages:
			err := k.Writer.WriteMessages(ctx, kafka.Message{
				Value: []byte(m.Value),
			})
			if err != nil {
				return err
			}
		}
	}
}

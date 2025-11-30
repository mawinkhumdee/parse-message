package producer

import (
	"context"
	"fmt"
	"log"
	"parse-message/config"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	Produce(ctx context.Context, message string) error
	Close() error
}

type producer struct {
	writer *kafka.Writer
}

func New(cfg config.Config) Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka.Brokers...),
		Topic:    cfg.Kafka.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &producer{
		writer: writer,
	}
}

func (p *producer) Produce(ctx context.Context, message string) error {
	msg := kafka.Message{
		Value: []byte(message),
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to write message to kafka: %w", err)
	}

	log.Printf("Produced message to topic %s: %s", p.writer.Topic, message)
	return nil
}

func (p *producer) Close() error {
	return p.writer.Close()
}

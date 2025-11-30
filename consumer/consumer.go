package consumer

import (
	"context"
	"fmt"
	"log"
	"parse-message/config"

	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Start(ctx context.Context, handler func(ctx context.Context, message string) error) error
	Close() error
}

type consumer struct {
	reader *kafka.Reader
}

func New(cfg config.Config) Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Kafka.Brokers,
		Topic:   cfg.Kafka.Topic,
		GroupID: cfg.Kafka.GroupID,
	})

	return &consumer{
		reader: reader,
	}
}

func (c *consumer) Start(ctx context.Context, handler func(ctx context.Context, message string) error) error {
	log.Printf("Starting consumer for topic %s, group %s", c.reader.Config().Topic, c.reader.Config().GroupID)

	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return fmt.Errorf("failed to read message: %w", err)
		}

		log.Printf("Received message at offset %d: %s", m.Offset, string(m.Value))

		if err := handler(ctx, string(m.Value)); err != nil {
			log.Printf("Error handling message: %v", err)
			// Decide whether to commit or not based on error type
			// For now, we continue processing
		}
	}
}

func (c *consumer) Close() error {
	return c.reader.Close()
}

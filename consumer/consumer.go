package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Start(ctx context.Context, handler func(ctx context.Context, message string) error) error
	Close() error
}

type consumer struct {
	reader *kafka.Reader
}

func New(brokers []string, topic string, groupID string) Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
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
			log.Printf("failed to read message for topic %s: %v", c.reader.Config().Topic, err)
			return fmt.Errorf("failed to read message: %w", err)
		}

		log.Printf("Received message at offset %d: %s", m.Offset, string(m.Value))

		if err := handler(ctx, string(m.Value)); err != nil {
			log.Printf("Error handling message: %v", err)
		}
	}
}

func (c *consumer) Close() error {
	return c.reader.Close()
}

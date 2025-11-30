package service

import (
	"context"
	"parse-message/model"
)

func (s *service) InsertMessage(ctx context.Context, message model.Message) error {
	// 1. Insert into DB
	_, err := s.db.Message.Insert(ctx, message)
	if err != nil {
		return err
	}

	// 2. Produce to Kafka
	// We send the content as the message for now, as the consumer expects a string to parse
	// Ideally, we might want to send the full message object or ID, but based on current consumer:
	// handler := func(ctx context.Context, message string) error { ... service.ParseMessage(ctx, message) ... }
	// It expects the raw text to parse.
	if err := s.producer.Produce(ctx, message.Content); err != nil {
		return err
	}

	return nil
}

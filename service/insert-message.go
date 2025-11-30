package service

import (
	"context"
	"encoding/json"
	"parse-message/model"
)

func (s *service) InsertMessage(ctx context.Context, message model.Message) error {
	insertedID, err := s.db.Message.Insert(ctx, message)
	if err != nil {
		return err
	}

	payload := map[string]string{
		"id":      insertedID,
		"content": message.Content,
		"user_id": message.UserID,
		"source":  message.Source,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := s.parseProducer.Produce(ctx, string(payloadBytes)); err != nil {
		return err
	}

	return nil
}

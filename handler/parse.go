package handler

import (
	"context"
	"encoding/json"
	"parse-message/model"
	"parse-message/service"
)

type parsePayload struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  string `json:"user_id,omitempty"`
	Source  string `json:"source,omitempty"`
}

func newParseHandler(service service.Service) func(ctx context.Context, message string) error {
	return func(ctx context.Context, message string) error {
		payload := parsePayload{}
		if err := json.Unmarshal([]byte(message), &payload); err != nil || payload.Content == "" {
			// fallback: treat raw message as content if payload missing
			payload.Content = message
		}

		msg := model.Message{
			ID:      payload.ID,
			UserID:  payload.UserID,
			Content: payload.Content,
			Source:  payload.Source,
		}

		_, err := service.ParseMessage(ctx, msg)
		return err
	}
}

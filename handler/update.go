package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"parse-message/model"
	"parse-message/service"
)

type updatePayload struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func newUpdateHandler(service service.Service) func(ctx context.Context, message string) error {
	return func(ctx context.Context, message string) error {
		var payload updatePayload
		if err := json.Unmarshal([]byte(message), &payload); err != nil {
			return fmt.Errorf("invalid update payload: %w", err)
		}

		return service.UpdateMessage(ctx, model.Message{ID: payload.ID, Status: payload.Status})
	}
}

package service

import (
	"context"
	"parse-message/model"
)

func (s *service) UpdateMessage(ctx context.Context, message model.Message) error {
	if err := s.db.Message.UpdateStatus(ctx, message.ID, message.Status); err != nil {
		return err
	}
	return nil
}

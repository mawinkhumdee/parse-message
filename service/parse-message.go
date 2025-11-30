package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"parse-message/model"
	"time"
)

func (s *service) ParseMessage(ctx context.Context, message model.Message) (model.ParseResult, error) {
	fields, err := s.db.Standardfield.ListActive(ctx)
	if err != nil {
		return model.ParseResult{}, err
	}

	fieldMap := make(map[string]model.StandardField, len(fields))
	for _, f := range fields {
		fieldMap[f.Key] = f
	}

	values, err := s.utils.Gemini.ExtractFields(ctx, message.Content, fields)
	if err != nil {
		return model.ParseResult{}, err
	}
	intent := inferIntentFromCatalog(values, fieldMap)

	result := model.ParseResult{
		Intent:    intent,
		Fields:    values,
		RawText:   message.Content,
		MessageID: message.ID,
		UserID:    message.UserID,
		CreatedAt: time.Now(),
	}

	insertedID, err := s.db.ParseResult.Insert(ctx, result)
	if err != nil {
		return model.ParseResult{}, fmt.Errorf("failed to insert result: %w", err)
	}
	result.ID = insertedID

	if message.ID != "" {
		updateMsg := map[string]string{
			"id":     message.ID,
			"status": "success",
		}
		if payload, err := json.Marshal(updateMsg); err == nil {
			if err := s.updateProducer.Produce(ctx, string(payload)); err != nil {
				log.Printf("Failed to produce update message: %v", err)
			}
		}
	}

	return result, nil
}

func inferIntentFromCatalog(values []model.FieldValue, catalog map[string]model.StandardField) string {
	for _, v := range values {
		if sf, ok := catalog[v.Key]; ok {
			if len(sf.Intents) > 0 {
				return sf.Intents[0]
			}
			if sf.Category != "" {
				return sf.Category
			}
		}
	}
	return "note"
}

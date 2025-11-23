package service

import (
	"context"
	"parse-message/model"
	"time"
)

func (s *service) ParseMessage(ctx context.Context, message string) (model.ParseResult, error) {
	fields, err := s.db.Standardfield.ListActive(ctx)
	if err != nil {
		return model.ParseResult{}, err
	}

	fieldMap := make(map[string]model.StandardField, len(fields))
	for _, f := range fields {
		fieldMap[f.Key] = f
	}

	values, err := s.utils.Gemini.ExtractFields(ctx, message, fields)
	if err != nil {
		return model.ParseResult{}, err
	}
	intent := inferIntentFromCatalog(values, fieldMap)

	result := model.ParseResult{
		Intent:    intent,
		Fields:    values,
		RawText:   message,
		CreatedAt: time.Now(),
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

package gemini

import (
	"context"
	"parse-message/model"
)

type GeminiClient interface {
	ExtractFields(ctx context.Context, text string, fields []model.StandardField) ([]model.FieldValue, error)
}

type gemini struct {
}

func New() GeminiClient {
	return &gemini{}
}

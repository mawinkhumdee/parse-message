package gemini

import (
	"context"
	"parse-message/model"
)

func (g *gemini) ExtractFields(ctx context.Context, text string, fields []model.StandardField) ([]model.FieldValue, error) {

	// TODO: Call Gemini API
	var amount *model.FieldValue
	if containsThaiBaht(text) {
		amount = &model.FieldValue{
			Key:   "expense.amount",
			Type:  "number",
			Value: 55.0,
		}
	}

	var result []model.FieldValue
	if amount != nil {
		result = append(result, *amount)
	}

	return result, nil
}

func containsThaiBaht(text string) bool {
	return len(text) > 0 && (contains(text, "บาท") || contains(text, "฿"))
}

func contains(s, sub string) bool {
	return stringContains(s, sub)
}

func stringContains(s, sub string) bool {
	return len(sub) > 0 && len(s) >= len(sub) && (indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
	for i := range s {
		if len(s)-i < len(sub) {
			return -1
		}
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"parse-message/model"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func (g *gemini) ExtractFields(ctx context.Context, text string, fields []model.StandardField) ([]model.FieldValue, error) {
	if g.client == nil {
		return nil, fmt.Errorf("gemini client not initialized")
	}

	prompt := g.buildPrompt(text, fields)

	genModel := g.client.GenerativeModel(g.config.Gemini.Model)
	genModel.SetTemperature(0.2)
	genModel.ResponseMIMEType = "application/json"

	resp, err := genModel.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("gemini api call failed: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return []model.FieldValue{}, nil
	}

	responseText := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	var result struct {
		Fields []model.FieldValue `json:"fields"`
	}

	if err := json.Unmarshal([]byte(responseText), &result); err != nil {
		return nil, fmt.Errorf("failed to parse gemini response: %w", err)
	}

	return result.Fields, nil
}

func (g *gemini) buildPrompt(text string, fields []model.StandardField) string {
	fieldsJSON, _ := json.MarshalIndent(fields, "", "  ")

	prompt := g.promptTemplate
	prompt = strings.ReplaceAll(prompt, "{MESSAGE}", text)
	prompt = strings.ReplaceAll(prompt, "{FIELDS}", string(fieldsJSON))

	return prompt
}

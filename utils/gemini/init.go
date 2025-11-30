package gemini

import (
	"context"
	"log"
	"parse-message/config"
	"parse-message/model"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient interface {
	ExtractFields(ctx context.Context, text string, fields []model.StandardField) ([]model.FieldValue, error)
}

type gemini struct {
	client         *genai.Client
	config         config.Config
	promptTemplate string
}

func New(cfg config.Config) GeminiClient {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.Gemini.ApiKey))
	if err != nil {
		log.Printf("Warning: Failed to initialize Gemini client: %v", err)
	}

	if cfg.Prompt.ExtractFields == "" {
		panic("prompt template for extract_fields is not loaded in config")
	}

	return &gemini{
		client:         client,
		config:         cfg,
		promptTemplate: cfg.Prompt.ExtractFields,
	}
}

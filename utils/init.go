package utils

import (
	"parse-message/config"
	"parse-message/utils/gemini"
)

type Utils struct {
	Gemini gemini.GeminiClient
}

func New(cfg config.Config) Utils {
	return Utils{
		Gemini: gemini.New(cfg),
	}
}

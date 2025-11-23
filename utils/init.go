package utils

import "parse-message/utils/gemini"

type Utils struct {
	Gemini gemini.GeminiClient
}

func New() Utils {
	return Utils{
		Gemini: gemini.New(),
	}
}

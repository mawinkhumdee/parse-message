package handler

import (
	"context"
	"parse-message/service"
)

type Handler struct {
	Parse  func(ctx context.Context, message string) error
	Update func(ctx context.Context, message string) error
}

func New(service service.Service) Handler {
	return Handler{
		Parse:  newParseHandler(service),
		Update: newUpdateHandler(service),
	}
}

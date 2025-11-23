package service

import (
	"context"
	"parse-message/config"
	"parse-message/db"
	"parse-message/model"
	"parse-message/utils"
)

type Service interface {
	ParseMessage(ctx context.Context, message string) (model.ParseResult, error)
	InsertMessage(ctx context.Context, message model.Message) error
}

type service struct {
	config config.Config
	db     db.DB
	utils  utils.Utils
}

func New(config config.Config, db db.DB, utils utils.Utils) Service {
	return &service{
		config: config,
		db:     db,
		utils:  utils,
	}
}

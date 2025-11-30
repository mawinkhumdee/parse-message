package service

import (
	"context"
	"parse-message/config"
	"parse-message/db"
	"parse-message/model"
	"parse-message/producer"
	"parse-message/utils"
)

type Service interface {
	ParseMessage(ctx context.Context, message model.Message) (model.ParseResult, error)
	InsertMessage(ctx context.Context, message model.Message) error
	UpdateMessage(ctx context.Context, message model.Message) error
}

type service struct {
	config         config.Config
	db             db.DB
	utils          utils.Utils
	parseProducer  producer.Producer
	updateProducer producer.Producer
}

func New(config config.Config, db db.DB, utils utils.Utils, parseProducer producer.Producer, updateProducer producer.Producer) Service {
	return &service{
		config:         config,
		db:             db,
		utils:          utils,
		parseProducer:  parseProducer,
		updateProducer: updateProducer,
	}
}

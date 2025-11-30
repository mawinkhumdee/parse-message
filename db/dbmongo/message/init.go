package message

import (
	"context"
	"parse-message/config"
	"parse-message/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageDB interface {
	Insert(ctx context.Context, message model.Message) (string, error)
	UpdateStatus(ctx context.Context, id string, status string) error
}

type db struct {
	collection *mongo.Collection
}

func New(cfg config.DB) MessageDB {
	clientOptions := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	collection := client.Database(cfg.DatabaseName).Collection("message")
	return &db{
		collection: collection,
	}
}

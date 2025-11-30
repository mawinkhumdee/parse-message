package parseresult

import (
	"context"
	"parse-message/config"
	"parse-message/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const parseResultCollection = "parse_result"

type ParseResultDB interface {
	Insert(ctx context.Context, pr model.ParseResult) (string, error)
}

type db struct {
	collection *mongo.Collection
}

func New(config config.DB) ParseResultDB {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil
	}
	collection := client.Database(config.DatabaseName).Collection(parseResultCollection)
	return &db{
		collection: collection,
	}
}

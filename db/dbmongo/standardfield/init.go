package standardfield

import (
	"context"
	"parse-message/config"
	"parse-message/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const standardFieldCollection = "standard_field"

type StandardFieldDB interface {
	Insert(ctx context.Context, f model.StandardField) error
	Upsert(ctx context.Context, f model.StandardField) error
	FindByKey(ctx context.Context, key string) (model.StandardField, error)
	ListActive(ctx context.Context) ([]model.StandardField, error)
}

type db struct {
	collection *mongo.Collection
}

func New(config config.DB) StandardFieldDB {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil
	}
	collection := client.Database(config.DatabaseName).Collection(standardFieldCollection)
	return &db{
		collection: collection,
	}
}

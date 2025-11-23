package standardfield

import (
	"context"
	"errors"
	"parse-message/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *db) Insert(ctx context.Context, f model.StandardField) error {
	doc, err := fromModel(&f)
	if err != nil {
		return err
	}

	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	if oid, ok := res.InsertedID.(interface{ Hex() string }); ok {
		f.ID = oid.Hex()
	}

	return nil
}

func (r *db) Upsert(ctx context.Context, f model.StandardField) error {
	doc, err := fromModel(&f)
	if err != nil {
		return err
	}

	filter := bson.M{"key": doc.Key}

	update := bson.M{
		"$set": bson.M{
			"type":        doc.Type,
			"description": doc.Description,
			"category":    doc.Category,
			"intents":     doc.Intents,
			"examples":    doc.Examples,
			"deprecated":  doc.Deprecated,
			"version":     doc.Version,
			"tags":        doc.Tags,
		},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	_ = res

	return nil
}

func (r *db) FindByKey(ctx context.Context, key string) (model.StandardField, error) {
	filter := bson.M{"key": key}

	var doc standardFieldSchema
	err := r.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.StandardField{}, nil
		}
		return model.StandardField{}, err
	}

	sf := toModel(&doc)
	return sf, nil
}

func (r *db) ListActive(ctx context.Context) ([]model.StandardField, error) {
	filter := bson.M{"deprecated": false}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []model.StandardField
	for cur.Next(ctx) {
		var doc standardFieldSchema
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		sf := toModel(&doc)
		result = append(result, sf)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

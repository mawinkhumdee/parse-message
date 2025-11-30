package message

import (
	"context"
	"fmt"
	"parse-message/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *db) Insert(ctx context.Context, message model.Message) (string, error) {
	schema := fromModel(message)
	schema.CreatedAt = time.Now()
	schema.ID = primitive.NewObjectID()

	result, err := d.collection.InsertOne(ctx, schema)
	if err != nil {
		return "", fmt.Errorf("failed to insert message: %w", err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d *db) UpdateStatus(ctx context.Context, id string, status string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := primitive.M{"_id": objectID}
	update := primitive.M{"$set": primitive.M{"status": status}}

	_, err = d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}

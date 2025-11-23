package standardfield

import "go.mongodb.org/mongo-driver/bson/primitive"

type standardFieldSchema struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Key         string             `bson:"key"`
	Type        string             `bson:"type"`
	Description string             `bson:"description,omitempty"`
	Category    string             `bson:"category"`
	Intents     []string           `bson:"intents"`
	Examples    []interface{}      `bson:"examples,omitempty"`
	Deprecated  bool               `bson:"deprecated"`
	Version     int                `bson:"version"`
	Tags        []string           `bson:"tags,omitempty"`
}

package message

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type messageSchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Content   string             `bson:"content"`
	Source    string             `bson:"source"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
}

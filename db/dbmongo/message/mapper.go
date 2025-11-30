package message

import (
	"parse-message/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func toModel(s messageSchema) model.Message {
	return model.Message{
		ID:        s.ID.Hex(),
		UserID:    s.UserID,
		Content:   s.Content,
		Source:    s.Source,
		Status:    s.Status,
		CreatedAt: s.CreatedAt,
	}
}

func fromModel(m model.Message) messageSchema {
	id, _ := primitive.ObjectIDFromHex(m.ID)
	return messageSchema{
		ID:        id,
		UserID:    m.UserID,
		Content:   m.Content,
		Source:    m.Source,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
	}
}

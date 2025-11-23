package standardfield

import (
	"parse-message/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func fromModel(sf *model.StandardField) (*standardFieldSchema, error) {
	var id primitive.ObjectID
	var err error

	if sf.ID != "" {
		id, err = primitive.ObjectIDFromHex(sf.ID)
		if err != nil {
			return nil, err
		}
	}

	return &standardFieldSchema{
		ID:          id,
		Key:         sf.Key,
		Type:        sf.Type,
		Description: sf.Description,
		Category:    sf.Category,
		Intents:     sf.Intents,
		Examples:    sf.Examples,
		Deprecated:  sf.Deprecated,
		Version:     sf.Version,
		Tags:        sf.Tags,
	}, nil
}

func toModel(doc *standardFieldSchema) model.StandardField {
	var id string
	if !doc.ID.IsZero() {
		id = doc.ID.Hex()
	}

	return model.StandardField{
		ID:          id,
		Key:         doc.Key,
		Type:        doc.Type,
		Description: doc.Description,
		Category:    doc.Category,
		Intents:     doc.Intents,
		Examples:    doc.Examples,
		Deprecated:  doc.Deprecated,
		Version:     doc.Version,
		Tags:        doc.Tags,
	}
}

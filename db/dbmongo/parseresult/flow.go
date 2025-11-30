package parseresult

import (
	"context"
	"parse-message/model"
)

func (r *db) Insert(ctx context.Context, pr model.ParseResult) (string, error) {
	doc, err := fromModel(&pr)
	if err != nil {
		return "", err
	}

	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}

	if oid, ok := res.InsertedID.(interface{ Hex() string }); ok {
		return oid.Hex(), nil
	}

	return "", nil
}

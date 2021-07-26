package mongo

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	stdmongo "go.mongodb.org/mongo-driver/mongo"
)

type (
	MongoDriver struct {
		Client *stdmongo.Collection
	}
)

func (driver *MongoDriver) Save(ctx context.Context, entity interface{}) error {
	document, err := bson.Marshal(entity)
	if err != nil {
		return err
	}

	_, err = driver.Client.InsertOne(ctx, document)
	return err
}

func (driver *MongoDriver) Find(ctx context.Context, criteria interface{}) ([]byte, error) {
	cur, err := driver.Client.Find(ctx, criteria)
	if err != nil {
		return nil, err
	}

	dbResult := make([]bson.M, 0)
	err = cur.All(ctx, &dbResult)
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(dbResult)
	if err != nil {
		return nil, err
	}

	return result, nil
}

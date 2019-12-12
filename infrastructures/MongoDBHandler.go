package infrastructures

import (
	"context"
	"github.com/djamboe/mtools-login-service/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandler struct {
	Conn *mongo.Client
}

func (handler *MongoDBHandler) FindOne(filter bson.M, collectionName string, dbName string) (interfaces.IRowMongoDB, error) {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	rows := collection.FindOne(context.TODO(), filter)
	row := new(MongoRow)
	row.Rows = rows
	return row, nil
}

type MongoRow struct {
	Rows *mongo.SingleResult
}

func (r *MongoRow) DecodeResults(v interface{}) error {
	return r.Rows.Decode(v)
}

func (r MongoRow) Next() bool {
	return r.Next()
}

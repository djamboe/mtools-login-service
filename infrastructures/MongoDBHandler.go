package infrastructures

import (
	"context"
	"fmt"
	"github.com/djamboe/mtools-login-service/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandler struct {
	Conn *mongo.Client
}

type Hero struct {
	Username string `json:"username"`
}

func (handler *MongoDBHandler) FindOne(filter bson.M, collectionName string, dbName string) (interfaces.IRowMongoDB, error) {
	var hero Hero
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	rows := collection.FindOne(context.TODO(), filter)
	row := new(MongoRow)
	row.Rows = rows
	rows.Decode(&hero)
	fmt.Println("from single result", hero.Username)
	return row, nil
}

type MongoRow struct {
	Rows *mongo.SingleResult
}

func (r *MongoRow) DecodeResults(v interface{}) error {
	fmt.Println("from interface", r.Rows.Decode(v))
	return r.Rows.Decode(v)
}

func (r MongoRow) Next() bool {
	return r.Next()
}

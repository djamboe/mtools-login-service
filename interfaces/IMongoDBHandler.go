package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
)

type IMongoDBHandler interface {
	FindOne(filter bson.M, collectionName string, databaseName string) (IRowMongoDB, error)
}
type IRowMongoDB interface {
	Next() bool
	DecodeResults(v interface{}) error
}

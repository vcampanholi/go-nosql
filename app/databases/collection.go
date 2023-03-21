package databases

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go_nosql_db").Collection("Routes")
	return collection
}

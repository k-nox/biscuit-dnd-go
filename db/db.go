package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func getOptions() error {
	return nil
}

func getClient() (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	return client, err
}

func getDatabase(client *mongo.Client, database string, collection string) *mongo.Collection {
	db := client.Database(database).Collection(collection)
	return db
}

func GetCollection(collection string) *mongo.Collection {
	client, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	coll := client.Database("5e-database").Collection(collection)
	return coll
}

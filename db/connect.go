package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error when connecting to Database: %v", err.Error())
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error when pinging to Database: %v", err.Error())
	}
	return client
}

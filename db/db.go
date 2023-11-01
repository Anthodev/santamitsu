package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type CollectionRequest struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func connect() *mongo.Client {
	uri := os.Getenv("DATABASE_URI")

	c, errCo := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if errCo != nil {
		log.Fatal(errCo)
	}

	return c
}

func disconnect(c *mongo.Client) {
	defer func() {
		if err := c.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
}

func GetCollection() CollectionRequest {
	var dbName = os.Getenv("DATABASE_NAME")
	var collectionName = os.Getenv("DATABASE_COLLECTION")

	client := connect()

	db := client.Database(dbName)

	collection := CollectionRequest{
		client:     client,
		collection: db.Collection(collectionName),
	}

	if collection.collection == nil {
		err := db.CreateCollection(context.TODO(), collectionName)

		if err != nil {
			log.Fatal(err)
		}

		collection := CollectionRequest{
			client:     client,
			collection: db.Collection(collectionName),
		}

		return collection
	}

	return collection
}

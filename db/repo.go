package db

import (
	"anthodev/santamitsu/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InsertSantaSecret(s model.SantaSecret) {
	c := GetCollection()

	_, err := c.collection.InsertOne(context.TODO(), s)

	if err != nil {
		panic(err)
	}

	defer disconnect(c.client)
}

func FindOneSantaSecret(channelID string) model.SantaSecret {
	c := GetCollection()

	filter := bson.M{"channelid": channelID}

	var s model.SantaSecret

	err := c.collection.FindOne(context.TODO(), filter).Decode(&s)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return s
		}
	}

	defer disconnect(c.client)

	return s
}

func UpdateSantaSecret(s model.SantaSecret) model.SantaSecret {
	c := GetCollection()

	filter := bson.M{"channelid": s.ChannelID}

	update := bson.M{
		"$set": s,
	}

	err := c.collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&s)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return s
		}
		log.Fatal(err)
	}

	defer disconnect(c.client)

	return s
}

func DeleteOneSantaSecret(channelID string) {
	c := GetCollection()

	filter := bson.M{"channelid": channelID}

	_, err := c.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	defer disconnect(c.client)
}

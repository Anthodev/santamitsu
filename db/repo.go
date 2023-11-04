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

func UpdateSantaSecret(s *model.SantaSecret) {
	c := GetCollection()

	filter := bson.M{"channel_id": s.ChannelID}

	var updatedSantaSecret bson.M

	doc := bson.M{
		"title":          s.Title,
		"description":    s.Description,
		"max_price":      s.MaxPrice,
		"participants":   s.Participants,
		"excluded_pairs": s.ExcludedPairs,
	}

	err := c.collection.FindOneAndUpdate(context.TODO(), filter, doc).Decode(&updatedSantaSecret)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		log.Fatal(err)
	}

	defer disconnect(c.client)
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

package core

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateRoom is responsible for create a new document at mongodb
func CreateRoom(name string) (id *mongo.InsertOneResult) {
	var result mongo.InsertOneResult
	Exec(10*time.Second, "rooms", func(ctx context.Context, collection *mongo.Collection) {
		insertResult, err := collection.InsertOne(ctx, bson.D{
			{Key: "name", Value: name},
		})

		if err != nil {
			log.Fatalln("Error on insertOne", err)
			return
		}
		result = *insertResult
	})
	return &result
}

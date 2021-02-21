package core

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GetMongoClient gets a client connection
func GetMongoClient() (mongo.Client, context.Context) {
	uri := "mongodb://localhost:27017/"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return *client, ctx
}

// ConnectionPing pings the mongodb connection
func ConnectionPing() {
	client, ctx := GetMongoClient()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

// RegisterRoom is responsible for create a new document at mongodb
func RegisterRoom(name string) (id *mongo.InsertOneResult) {
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

// Exec executes some operation in mongodb
func Exec(timeout time.Duration, collectionName string, query func(context.Context, *mongo.Collection)) {
	//client, ctx := GetMongoClient()
	uri := "mongodb://localhost:27017/"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln("error trying to connect with mongodb", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalln("error on disconnect mongodb", err)
		}
	}()

	collection := client.Database("hotel1").Collection(collectionName)

	query(ctx, collection)
}

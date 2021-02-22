package core

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
	Next tutorial, READ and FILTER
	https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
	https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver-pt
*/

// ConnectionPing pings the mongodb connection
func ConnectionPing() {
	ExecClient(5*time.Second, func(ctx context.Context, client *mongo.Client) {
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected and pinged.")
	})
}

// Exec gives an execution context with collection
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

// ExecClient gives an execution context with a mongo client
func ExecClient(timeout time.Duration, query func(context.Context, *mongo.Client)) {
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

	query(ctx, client)
}

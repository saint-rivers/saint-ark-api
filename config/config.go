package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017")
	// clientOptions := options.Client().ApplyURI("mongodb://root:password@mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	// defer client.Disconnect(context.TODO())

	return client

	// collection := client.Database("saint-ark").Collection("resources")

	// return collection
}

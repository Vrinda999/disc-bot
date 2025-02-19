package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() *mongo.Client {
	// Set MongoDB URI
	uri := "mongodb://localhost:27017"

	// clientOptions := options.Client().ApplyURI(uri)

	// client, err := mongo.NewClient(clientOptions)
	// if err != nil {
	// 	log.Fatal("Error creating MongoDB client:", err)
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal("Error connecting to MongoDB:", err)
	// }

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		// error
	}

	// Check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Assign the database
	DB = client.Database("iconic-lines")

	return client
}

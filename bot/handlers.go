package bot

import (
	"context"
	"fmt"
	"iconic-lines/models"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StoreMessage(db *mongo.Database, content string, author string) {
	collection := db.Collection("messages")
	message := models.Message{
		Content: content,
		Author:  author,
	}

	res, err := collection.InsertOne(context.TODO(), message)
	if err != nil {
		log.Println("Error storing message:", err)
	}

	// Console log for success
	fmt.Printf("✅ Message stored successfully! ID: %v | Content: %s | Author: %s\n", res.InsertedID, content, author)

}

// Create a new random generator
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomMessage(db *mongo.Database) string {
	collection := db.Collection("messages")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error retrieving messages:", err)
		return "I have no messages stored yet."
	}
	defer cursor.Close(context.TODO())

	var messages []models.Message
	for cursor.Next(context.TODO()) {
		var msg models.Message
		cursor.Decode(&msg)
		messages = append(messages, msg)
	}

	if len(messages) == 0 {
		return "No messages found."
	}

	// Use the local random generator
	randomMsg := messages[rng.Intn(len(messages))]

	return randomMsg.Content
}

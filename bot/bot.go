package bot

import (
	"fmt"
	"iconic-lines/config"
	"log"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
)

var BotID string
var goBot *discordgo.Session
var db *mongo.Database // Global database variable

func Start(client *mongo.Client) {
	db = client.Database("discordbot") // Assign the database

	var err error
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal("Error creating bot session:", err.Error())
	}

	// Set Intents BEFORE opening connection
	goBot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println("Error getting bot user:", err.Error())
		return
	}
	BotID = user.ID

	goBot.AddHandler(messageHandler)

	// Open connection
	err = goBot.Open()
	if err != nil {
		log.Fatal("Error opening connection:", err.Error())
	}

	fmt.Println("Bot is Running!")
}

func messageHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	// Makes sure that Bot does not take up its own Msgs as Commands
	if msg.Author.ID == BotID {
		return
	}

	if msg.Content == "ping" {
		sess.ChannelMessageSend(msg.ChannelID, "pong")
	}
	if msg.Content == "!store" {
		StoreMessage(db, msg.Content, msg.Author.Username)
		sess.ChannelMessageSend(msg.ChannelID, "Message stored!")
	}

	if msg.Content == "!respond" {
		message := GetRandomMessage(db)
		sess.ChannelMessageSend(msg.ChannelID, message)
	}

	fmt.Println("Received message:", msg.Content)
}

package bot

import (
	"fmt"
	"iconic-lines/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	var err error
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error creating bot session:", err.Error())
		return
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
		fmt.Println("Error opening connection:", err.Error())
		return
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

	fmt.Println("Received message:", msg.Content)
}

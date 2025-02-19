package main

import (
	"fmt"
	"iconic-lines/bot"
	"iconic-lines/config"
	"iconic-lines/db"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Initialize MongoDB
	client := db.ConnectDB()
	defer client.Disconnect(nil) // Ensure we close the connection

	// Start the Discord bot
	bot.Start(client)

	<-make(chan struct{})

	fmt.Println("Bot is running...")
	select {} // Keeps the bot running
}

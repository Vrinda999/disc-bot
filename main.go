package main

import (
	"fmt"
	"iconic-lines/bot"
	"iconic-lines/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}

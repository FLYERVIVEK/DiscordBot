package main

import (
	"fmt"
	"github/flyervivek/discordbot/bot"
	"github/flyervivek/discordbot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}

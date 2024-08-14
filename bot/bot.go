package bot

import (
	"fmt"
	"github/flyervivek/discordbot/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "Hii" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hey Bro, How are You Doing!")
	}

	if m.Content == "Who are you" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "I am a discorbot made by Vivek Yadav")
	}

}

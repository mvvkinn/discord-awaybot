package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// shrc export AWAYBOTTOKEN="put your token here"
	token := os.Getenv("AWAYBOTTOKEN")
	fmt.Println(token)

	// Create new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session : ", err)
		return
	}

	dg.AddHandler(messageCreate)

	// Establish connection to discord server
	err = dg.Open() // connection variable
	if err != nil {
		fmt.Println("Error establish connection, ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close discord session
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate, stat *discordgo.State) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")

		fmt.Println(stat)
	}
}

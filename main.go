package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/poodlenoodle42/Discord_Survey_Bot/config"
	"github.com/poodlenoodle42/Discord_Survey_Bot/database"
)

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	} else if m.Content[0] != '$' {
		return
	} else if m.Content[1:3] == "new" {
		poll, err := database.New_Poll(m.Author.Username, m.Content, m.MentionRoles[0])
		//Sent poll
	} else if m.Content[1:4] == "vote" {

	}

	//m.Member.Roles
}

func main() {
	config := config.ReadConfig("config.yaml")
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err)
	}
	discord.AddHandler(onMessageCreate)
	err = discord.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

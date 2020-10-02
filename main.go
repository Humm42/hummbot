package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var stop chan os.Signal

func Die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}

func Usage() {
	Die("usage: %s token", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		Usage()
	}
	token := os.Args[1]
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		Die("creating Discord session: %v", err)
	}

	session.AddHandler(messageHandler)

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = session.Open()
	if err != nil {
		Die("connecting with Discord: %v", err)
	}
	defer session.Close()

	stop = make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
}

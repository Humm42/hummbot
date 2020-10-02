package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// TODO: use $PLAN9, I guess
var fortunes = map[string]string{
	"dougfacts": "/usr/src/os/plan9front/lib/dougfacts",
	"ken":       "/usr/src/os/plan9front/lib/ken",
	"rob":       "/usr/src/os/plan9front/lib/rob",
	"rsc":       "/usr/src/os/plan9front/lib/rsc",
	"terry":     "/usr/src/os/plan9front/lib/terry",
	"theo":      "/usr/src/os/plan9front/lib/theo",
	"troll":     "/usr/src/os/plan9front/lib/troll",
	"uriel":     "/usr/src/os/plan9front/lib/uriel",
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	content := strings.TrimPrefix(m.Content, "HummBot: ")
	if content == m.Content {
		return
	}
	fmt.Printf("%s:%s <%s> %s\n", m.GuildID, m.ChannelID, m.Author.ID, m.Content)
	var message string
	if content == "names" {
		for name := range fortunes {
			message += " " + name
		}
		message += "\n"
	} else {
		message = fortune(content)
	}
	fmt.Printf(" -> %s", message)
	s.ChannelMessageSend(m.ChannelID, message)
}

func fortune(s string) string {
	path, ok := fortunes[s]
	if !ok {
		return "Misfortune!\n"
	}
	cmd := exec.Command("9", "fortune", path)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("`9 fortune %s`: %v", path, err)
	}
	return stdout.String()
}

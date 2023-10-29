package utils

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func CreateDmChannel(s *discordgo.Session, id string) *discordgo.Channel {
	uc, err := s.UserChannelCreate(id)

	if err != nil {
		log.Fatalf("error creating dm channel: %v", err)
	}

	return uc
}

func SendDm(s *discordgo.Session, id string, content string) *discordgo.Message {
	msg, err := s.ChannelMessageSend(id, content)

	if err != nil {
		log.Fatalf("error sending dm: %v", err)
	}

	return msg
}

func SendDmEmbed(s *discordgo.Session, id string, embed *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := s.ChannelMessageSendEmbed(id, embed)

	if err != nil {
		log.Fatalf("error sending dm embed: %v", err)
	}

	return msg
}

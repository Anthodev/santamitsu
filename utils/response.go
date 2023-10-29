package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func SendInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, content string, ephemeral bool) {
	ird := &discordgo.InteractionResponseData{
		Content: content,
		TTS:     false,
	}

	if ephemeral {
		ird.Flags = discordgo.MessageFlagsEphemeral
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: ird,
	})
}

func SendFollowupMessage(s *discordgo.Session, i *discordgo.InteractionCreate, content string) *discordgo.Message {
	msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: content,
	})

	if err != nil {
		fmt.Println("error creating a followup message", err)
	}

	return msg
}

func SendFollowupEmbedMessage(s *discordgo.Session, i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{embed},
	})

	if err != nil {
		fmt.Println("error creating a followup message", err)
	}

	return msg
}

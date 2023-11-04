package response

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func SendInteractionResponse(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	content string,
	ephemeral bool,
) {
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

func SendInteractionEmbedResponse(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	content string,
	embed *discordgo.MessageEmbed,
) {
	ird := &discordgo.InteractionResponseData{
		Content: content,
		Embeds:  []*discordgo.MessageEmbed{embed},
		TTS:     false,
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: ird,
	})
}

func SendInteractionEmbedResponseWithActionRow(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ird *discordgo.InteractionResponseData,
) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: ird,
	})
}

func BuildComponent(ar *discordgo.ActionsRow) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		ar,
	}
}

func BuildInteractionResponseData(
	t string,
	c string,
	f discordgo.MessageFlags,
	cmp []discordgo.MessageComponent,
) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Title:      t,
		Content:    c,
		Flags:      f,
		Components: cmp,
		TTS:        false,
	}
}

func UpdateInteractionEmbedResponse(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	e []*discordgo.MessageEmbed,
) {
	_, err := s.FollowupMessageEdit(i.Interaction, i.Message.ID, &discordgo.WebhookEdit{
		Embeds: &e,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func BuildInteractionEmbedResponseData(
	f discordgo.MessageFlags,
	e *discordgo.MessageEmbed,
	cmp []discordgo.MessageComponent,
) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Flags:      f,
		Embeds:     []*discordgo.MessageEmbed{e},
		Components: cmp,
		TTS:        false,
	}
}

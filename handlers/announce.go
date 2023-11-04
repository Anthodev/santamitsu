package handlers

import (
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func AnnounceHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SantaSecret) {
	embed := &discordgo.MessageEmbed{
		Title:       "Launch of the Secret Santa " + ss.Title + "!",
		Description: ss.Description + "\n\nMaximum budget per participant:\n" + ss.MaxPrice + ss.Currency + "\n\nParticipants:\n\n",
	}

	for _, p := range ss.Participants {
		pcp, err := s.GuildMember(i.ChannelID, p.UserId)
		if err != nil {
			continue
		}

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   p.UserId,
			Value:  pcp.Mention(),
			Inline: false,
		})
	}

	ar := actionRow()

	ird := response.BuildInteractionEmbedResponseData(
		0,
		embed,
		[]discordgo.MessageComponent{ar},
	)

	response.SendInteractionEmbedResponseWithActionRow(s, i, ird)
}

func actionRow() *discordgo.ActionsRow {
	return &discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			&discordgo.Button{
				Label:    "Join",
				Style:    discordgo.SuccessButton,
				CustomID: "join",
			},
		},
	}
}

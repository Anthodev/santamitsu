package handlers

import (
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/component"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func infoHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SecretSanta) {
	embed := &discordgo.MessageEmbed{
		Title:       "Secret Santa: " + ss.Title,
		Description: ss.Description + "\n\nMaximum budget per participant:\n" + ss.MaxPrice + ss.Currency + "\n\nParticipants:\n\n",
	}

	for _, p := range ss.Participants {
		pcp, err := s.GuildMember(i.GuildID, p.UserId)
		if err != nil {
			continue
		}

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   pcp.Nick,
			Value:  pcp.Mention(),
			Inline: false,
		})
	}

	ar := component.JoinLeaveActionRow()

	ird := response.BuildInteractionEmbedResponseData(
		0,
		embed,
		[]discordgo.MessageComponent{ar},
	)

	response.SendInteractionEmbedResponseWithActionRow(s, i, ird)
}

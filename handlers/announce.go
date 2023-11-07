package handlers

import (
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/service"
	"anthodev/santamitsu/utils/component"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func announceHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SecretSanta) {
	service.LockState(s, i, ss)
	service.IsMemberAuthorized(s, i, ss)

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

	ar := component.JoinLeaveActionRow()

	ird := response.BuildInteractionEmbedResponseData(
		0,
		embed,
		[]discordgo.MessageComponent{ar},
	)

	response.SendInteractionEmbedResponseWithActionRow(s, i, ird)
}

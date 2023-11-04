package response

import (
	"anthodev/santamitsu/model"
	"github.com/bwmarrin/discordgo"
)

func UpdateEmbedParticipantList(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SantaSecret) {
	msgEmbeds := i.Message.Embeds
	me := msgEmbeds[0]

	for _, p := range ss.Participants {
		m, _ := s.GuildMember(i.GuildID, p.UserId)

		me.Fields = append(me.Fields, &discordgo.MessageEmbedField{
			Name:   m.Nick,
			Value:  m.Mention(),
			Inline: false,
		})
	}

	UpdateInteractionEmbedResponse(s, i, msgEmbeds)
}

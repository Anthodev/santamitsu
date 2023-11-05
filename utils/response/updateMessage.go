package response

import (
	"anthodev/santamitsu/model"
	"github.com/bwmarrin/discordgo"
)

func UpdateEmbedParticipantList(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SecretSanta) {
	msgEmbeds := i.Message.Embeds
	me := msgEmbeds[0]

	var fields []*discordgo.MessageEmbedField

	for _, p := range ss.Participants {
		m, _ := s.GuildMember(i.GuildID, p.UserId)

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   m.Nick,
			Value:  m.Mention(),
			Inline: false,
		})
	}

	me.Fields = fields

	UpdateInteractionEmbedResponse(s, i, msgEmbeds)
}

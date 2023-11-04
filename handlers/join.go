package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func JoinHandler(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	ss model.SantaSecret,
	isMsg bool,
) {
	p := model.SantaParticipant{
		UserId: i.Member.User.ID,
	}

	for _, p := range ss.Participants {
		if p.UserId == i.Member.User.ID {
			response.SendInteractionResponse(s, i, "You have already joined the secret santa!", true)

			return
		}
	}

	ss = model.AddParticipant(ss, p)

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "You have joined the secret santa!", true)

	if isMsg == true {
		response.UpdateEmbedParticipantList(s, i, ss)
	}
}

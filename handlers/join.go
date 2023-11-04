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

	ss = model.AddParticipant(ss, p)

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "You have joined the secret santa!", true)

	if isMsg {
		response.UpdateEmbedParticipantList(s, i, ss)
	}
}

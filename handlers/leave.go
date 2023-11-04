package handlers

import (
	"anthodev/santamitsu/db"
	"anthodev/santamitsu/model"
	"anthodev/santamitsu/utils/response"
	"github.com/bwmarrin/discordgo"
)

func LeaveHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ss model.SantaSecret) {
	var u model.SantaParticipant

	for _, p := range ss.Participants {
		if p.UserId == i.Member.User.ID {
			u = p
			break
		}
	}

	if u.UserId == "" {
		return
	}

	ss = model.RemoveParticipant(ss, u)

	db.UpdateSantaSecret(ss)

	response.SendInteractionResponse(s, i, "You have left the secret santa!", true)
}
